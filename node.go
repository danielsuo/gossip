package gossip

import "fmt"
// import "sync"
import "bytes"
import "math/rand"

var nodeCounter = 0

type Data int

type Gossiper interface {
	Start()
  Stop()
	Push(data Data)
  Gossip(data Data) bool

	Id() int

	SelectPeers() []Gossiper
	AddPeer(peer Gossiper)

	// TODO: Should be able to remove
	SetState(state State)
	Infected() bool
	Susceptible() bool
	Removed() bool
}

type Node struct {
	id    int
	peers []Gossiper
	data  Data
	state State
	inbox chan Data
}

func NewNode() *Node {
	n := Node{
		id:    nodeCounter,
		state: Susceptible,
		inbox: make(chan Data)}
	nodeCounter++
	return &n
}

func (node *Node) Start() {
  fmt.Printf("Starting node %d\n", node.id)
  for {
    a := <-node.inbox
    if !node.Gossip(a) {
      break
    }
  }
}

func (node *Node) Stop() {
  node.Push(-1)
}

func (node *Node) Push(data Data) {
	node.inbox <- data
}

func (node *Node) Gossip(data Data) bool {
  if data == -1 {
    return false
  }
  fmt.Printf("Received %d at %d\n", data, node.id)

  node.SetState(Infected)
  peers := node.SelectPeers()

  for _, peer := range peers {
    if peer.Susceptible() {
      peer.Push(data)
    }
  }
  return true
}

// TODO: This should really be same type, not Gossiper
func (node *Node) AddPeer(peer Gossiper) {
	node.peers = append(node.peers, peer)
}

func (node *Node) SelectPeers() []Gossiper {
	// NOTE: In default implementation, choose one peer uniformly at random
	peerId := rand.Intn(len(node.peers))
	return []Gossiper{node.peers[peerId]}
}

func (node *Node) Id() int {
	return node.id
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
	fmt.Fprintf(buffer, "Node #%d\n", node.Id())

	for _, peer := range node.peers {
		fmt.Fprintf(buffer, "\tPeer #%d\n", peer.Id())
	}

	return buffer.String()
}
