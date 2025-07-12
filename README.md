# Snug

Is gonna be an open source renting platform API.

## Dev env

I am using a docker compose file for deploy a db and expose port to my host machine.

## ✅ Standard Library

| Package                                             | Use Case                                                   |
| --------------------------------------------------- | ---------------------------------------------------------- |
| [`log`](https://pkg.go.dev/log)                     | Logging errors, warnings, and debug messages.              |
| [`os`](https://pkg.go.dev/os)                       | Reading environment variables and interacting with the OS. |
| [`time`](https://pkg.go.dev/time)                   | Handling timeouts, durations, and connection lifetimes.    |
| [`net/http`](https://pkg.go.dev/net/http)           | Building HTTP server, routes, and middleware.              |
| [`encoding/json`](https://pkg.go.dev/encoding/json) | Marshaling/unmarshaling JSON payloads.                     |
| [`errors`](https://pkg.go.dev/errors)               | Working with Go error types and error comparisons.         |
| [`errors`](https://pkg.go.dev/net/mail)             | Package mail implements parsing of mail messages.          |

---

## 📦 Third-Party Libraries

| Package                                                                                      | Use Case                                                                             |
| -------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------ |
| [`gorm`](https://gorm.io/)                                                                   | Object-Relational Mapping (ORM) for interacting with PostgreSQL and managing models. |
| [`gorm.io/driver/postgres`](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL) | GORM PostgreSQL driver for database connection.                                      |
| [`github.com/joho/godotenv`](https://github.com/joho/godotenv)                               | Loads `.env` files into environment variables during development.                    |

---

## 🏗️ Internal Project Packages

| Package                                     | Use Case                                                |
| ------------------------------------------- | ------------------------------------------------------- |
| `github.com/hrrydgls/snug/models`           | Application models used by GORM for database structure. |
| `github.com/hrrydgls/snug/models/responses` | Structs for standardizing API responses.                |
| `github.com/hrrydgls/snug/exceptions`       | Custom error types used in centralized error handling.  |
| `github.com/hrrydgls/snug/db`               | Database connection configuration and pooling.          |
| `github.com/hrrydgls/snug/routes`           | HTTP route definitions and handler registration.        |
| `github.com/hrrydgls/snug/middlewares`      | HTTP middleware such as panic recovery.                 |

## Migrations

### Create a new migration

```bash
> make migrate-new                      
Enter name: create_languages_table
/home/hamid/Codes/snug/db/migrations/000002_create_languages_table.up.sql
/home/hamid/Codes/snug/db/migrations/000002_create_languages_table.down.sql
```
