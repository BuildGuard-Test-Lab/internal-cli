# Internal CLI

Command-line tool for internal developer workflows — deployments, database migrations, and service management.

## Installation
```bash
go install github.com/BuildGuard-Test-Lab/internal-cli@latest
```

## Commands
```bash
icli deploy <service> <env>
icli db migrate <service>
icli logs <service> --tail
```
