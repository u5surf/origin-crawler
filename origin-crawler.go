package main

import (
  "os"
  "os/exec"
  "fmt"
  "github.com/PuerkitoBio/goquery"
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

func GetProxy(path string) {
  out, err := exec.Command("curl", "-k", "-v", proxy+path, "-o", "/dev/null", "-s").CombinedOutput()
  fmt.Println(string(out))
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func main() {
  CrawulOrigin()
}

