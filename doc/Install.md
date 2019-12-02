# How to Use DataSource


# Install these software first

## go

Go Version 1.13 or higher

https://golang.google.cn/

## Mysql

Mysql

https://www.mysql.com/

## Redis

https://redis.io/

## rabbit mq

https://www.rabbitmq.com/

## consul

https://www.consul.io/


# Get Code Source


## git clone the code source

```shell
git clone https://github.com/tsbxmw/datasource
```

## modify the config in auth/config and data/config


**change all configuration to yours!**

```json
{
  "address": "172.22.96.83",
  "consul_addr": "http://172.23.152.46",
  "consul_port": 8500,
  "db_uri": "root:mengwei@(tcloud.tsbx.com:3306)/datasource?parseTime=true&loc=Local",
  "jaeger_addr": "tcloud-api.tsbx.com:5775",
  "log_file": "log/auth.log",
  "port": 9040,
  "grpc_port": 9041,
  "retry_max": 3,
  "retry_timeout": 500,
  "service_name": "auth_v1",
  "redis_host": "tcloud.tsbx.com",
  "redis_password": "",
  "redis_port": "6379",
  "redis_db": 1
}
```

# Go Build the Code

## build the auth

```shell
cd datasource
go build -o auth1 apps/auth/auth.go
```

## build the data

```shell
cd datasource
go build -o data1 apps/data/data.go
```


# Run the code

## auth service

```shell
./auth1 --config=./auth/config/dev.json httpserver
```

## data service

```shell
./data1 --config=./data/config/dev.json httpserver
```

## data worker

```shell
./data1 --config=./data/config/dev.json worker-server
```