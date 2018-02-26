package main

import "fmt"
import "github.com/danielsuo/gossip"

func main() {
  net := gossip.NewNetwork(20)

  fmt.Println("Starting")
  net.Gossip(1)
}
