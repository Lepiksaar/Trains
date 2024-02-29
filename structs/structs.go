package structs

type Station struct {
	Name        string
	Distance    int // can see if visited, by seeing that distance is not infinite
	X           int
	Y           int
	Vistited    bool     // added just in case maybe delete later
	Connections []string // actually it points to a map names
}
