package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main(){
  s := rand.NewSource(time.Now().UnixNano())
  r := rand.New(s)
  for i := 0; i < 10; i++ {
    fmt.Println(r.Intn(10))
  }
  animals := []string{"Monkey", "Fish", "Tiger"}
  fmt.Println(animals[r.Intn(len(animals))])
}
