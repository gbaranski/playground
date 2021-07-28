use criterion::{black_box, criterion_group, criterion_main, Criterion};
use bincode_vs_json::Block;

fn criterion_benchmark(c: &mut Criterion) {
    let block = Block {
        index: 0,
        timestamp: chrono::MAX_DATETIME,
        message: String::new(),
        prev_hash: String::new(),
    };
    c.bench_function("serde json", |b| b.iter(|| serde_json::to_string(&block)));
    c.bench_function("bincode", |b| b.iter(|| bincode::serialize(&block)));
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);
