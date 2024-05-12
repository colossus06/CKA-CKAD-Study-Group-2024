from flask import Flask, request, jsonify
import mysql.connector
import os
from retry_utils import retry_on_operational_error

app = Flask(__name__)
db = mysql.connector.connect(
    #kubernetes
    host = 'mysql-svc',
    #docker-compose
    # host="mysql",
    user=os.getenv('MYSQL_USER', 'colo'),
    password=os.getenv('MYSQL_PASSWORD', 'colo'),
    database=os.getenv('MYSQL_DB', 'ckad_crud'),
    port=os.getenv('MYSQL_PORT', '3306')
)
cursor = db.cursor()

class Task:
    def __init__(self, id, taskName):
        self.id = id
        self.taskName = taskName

@app.route('/tasks', methods=['GET'])
def get_tasks():
    try:
        if not db.is_connected():
            db.reconnect()

        cursor.execute("SELECT * FROM tasks")
        tasks = cursor.fetchall()
        result = []
        for task in tasks:
            result.append({'id': task[0], 'taskName': task[1]})
        return jsonify(result)
    except mysql.connector.Error as err:
        return jsonify({'error': f"Failed to fetch tasks: {err}"}), 500



@app.route('/tasks/<int:id>', methods=['GET'])
def get_task_by_id(id):
    cursor.execute("SELECT * FROM tasks WHERE id = %s", (id,))
    task = cursor.fetchone()
    if not task:
        return jsonify({'error': 'Task not found'}), 404
    return jsonify({'id': task[0], 'taskName': task[1]})

@app.route('/tasks', methods=['POST'])
def add_task():
    data = request.get_json()
    task_name = data.get('taskName')
    if not task_name:
        return jsonify({'message': 'Task name cannot be empty'}), 400

    cursor.execute("INSERT INTO tasks (taskName) VALUES (%s)", (task_name,))
    db.commit()

    return jsonify({'message': 'Task added successfully'}), 200

@app.route('/tasks/edit/<int:id>', methods=['PUT'])
def update_task(id):
    data = request.get_json()
    task_name = data['taskName']
    cursor.execute("UPDATE tasks SET taskName = %s WHERE id = %s", (task_name, id))
    db.commit()
    return jsonify({'message': 'Task updated successfully'})



@app.route('/tasks/delete/<int:id>', methods=['DELETE'])
def delete_task(id):
    cursor.execute("DELETE FROM tasks WHERE id = %s", (id,))
    db.commit()
    return jsonify({'message': 'Task deleted successfully'})


@app.route('/tasks/count', methods=['GET'])
def get_task_count():
    cursor.execute("SELECT COUNT(*) FROM tasks")
    count = cursor.fetchone()[0]
    if count > 3:
        # Send email logic here
        print("Sending email notification: Too many tasks!")
    return jsonify({'count': count})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
