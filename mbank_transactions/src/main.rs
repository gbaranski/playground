use scraper::{Html, Selector};
use serde::{Deserialize, Serialize};
use std::fs::File;
use std::io::Read;

const FILE_PATH: &str = "transactions.csv";

#[derive(Debug, Serialize, Deserialize)]
struct Transaction {
    date: String,
    location: String,
    price: f64,
}

fn main() {
    let mut html = String::new();
    std::io::stdin().read_to_string(&mut html).unwrap();
    let html = Html::parse_fragment(&html);
    let selector = Selector::parse("tr").unwrap();
    let transactions = html
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
        });

        let file = File::create(FILE_PATH).unwrap();
        let mut wtr = csv::Writer::from_writer(file);
        for transaction in transactions {
            wtr.serialize(transaction).unwrap();
        }
        println!("wrote transactions to {}", FILE_PATH);
}
