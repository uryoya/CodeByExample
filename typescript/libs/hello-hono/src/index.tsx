import { serve } from "@hono/node-server";
import { Hono } from "hono";

const app = new Hono();

/**
 * Hello World
 * https://hono.dev/getting-started/basic#hello-world
 */
app.get("/", (c) => {
  return c.text("Hello Hono!");
});

/**
 * Return JSON
 * https://hono.dev/getting-started/basic#return-json
 */
app.get("/api/hello", (c) => {
  return c.json({
    ok: true,
    message: "Hello Hono!",
  });
});

/**
 * Request and Response
 * https://hono.dev/getting-started/basic#request-and-response
 */
app.get("/posts/:id", (c) => {
  const page = c.req.query("page");
  const id = c.req.param("id");
  c.header("X-Message", "Hi!");
  return c.text(`You want see ${page} of ${id}`);
});

app.post("/posts", (c) => c.text("Created!", 201));
app.delete(`/posts/:id`, (c) => c.text(`${c.req.param("id")} is deleted!`));

/**
 * Return raw Response
 * https://hono.dev/getting-started/basic#return-raw-response
 */
app.get("/raw-response", (c) => new Response("Good morning!"));

/**
 * Return HTML
 * https://hono.dev/getting-started/basic#return-html
 */
const View = () => (
  <>
    <html>
      <body>
        <h1>Hello Hono!</h1>
      </body>
    </html>
  </>
);

app.get("/page", (c) => c.html(<View />));

/**
 * Using Middleware
 * https://hono.dev/getting-started/basic#using-middleware
 */
import { basicAuth } from "hono/basic-auth";

app.use(
  "/admin/*",
  basicAuth({
    username: "admin",
    password: "secret",
  })
);

app.get("/admin", (c) => c.text("Welcome to admin page!"));

/**
 * Adapter
 * https://hono.dev/getting-started/basic#adapter
 */
// import { serveStatic } from "hono/cloudflare-workers";

// app.get("/static/*", serveStatic({ root: "./" }));

const port = 3000;
console.log(`Server is running on port ${port}`);

serve({
  fetch: app.fetch,
  port,
});
