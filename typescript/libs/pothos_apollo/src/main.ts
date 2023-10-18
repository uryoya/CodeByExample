import SchemaBuilder from "@pothos/core";
import { ApolloServer } from "@apollo/server";
import { startStandaloneServer } from "@apollo/server/standalone";

const builder = new SchemaBuilder({});

builder.queryType({
  fields: (t) => ({
    hello: t.string({
      args: {
        name: t.arg.string(),
      },
      resolve: (parent, { name }) => `hello, ${name || "World"}`,
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
