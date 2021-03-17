### Go Gin CRUD RESTful APIs Example

**About the tech stack**

    |Go Modules| is a dependency management system introduced since Go 1.1+ https://blog.golang.org/using-go-modules
    |Wire|       is a code generation tool providing compile-time dependency injection for Go https://github.com/google/wire
    |Gin|        is a popular web framework written in Go https://github.com/gin-gonic/gin
    |Gorm|       is an ORM library https://gorm.io/
    |GoDotEnv|   is a Go port of the Ruby dotenv project (which loads env vars from a .env file) https://github.com/joho/godotenv
    |go.uuid|    is a package provides pure Go implementation of Universally Unique Identifier (UUID) https://github.com/satori/go.uuid   

**Prerequisites**

    This example's tested on go1.16.2 and 10.4.17-MariaDB