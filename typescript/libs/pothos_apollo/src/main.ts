import SchemaBuilder from "@pothos/core";
import TracingPlugin, {
  wrapResolver,
  isRootField,
} from "@pothos/plugin-tracing";
import { ApolloServer } from "@apollo/server";
import { startStandaloneServer } from "@apollo/server/standalone";

export class Post {
  id: string;
  title: string;
  content: string;

  constructor(id: string, title: string, content: string) {
    this.id = id;
    this.title = title;
    this.content = content;
  }
}

const builder = new SchemaBuilder({
  plugins: [TracingPlugin],
  tracing: {
    default: (config) => isRootField(config),
    wrap: (resolver, options, config) =>
      wrapResolver(resolver, (error, duration) => {
        console.log(
          `Executed resolver ${config.parentType}.${config.name} in ${duration}ms`,
        );
      }),
  },
});

builder.objectType(Post, {
  name: "Post",
  description: "A blog post",
  fields: (t) => ({
    id: t.exposeID("id"),
    title: t.exposeString("title", {
      description: "The title of the blog post",
    }),
    content: t.exposeString("content", {
      description: "The content of the blog post",
    }),
  }),
});

builder.queryType({
  fields: (t) => ({
    hello: t.string({
      args: {
        name: t.arg.string(),
      },
      resolve: (parent, { name }) => `hello, ${name || "World"}`,
    }),
    post: t.field({
      type: Post,
      resolve: () => new Post("1", "Hello World", "This is a blog post."),
    }),
  }),
});

builder.mutationType({
  fields: (t) => ({
    post: t.boolean({
      args: {
        message: t.arg.string(),
      },
      resolve: (root, args) => {
        // do something
        return true;
      },
    }),
  }),
});

const schema = builder.toSchema();

const server = new ApolloServer({
  schema,
});

startStandaloneServer(server, { listen: { port: 3000 } }).then(({ url }) =>
  console.log(`🚀 Server ready at ${url}`),
);
