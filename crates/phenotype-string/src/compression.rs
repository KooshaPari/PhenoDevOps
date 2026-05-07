//! String compression utilities (stub).

use super::Error;

/// Compress data using zstd (stub: returns uncompressed data).
pub fn compress(data: &[u8]) -> Result<Vec<u8>, Error> {
    Ok(data.to_vec())
}

/// Decompress zstd data (stub: returns data as-is).
pub fn decompress(data: &[u8]) -> Result<Vec<u8>, Error> {
    Ok(data.to_vec())
}
