package structs

// for searching the track
type Station struct {
	Name string
	//Distance    int // needed for dijkstra. do not need for depth-first search
	X           int
	Y           int
	Vistited    bool
	Connections []*Station // points to a connected stations
}

/* did not use in the new and faster printing

// for printing functions
type Track struct {
	Name  []string
	Train []int
	InUse bool // need it to check if longer routes are used
}
*/
