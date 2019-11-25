package common

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
)

var (
    RedisPool *redis.Pool
)

func InitRedisPool(network string, host string, password string, database int) (pool *redis.Pool) {
    pool = &redis.Pool{
        MaxIdle:     16,
        MaxActive:   0,
        IdleTimeout: 300,
        Dial: func() (redis.Conn, error) {
            conn, err := redis.Dial(network, host)
            if err != nil {
                return nil, err
            }
            if password != "" {

                if _, err := conn.Do("AUTH", password); err != nil {
                    conn.Close()
                    return nil, err
                }
            }
            if _, err := conn.Do("SELECT", database); err != nil {
                panic(err)
                return nil, err
            }
            return conn, nil
        },
    }
    RedisPool = pool
    return
}

func RedisSet(redisConn redis.Conn, key string, value string) (code int, err error) {
    //if parent := opentracing.SpanFromContext(ctx); parent != nil {
    //    pctx := parent.Context()
    //    if tracer := opentracing.GlobalTracer(); tracer != nil {
    //        redisSpan := tracer.StartSpan("RedisSpan", opentracing.ChildOf(pctx))
    //        defer redisSpan.Finish()
    //    }
    //}

    if _, err := redisConn.Do("Set", key, value); err != nil {
        panic(err)
    }
    return
}

func RedisGet(redisConn redis.Conn, key string) (value string, err error) {
    fmt.Println(key)
    if value, err = redis.String(redisConn.Do("Get", key)); err != nil {
        value = "0"
    }
    return
}
