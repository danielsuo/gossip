package gossip

import "fmt"
import "bytes"

var nodeCounter = 0

type Data struct {
	Value int
	Round int
}

type Node struct {
	Id    int
	peers []*Node
	data  Data
	state State
}

func NewNode() *Node {
	n := Node{Id: nodeCounter, state: Susceptible}
	nodeCounter++
	return &n
}

func (node *Node) AddPeer(peer *Node) {
	node.peers = append(node.peers, peer)
}

func (node *Node) SetState(state State) {
	node.state = state
}

func (node *Node) Infected() bool {
	return node.state == Infected
}

func (node *Node) Susceptible() bool {
	return node.state == Susceptible
}

func (node *Node) Removed() bool {
	return node.state == Removed
}

func (node *Node) String() string {
	buffer := bytes.NewBufferString("")
	fmt.Fprintf(buffer, "Node #%d\n", node.Id)

	for _, peer := range node.peers {
		fmt.Fprintf(buffer, "\tPeer #%d\n", peer.Id)
	}

	return buffer.String()
}
