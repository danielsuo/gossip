package gossip

import "fmt"
import "bytes"

var nodeCounter = 0

type Node struct {
	Id    int
	peers []*Node
}

func NewNode() *Node {
	n := Node{Id: nodeCounter}
	nodeCounter++
	return &n
}

func (node *Node) AddPeer(peer *Node) {
  node.peers  = append(node.peers, peer)
}

func (node *Node) String() string {
  buffer := bytes.NewBufferString("")
  fmt.Fprintf(buffer, "Node #%d\n", node.Id)

  for _, peer := range node.peers {
    fmt.Fprintf(buffer, "\tPeer #%d\n", peer.Id)
  }

  return buffer.String()
}
