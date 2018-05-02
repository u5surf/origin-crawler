package main

import (
  "os"
  "os/exec"
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "github.com/gomodule/redigo/redis"
)

const origin = "https://anond.hatelabo.jp"
const proxy = "https://153.126.171.175:8443"

func CrawulOrigin() {
  doc,_ := goquery.NewDocument(origin)
  doc.Find("#hotentriesblock > ul > li").Each(func(_ int, s *goquery.Selection) {
    path,_ := s.Find("a").Attr("href")
    PushRedis(path)
    GetProxy(path)
  })
}

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

func GetProxy(path string) {
  out, err := exec.Command("curl", "-k", "-v", proxy+path, "-o", "/dev/null", "-s").CombinedOutput()
  fmt.Println(string(out))
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func getOriginStatus(path string) string {
  out, err := exec.Command("curl", "-k", origin+path, "-o", "/dev/null", "-w", "'%{http_code}'","-s").Output()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return string(out)
}

func main() {
  CrawulOrigin()
  for _, path := range PopRedis() {
    status := getOriginStatus(path)
    fmt.Println(status)
    if status == "404" {
      fmt.Println(path)
    }
  }
}

