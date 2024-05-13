import { defineConfig } from 'drizzle-kit';

export default defineConfig({
  schema: './db/*.ts',
  dialect: 'mysql',
  dbCredentials: {
    url: process.env.DATABASE_URL,
  },
  verbose: true,
  strict: true,
});
