package inout

func Printout(tracks [][]string, numTrains int) {

}

package main

import (
	"fmt"
	"your_package/structs" // Replace with your actual package path
)

// Station represents a station in the graph.
type Station struct {
	Name        string
	Visited     bool
	Connections []*Station
}

// FindAllRoutes finds all routes from start to end station.
func FindAllRoutes(mainMap map[string]*Station, start, end string) [][]string {
	allRoutes := [][]string{}
	visited := make(map[string]bool)
	var currentRoute []string
	dfsFindRoutes(mainMap, visited, &allRoutes, currentRoute, start, end)
	return allRoutes
}

// dfsFindRoutes is a recursive helper function for FindAllRoutes.
func dfsFindRoutes(mainMap map[string]*Station, visited map[string]bool, allRoutes *[][]string, currentRoute []string, current, end string) {
	if visited[current] {
		return // Avoid revisiting
	}
	if current == end {
		// End station reached, add route to allRoutes
		route := append([]string(nil), currentRoute...) // Make a copy of currentRoute
		route = append(route, current)                  // Add end station
		*allRoutes = append(*allRoutes, route)
		return
	}
	// Mark the current station as visited
	visited[current] = true
	currentRoute = append(currentRoute, current) // Add current station to the route

	// Explore each connection recursively
	for _, station := range mainMap[current].Connections {
		if !visited[station.Name] {
			dfsFindRoutes(mainMap, visited, allRoutes, currentRoute, station.Name, end)
		}
	}
	// Unmark the current station as visited before backtracking
	visited[current] = false
}

func main() {
	stations := make(map[string]*Station)
	// Populate your stations map with Station structs

	// Example usage of FindAllRoutes
	start := "a" // Replace with your actual start station name
	end := "b"   // Replace with your actual end station name
	allRoutes := FindAllRoutes(stations, start, end)

	fmt.Println("All possible routes from", start, "to", end, ":")
	for _, route := range allRoutes {
		fmt.Println(route)
	}
}