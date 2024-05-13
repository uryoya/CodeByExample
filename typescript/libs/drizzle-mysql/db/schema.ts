import { json, mysqlTable, varchar } from 'drizzle-orm/mysql-core';

export const json_tables = mysqlTable('json_tables', {
  id: varchar('id', { length: 36 }).primaryKey(),
  json: json('json'),
});
