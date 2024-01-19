from flask import Flask
import time

app = Flask(__name__)

@app.route('/')
def hello_gang():
    current_time = time.strftime("%Y-%m-%d %H:%M")
    return f'{current_time} - Good job on sharpening skills! Time to dockerize.'
hello_gang()

if __name__ == '__main__':
    app.run(port=8000)
