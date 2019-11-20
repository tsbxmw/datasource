package main

import (
    "datasource/public"
    "fmt"
    "os"
)

func main() {
    app, err := public.App("datasource", "usage for datasource")
    if err!=nil{
        panic(err)
    }
    fmt.Println("start main")
    app.Run(os.Args)
}