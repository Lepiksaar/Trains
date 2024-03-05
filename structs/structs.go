package structs

// for searching the track
type Station struct {
	Name string
	//Distance    int // needed for dijkstra. do not need for depth-first search
	//X           int // they gave us location. do not need it for exercise
	//Y           int // they gave us location. do not need it for exercise
	Vistited    bool
	Connections []*Station // points to a connected stations
}

// for printing functions
type Track struct {
	Name  []string
	Train []int
	InUse bool // need it to check if longer routes are used
}
