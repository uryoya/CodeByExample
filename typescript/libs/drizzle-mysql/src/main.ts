import { sql } from 'drizzle-orm';
import { drizzle } from 'drizzle-orm/mysql2';
import mysql from 'mysql2/promise';
import { json_tables } from '../db/schema';
import { randomUUID } from 'crypto';

const main = async () => {
  const client = await mysql.createConnection({
    uri: process.env.DATABASE_URL,
  });
  const db = drizzle(client);

  // レコードの挿入
  await db.insert(json_tables).values({
    id: randomUUID(),
    json: {
      foo: 'bar',
      baz: Math.floor(Math.random() * 100),
      array: [1, 2, 3],
    },
  });

  // 全件取得
  const jsons = await db
    .select({
      id: json_tables.id,
      json: json_tables.json,
    })
    .from(json_tables);

  console.log('all jsons:');
  console.log(jsons);

  // JSONの中身でフィルタリングして取得
  // JSON用の構文はdrizzleでサポートされていなさそうなのでsqlオペレーターを使って取得する
  // https://orm.drizzle.team/docs/sql
  const [filteredJsons] = await db.execute(
    sql`SELECT * FROM ${json_tables} WHERE json->>'$.baz' > 50`,
  );

  console.log('filtered jsons:');
  console.log(filteredJsons);

  process.exit(0);
};

try {
  main();
} catch (e) {
  console.error(e);
  process.exit(1);
}
