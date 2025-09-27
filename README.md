# bible-es

Programa para crear una base de datos de PostgreSQL con la Biblia en diferentes versiones, organizadas en **libros**, **capítulos** y **versículos**.

## Dependencias

- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [sqlc](https://sqlc.dev/)
- [goose](https://github.com/pressly/goose)

## Variables de Entorno

```
DATABASE_URL="postgres://admin:admin@localhost:5432/db_name"
DATA_SOURCE_PATH="./data/json" # directorio con los datos en formato json

GOOSE_DRIVER=postgres
GOOSE_DBSTRING="postgres://admin:admin@localhost:5432/db_name"
GOOSE_MIGRATION_DIR=./sql/schema # directorio con la definicion de las tablas sql
```

## Migraciones

```bash
goose up
```

Para mas informacion de las migraciones echa un vistazo a la documentacion de [goose](https://github.com/pressly/goose)

## Compilacion y Ejecucion

- `go build -o bible-db .`
- `./bible-db`

## Metodo Rapido

Puedes descargar y/o usar el siguiente script de postgreSQL para crear la base de datos con los datos:

[Descargar script de PostgreSQL](https://github.com/hiahir357/bible-to-db-es/blob/main/bible_es)
