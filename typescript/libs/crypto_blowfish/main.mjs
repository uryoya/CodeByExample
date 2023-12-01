/**
 * Blowfish暗号はOpenSSLv3ではデフォルトでサポートされていないため、NODE_OPTIONSに--openssl-legacy-providerを指定して実行する必要がある。
 * そのうえでのサンプル。
 * https://nodejs.org/api/crypto.html
 */
import crypto from "crypto";

const pw = "password";
const cipher = crypto.createCipheriv("bf-ecb", pw, null);
const decipher = crypto.createDecipheriv("bf-ecb", pw, null);

const text = "このテキストが暗号化される";
console.log(`元の文字列: ${text}`);

let es = cipher.update(text, "utf8", "hex");
es += cipher.final("hex");
console.log(`暗号化した文字列: ${es}`);

let ds = decipher.update(es, "hex", "utf8");
ds += decipher.final("utf8");
console.log(`復号した文字列: ${ds}`);
