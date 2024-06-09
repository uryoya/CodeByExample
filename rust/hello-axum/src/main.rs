use axum::{extract::State, routing::get, Router};
use std::sync::Arc;

struct AppState {
    foo: String,
}

#[tokio::main]
async fn main() {
    let shared_state = Arc::new(AppState {
        foo: "foooooooooo".to_string(),
    });

    let app = Router::new()
        .route("/", get(root))
        .route("/foo", get(get_foo).post(post_foo))
        .with_state(shared_state);

    let listener = tokio::net::TcpListener::bind("localhost:3000")
        .await
        .unwrap();
    axum::serve(listener, app).await.unwrap();
}

async fn root(State(state): State<Arc<AppState>>) -> String {
    String::from(&state.foo)
}

async fn get_foo() -> &'static str {
    "foo"
}

async fn post_foo() -> &'static str {
    "Foo"
}
