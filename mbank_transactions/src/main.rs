use scraper::{Html, Selector};
use serde::{Deserialize, Serialize};
use std::fs::File;
use std::io::{Read, Write};

const FILE_PATH: &'static str = "settlements.json";

#[derive(Debug, Serialize, Deserialize)]
struct Transaction {
    date: String,
    location: String,
    price: f64,
}

#[derive(Debug, Default, Serialize, Deserialize)]
struct Summary {
    mom: PerSummary,
    me: PerSummary,
    shared: PerSummary,
}

#[derive(Debug, Default, Serialize, Deserialize)]
struct PerSummary {
    total: f64,
    transactions: Vec<Transaction>,
}

fn main() {
    let dialoguer_theme = dialoguer::theme::ColorfulTheme::default();
    let mut html = String::new();
    std::io::stdin().read_to_string(&mut html).unwrap();
    let html = Html::parse_fragment(&html);
    let selector = Selector::parse("tr").unwrap();
    let summary = html
        .select(&selector)
        .skip(6)
        .map(|element| element.text().map(|v| v.trim()).filter(|v| !v.is_empty()))
        .map(|mut v| {
            let count = v.clone().count();
            let date = v.next().unwrap();
            let location = format!("{}", v.next().unwrap());
            // get last but one
            let price = v.skip(count - 3).next().unwrap();
            let price = price
                .split(' ')
                .next()
                .unwrap()
                .replace(",", ".")
                .parse()
                .expect(&format!("invalid price: {}", price));
            Transaction {
                date: date.to_string(),
                location: location.to_string(),
                price,
            }
        })
        .filter(|transaction| transaction.price < 0.0) // expenses only
        .map(|mut transaction| {
            transaction.price = transaction.price.abs();
            transaction
        })
        .fold(Summary::default(), |mut acc: Summary, curr: Transaction| {
            println!("Location: {}", curr.location);
            println!("Price: {}", curr.price);
            println!("Date: {}", curr.date);
            println!("");

            let selection = dialoguer::Select::with_theme(&dialoguer_theme)
                .items(&["mom", "me", "shared"])
                .default(0)
                .interact()
                .unwrap();

            let per_summary_ref = match selection {
                0 => &mut acc.mom,
                1 => &mut acc.me,
                2 => &mut acc.shared,
                _ => unreachable!(),
            };
            per_summary_ref.total += curr.price;
            per_summary_ref.transactions.push(curr);

            acc
        });

    println!(
        "mom: {}zl, me: {}zl, shared: {}zl",
        summary.mom.total, summary.me.total, summary.shared.total
    );
    let mut file = File::create(FILE_PATH).expect("create file");
    let summary_json = serde_json::to_string_pretty(&summary).unwrap();
    file.write_all(summary_json.as_bytes()).expect("write all");
    println!("wrote output to {}", FILE_PATH);
}
