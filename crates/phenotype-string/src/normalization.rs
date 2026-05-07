//! String normalization utilities (stub).

/// Normalize a string to NFC form.
pub fn normalize_nfc(input: &str) -> String {
    use unicode_normalization::UnicodeNormalization;
    input.nfc().collect()
}

/// Normalize a string to lowercase ASCII.
pub fn normalize_ascii_lowercase(input: &str) -> String {
    input.to_ascii_lowercase()
}

/// Trim and collapse whitespace.
pub fn normalize_whitespace(input: &str) -> String {
    input.split_whitespace().collect::<Vec<_>>().join(" ")
}
