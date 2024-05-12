use actix_web::{web, App, HttpServer, Responder, HttpRequest, HttpResponse};
use serde::{Deserialize, Serialize};

#[derive(Debug, Deserialize, Serialize)]
struct TaskCount {
    count: usize,
}

async fn get_task_count() -> Result<TaskCount, reqwest::Error> {
    let response = reqwest::blocking::get("http://crud-svc:5000/tasks/count")?
        .json::<TaskCount>()?;
    Ok(response)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .route("/notification", web::get().to(get_task_count_handler))
    })
    .bind("0.0.0.0:8083")?
    .run()
    .await
}

async fn get_task_count_handler() -> impl Responder {
    match get_task_count().await {
        Ok(task_count) => {
            let notification_message = format!("Total task count is: {}. Notification sent", task_count.count);
            println!("{}", notification_message);
            notification_message
        }
        Err(_) => "Failed to fetch task count".to_string(),
    }
}