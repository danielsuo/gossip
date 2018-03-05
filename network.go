package gossip

import "fmt"
import "time"
import "sync"
import "bytes"
import "math/rand"

type Network struct {
	nodes []Gossiper
  
}

func NewNetwork(numNodes int) *Network {
	// TODO: Network should be able to change size
	n := Network{nodes: make([]Gossiper, numNodes, numNodes)}
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

func (network *Network) Start() {
  wg := new(sync.WaitGroup)
	for _, node := range network.nodes {
    wg.Add(1)
    go func(node Gossiper) {
      defer wg.Done()
      node.Start()
    }(node)
  }
  go network.nodes[0].Push(1)
  time.Sleep(100 * time.Millisecond)
  for _, node := range network.nodes {
    node.Stop()
  }

  wg.Wait()
}

func (network *Network) Gossip(value int) {
	patientZero := rand.Intn(len(network.nodes))
	fmt.Printf("Patient Zero is %d\n", patientZero)

	network.nodes[patientZero].SetState(Infected)

	rounds := 0

	for {
		if network.Infected() {
			fmt.Printf("Infected in %d rounds\n", rounds)
			break
		}

		rounds++

		for i := 0; i < len(network.nodes); i++ {
			if network.nodes[i].Infected() {
				// TODO: Infected should not propagate within a single round
				fmt.Printf("Round %d, infected %d\n", rounds, i)
				peers := network.nodes[i].SelectPeers()
				for _, peer := range peers {
					if peer.Susceptible() {
						peer.SetState(Infected)
					}
				}
			}
		}
	}
}

func (network *Network) Infected() bool {
	for i := 0; i < len(network.nodes); i++ {
		if !network.nodes[i].Infected() {
			return false
		}
	}

	return true
}

func (network *Network) String() string {
	buffer := bytes.NewBufferString("")

	for _, node := range network.nodes {
		fmt.Fprintf(buffer, "%s\n", node)
	}

	return buffer.String()
}
