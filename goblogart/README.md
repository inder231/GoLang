## Production ready project setup

- go mod init goblogart


#### Folder structure setup
- inits controllers middlwares migrations models

#### Libraries

1. CompileDaemon package for automatic builds
- go get github.com/githubnemo/CompileDaemon
- To install it RUN: go install github.com/githubnemo/CompileDaemon

2. env package for securing application secrets
- go get github.com/joho/godotenv

3. Gin framework
- go get -u github.com/gin-gonic/gin

4. Gorm package
- go get -u gorm.io/gorm

- We will need a database driver to work with Gorm.
5. Install mysql drivers
- go get -u gorm.io/driver/mysql

6. bcrypt package for hashing
- go get -u golang.org/x/crypto/bcrypt

7. jwt for generating tokens
- go get -u github.com/golang-jwt/jwt/v5


- Now create a main.go file in root of project

8. RUN: `CompileDaemon -command-"./goblogart"` to build project automatically every time we save a file. ( if didn't work run : `export PATH=$PATH:$(go env GOPATH)/bin`
 as well )