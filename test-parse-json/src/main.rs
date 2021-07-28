const JSON: &str = r#"{"jsonrpc":"2.0","method":"$/progress","params":{"token":"rustAnalyzer/Indexing","value":{"kind":"report","message":"65 /751 (subtle)","percentage":8}}}"#;
fn main() {
    println!("Hello, world!");
    let parsed: serde_json::Value = serde_json::from_str(JSON).unwrap();
    println!("Parsed: {}", parsed);
}
