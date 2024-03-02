package structs

type Station struct {
	Name string
	//Distance    int // needed for dijkstra. do not need for depth-first search
	//X           int // they gave us location. do not need it for exercise
	//Y           int // they gave us location. do not need it for exercise
	Vistited    bool
	Connections []*Station // points to a connected stations
}
