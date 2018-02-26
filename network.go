package gossip

import "fmt"
import "bytes"
import "math/rand"

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

func (network *Network) Gossip(value int) {
	patientZero := rand.Intn(len(network.nodes))
	fmt.Printf("Patient Zero is %d\n", patientZero)

	network.nodes[patientZero].SetState(Infected)

	fanout := 1
	rounds := 0

	for {
		if network.Infected() {
			fmt.Printf("Infected in %d rounds\n", rounds)
			break
		}

		rounds++

		for i := 0; i < len(network.nodes); i++ {
			if network.nodes[i].Infected() {
        fmt.Printf("Round %d, infected %d\n", rounds, i)
				for j := 0; j < fanout; j++ {
					// TODO: This should be list of peers
					target := rand.Intn(len(network.nodes))
					if network.nodes[target].Susceptible() {
						network.nodes[target].SetState(Infected)
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
