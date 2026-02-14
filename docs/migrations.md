# Quick: run migrations (dev)

1. Start the Mongo service (docker-compose):

```bash
docker compose up -d mongodb
```

1. Copy migrations into the running container (Git Bash / POSIX):

```bash
docker cp ./migrations/. ledger-mongodb:/tmp/migrations
```

1. Run migrations inside the container:

```bash
docker compose exec mongodb bash -lc 'for f in /tmp/migrations/*.js; do echo "Running $f"; mongosh "mongodb://root:rootpassword@localhost:27017/ledger?authSource=admin" --file "$f" || exit 1; done'
```

Notes

- `migrations/create_indexes.js` was adjusted to avoid creating an explicit `_id` unique index (MongoDB provides `_id` uniqueness automatically).
- If you have `mongosh` installed locally you can also run `./scripts/run-migrations.sh` from the repo root.
- To make the helper pick custom credentials/host, copy `env.template` to `.env` and edit values.

That's it — the container-based copy + in-container `mongosh` loop is the simplest option when `mongosh` is not installed locally.
