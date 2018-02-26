package gossip

import "fmt"
import "bytes"

type Network struct {
  nodes []*Node
}

func NewNetwork(numNodes int) *Network {
	// TODO: Network should be able to change size
	n := Network{nodes: make([]*Node, numNodes, numNodes)}
	for i := 0; i < numNodes; i++ {
		n.nodes[i] = NewNode()
	}

  // TODO: For now, all nodes peers with all other nodes
  for i := 0; i < numNodes; i++ {
    for j := 0; j < numNodes; j++ {
      if i != j {
        n.nodes[i].AddPeer(n.nodes[j])
      }
    }
  }

  return &n
}

// TODO: Strictly speaking, Network is the oracle; nodes should really only
// know about their peers and not have this global information
func (network *Network) Add(node *Node) {
	network.nodes = append(network.nodes, node)
}

func (network *Network) String() string {
  buffer := bytes.NewBufferString("")

  for _, node := range network.nodes {
    fmt.Fprintf(buffer, "%s\n", node)
 }

  return buffer.String()
}
