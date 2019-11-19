package http

import "github.com/jinzhu/gorm"


type (
    HttpServer struct {
        DB            *gorm.DB
        SvcName       string
        Address       string
        Port          int
        GrpcPort      string
        DbUri         string
        ConsulAddr    string
        JaegerAddr    string
        RedisHost     string
        RedisPort     string
        RedisPassword string
        RedisDB       int
    }
)

func (h HttpServer) Serve() {

}