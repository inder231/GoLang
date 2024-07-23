# What's a REST api

` Server <-----> Client `

### Project Description

- A Go-powered "Event Booking" REST API

1. `GET : /events`                  ------> Get a list of available events
2. `GET : /events/<id>`             ------> Get an available event
3. `POST : /events`                 ------> Create a new bookable event    `<Auth Required>`
4. `PUT : /events/<id>`             ------> Update an event                `<Auth Required> <Only by creator>`
5. `DELETE : /events/<id>`          ------> Delete an event                `<Auth Required> <Only by creator>`
6. `POST : /signup`                 ------> Create a new user
7. `POST : /login`                  ------> Authenticate user             `<Auth Token JWT>`
8. `POST : /events/<id>/register`   ------> Register user for event       `<Auth Required>`
9. `DELETE : /events/<id>/register` ------> Cancel registration           `<Auth Required>`

##### Packages

1. Gin
- `go get -u github.com/gin-gonic/gin`
2. go-sqlite3
- `go get github.com/mattn/go-sqlite3`
3. bcrypt packate for password hashing
- `go get -u golang.org/x/crypto`
4. JWT ( JsonWebToken ) for login token generation
- `go get -u github.com/golang-jwt/jwt/v5`

### Basic server setup

```
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	/* Create Instance of a server */
	server := gin.Default()

	// Setting handler to handle http get request
	server.GET("/", checkServerStatus)

	// Run server on required port 
	server.Run(":8080") // localhost:8080

}

func checkServerStatus ( c *gin.Context) {
	c.JSON(http.StatusOK, "Server is Up!")
}
```

* Import statements which are required but not used and should not be removed on saving the file -> we should add a _ before their import to tell go to avoid removing it from the final saved file

`
Preparing Statements vs Directly Executing Queries (Prepare() vs Exec()/Query())
In the previous lectures, we started sending SQL commands to the SQLite database.

And we did this by following different approaches:

DB.Exec() (when we created the tables)

Prepare() + stmt.Exec() (when we inserted data into the database)

DB.Query() (when we fetched data from the database)

Using Prepare() is 100% optional! You could send all your commands directly via Exec() or Query().

The difference between those two methods then just is whether you're fetching data from the database (=> use Query()) or your manipulating the database / data in the database (=> use Exec()).

But what's the advantage of using Prepare()?

Prepare() prepares a SQL statement - this can lead to better performance if the same statement is executed multiple times (potentially with different data for its placeholders).

This is only true, if the prepared statement is not closed (stmt.Close()) in between those executions. In that case, there wouldn't be any advantages.

And, indeed, in this application, we are calling stmt.Close() directly after calling stmt.Exec(). So here, it really wouldn't matter which approach you're using.

But in order to show you the different ways of using the sql package, I decided to also include this preparation approach in this course.

`