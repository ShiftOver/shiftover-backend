# shiftover-backend

Monolith GO backend API codebase for ShiftOver

## Development

Prepare development environments
```bash
make prepare
```

Rehook pre-commit
```bash
make rehooks
```

To start the local development server
```bash
make dev.up
```

To stop the local development server
```bash
make dev.down
```

## Database Migrations

To create new migration
```bash
migrate-mongo create <migration-name>
```

To migrate database up (updating database)
```bash
make migrate.up
```

To migrate database down (rolling back database)
```bash
make migrate.down
```

To view migration status
```bash
migrate-mongo status
```
