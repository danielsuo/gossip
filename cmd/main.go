package main

import "fmt"
import "github.com/danielsuo/gossip"

func main() {
  n := gossip.Node{1}
  fmt.Printf("Hello, %d!\n", n.Id)
}
