package main

import (
  "os"
  "fmt"
  "github.com/gomodule/redigo/redis"
)

func PushRedis(path string) {
  c, err := redis.Dial("tcp", "127.0.0.1:6379")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  defer c.Close()
  c.Send("SELECT",0)
  c.Send("SET", path, path)
  c.Flush()
}

func PopRedis() []string {
  c, err := redis.Dial("tcp", "127.0.0.1:6379")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  defer c.Close()
  c.Do("SELECT",0)
  paths , err := redis.Strings(c.Do("keys","*"))
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return paths
}

