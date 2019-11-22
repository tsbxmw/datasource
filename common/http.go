package common

import (
    "github.com/jinzhu/gorm"
)

type (
    HttpServer interface {
        Serve()
        Init(config *ServiceConfig) HttpServer
    }

    HttpServerImpl struct {
        DB            *gorm.DB
        SvcName       string
        Address       string
        Port          int
        GrpcPort      string
        DbUri         string
        ConsulAddr    string
        ConsulPort    int
        JaegerAddr    string
        RedisHost     string
        RedisPort     string
        RedisPassword string
        RedisDB       int
    }
)
