package main

import (
    "datasource/public"
    "fmt"
    "github.com/sirupsen/logrus"
    "os"
)

func main() {
    app, err := public.App("datasource", "usage for datasource")
    if err!=nil{
        panic(err)
    }
    fmt.Println("start main")
    if err:=app.Run(os.Args); err!=nil{
        logrus.Error(err)
        panic(err)
    }
}