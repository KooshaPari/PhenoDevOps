//! Router types (stub).

pub trait Router {
    fn route(&self, request: &super::models::RoutingRequest) -> super::error::Result<super::models::RouterDecision>;
}
