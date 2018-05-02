package main

import (
  "os"
  "os/exec"
  "fmt"
)

const origin = "https://anond.hatelabo.jp"
const proxy = "https://153.126.171.175:8443"


func getOriginStatus(path string) string {
  out, err := exec.Command("curl", "-k", origin+path, "-o", "/dev/null", "-w", "'%{http_code}'","-s").Output()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return string(out)
}

func main() {
  for _, path := range PopRedis() {
    status := getOriginStatus(path)
    fmt.Println(status)
    if status == "404" {
      fmt.Println(path)
    }
  }
}

