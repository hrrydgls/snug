# Snug

Is gonna be an open source renting platform API.

## Dev env

I am using a docker compose file for deploy a db and expose port to my host machine.

## Packages I am using

- Migrations: [https://github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- ORM: [https://github.com/go-gorm/gorm](https://github.com/go-gorm/gorm)
- Env vars: [https://github.com/joho/godotenv](https://github.com/joho/godotenv)

## Migrations

### Create a new migration

```bash
> make migrate-new                      
Enter name: create_languages_table
/home/hamid/Codes/snug/db/migrations/000002_create_languages_table.up.sql
/home/hamid/Codes/snug/db/migrations/000002_create_languages_table.down.sql
```
