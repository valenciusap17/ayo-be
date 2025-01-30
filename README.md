## Getting started
Install Go version 1.23.5 and run 
```bash
go mod tidy
go mod vendor
```

## Add schema migrations

- Create migrations: `make create-migration name={migration_file_name}`
- Move migration to migrations directory (Manually): `mv <"FILE_NAME">.sql migrations/<"FILE_NAME">.sql`

## Execute migrations

- Apply all available migrations: `make migrate-up`
- Apply a single migration from current version: `make migrate-up-by-one`
- Role back single migration from the current version: `make migrate-down`
