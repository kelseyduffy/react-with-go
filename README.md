# react-with-go
Code from following along with [Working with React and Go](https://www.udemy.com/course/working-with-react-and-go-golang) on Udemy


### PostgreSQL database setup

- Install docker
- Create a PostgreSQL container (running locally on 5432 on this example)
  - `docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=<password> -d postgres`
- Create a database (named 'go_movies' in this example)
- Run the provided go_movies.sql script against that container
  - `psql -h 127.0.0.1 -p 5432 -d go_movies -U postgres -W -f go_movies.sql`