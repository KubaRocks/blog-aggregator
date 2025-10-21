# Gator

- Ensure Postgres and Go are installed before running the program.
- Install the gator CLI using `go install .`.
- Set up your config file by creating and editing `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://user:password@localhost:5432/database?sslmode=disable",
}
```

- Run the program with `gator` and explore commands like `gator agg`.
