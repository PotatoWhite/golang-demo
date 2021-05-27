# Golang, GORM 을 이용한 Restful API 예제

# Golang init
```shell
go mod init github.com/potatowhite/golang-demo
```

# module Import
```shell
go get -u github.com/gorilla/mux
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```


## DBURL
```go
const DB_URL = "potato:potato@tcp(127.0.0.1:3306)/demo_golang?charset=utf8mb4&parseTime=True&loc=Local"
```