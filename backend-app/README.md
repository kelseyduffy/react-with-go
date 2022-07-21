### PostgreSQL database setup

- Install docker
- Create a PostgreSQL container (running locally on 5432 on this example)
  - `docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=<password> -d postgres`
- Create a database (named 'go_movies' in this example)
- Run the provided go_movies.sql script against that container
  - `psql -h 127.0.0.1 -p 5432 -d go_movies -U postgres -W -f go_movies.sql`


### Running application
- to use default values of each item, just run `go run cmd/api/*.go` from the backend-app folder
- to provide a new value for an item, use `go run cmd/api/*.go -<flag> <value>`
  - This is necessary for the postgres connection string because the default value has `<password>` for the password. It is to show the syntax without storing the password in version control. 
  - To run this successfully, copy the default string and replace the password placeholder with the actual password, and run `go run cmd/api/*.go -dsn <new dsn string>`