This is a small reproduction case for a bug in gqlgen 0.9.1, where null values are not bubbling up correctly through arrays of non-nullable elements.

To reproduce the bug, run:

```sh
go run server/server.go&
# wait until server is ready, then run:
curl -H 'Content-Type: application/json' -d '{"query": "{ todos { user { name } } }"}' http://localhost:8080/query
```

You should see:

```json
{"errors":[{"message":"oh no","path":["todos",0,"user"]}],"data":{"todos":[null]}}
```

The part to focus on is `"todos":[null]`, which should not be allowed according to the schema.
