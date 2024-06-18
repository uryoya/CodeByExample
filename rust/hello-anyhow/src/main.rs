// anyhow
// https://docs.rs/anyhow/latest/anyhow/
// https://github.com/dtolnay/anyhow
use anyhow::{Context, Result};

// 普通のResultを使うとErrorの型を指定する必要がある。
fn parse_string(str1: &str, str2: &str) -> std::result::Result<i32, std::num::ParseIntError> {
    let int1 = str1.parse::<i32>()?;
    let int2 = str2.parse::<i32>()?;
    Ok(int1 + int2)
}

// anyhowを使うとErrorの型を指定しなくてもanyhow::Result<T, anyhow::Error>になる。
// std::error::Errorを実装している型なら何でも返せる。
fn parse_string_anyhow(str1: &str, str2: &str) -> Result<i32> {
    let int1 = str1.parse::<i32>()?;
    let int2 = str2.parse::<i32>()?;
    Ok(int1 + int2)
}

// anyhow::Contextを使うとエラーメッセージを追加でき、エラーの原因を特定しやすくなる。
// 元のエラーが低レベルなエラーの場合はcontext()で文脈を追加すると良い。
fn parse_string_anyhow_context(str1: &str, str2: &str) -> Result<i32> {
    let int1 = str1.parse::<i32>().context("failed to parse str1")?;
    let int2 = str2.parse::<i32>().context("failed to parse str2")?;
    Ok(int1 + int2)
}

fn main() {
    let result = parse_string("1", "a");
    let result_anyhow = parse_string_anyhow("2", "b");
    let result_anyhow_context = parse_string_anyhow_context("3", "c");

    println!("std::result::Result => {:?}", result);
    // ↑
    // std::result::Result => Err(ParseIntError { kind: InvalidDigit })

    println!("anyhow::Result => {:?}", result_anyhow);
    // ↑
    // anyhow::Result => Err(invalid digit found in string)

    println!("anyhow::Result with Context => {:?}", result_anyhow_context);
    // ↑
    // anyhow::Result with Context => Err(failed to parse str2
    //
    // Caused by:
    // invalid digit found in string)
}
