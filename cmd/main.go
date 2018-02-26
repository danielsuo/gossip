package main

import "fmt"
import "github.com/danielsuo/gossip"

func main() {
  net := gossip.NewNetwork(6)

  fmt.Println(net)
}
