//! Structured logging utilities (stub).

use super::config::LoggingConfig;

#[derive(Debug, thiserror::Error)]
pub enum LogError {
    #[error("logging init error: {0}")]
    Init(String),
}

/// Initialize structured logging (stub: no-op).
pub fn init_logging(_config: &LoggingConfig) -> Result<(), LogError> {
    Ok(())
}
