package gossip

type State int

const (
	Infected    State = 0
	Susceptible State = 1
	Removed     State = 2
)

func (state State) String() string {
	names := [...]string{"Infected", "Susceptible", "Removed"}

	if state < Infected || state > Removed {
		return "Unknown"
	}

	return names[state]
}
