package main

import "fmt"
import "github.com/danielsuo/gossip"

func main() {
  net := gossip.NewNetwork(20)

  fmt.Println(net)
  // net.Gossip(1)
  net.Start()

}
