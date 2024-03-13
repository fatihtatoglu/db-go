# db-go

The Database Go Package is a collection of tools for working with databases in Go applications. It provides functionalities for managing database connections, executing commands, and handling rows and tables. With this package, developers can easily interact with databases in their Go projects, making database-related tasks more efficient and straightforward.

## Install

```bash
go get github.com/fatihtatoglu/db-go
```

## Usage

### Creating a Database Connection

```go
config := db.CreateNewDBConfig("mysql", "username", "password", "localhost", 3306, "database_name")
connection, err := db.CreateNewDBConnection(config)
if err != nil {
    // handle error
}

err = connection.Open()
if err != nil {
    // handle error
}

defer connection.Close()

```

### Executing Commands

```go
command, err := db.CreateNewDBCommand(connection)
if err != nil {
    // handle error
}

_, err = command.Execute("CREATE TABLE users (name VARCHAR(255), age INT)")
if err != nil {
    // handle error
}

```

### Executing Commands with Parameters

```go
result, err := db.Execute("INSERT INTO users (name, age) VALUES (?, ?)", "John Doe", 30)
 if err != nil {
  // handle error
 }
 fmt.Println("Rows affected:", result.RowsAffected)
```

### Querying Data

```go
query := "SELECT * FROM users"
result, err := command.Query(query)
if err != nil {
    // handle error
}

```

### Querying Data with Parameters

```go
query := "SELECT * FROM users WHERE age > ?"
result, err := command.Query(query, 25)
if err != nil {
    // handle error
}

```

## Features

- [X] Support for various database drivers including;
  - [X] MySQL
  - [ ] PostgreSQL
  - [ ] SQLite
  - [ ] SQL Server
- [ ] Easy-to-use API for executing SQL queries, fetching data, and managing database connections
- [ ] Efficient handling of database transactions and error handling

## Credits

- [Fatih Tatoğlu](https://github.com/fatihtatoglu)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
