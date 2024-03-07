package logic

import (
	"fmt"
	"os"
	"sort"
	"stations/structs"
)

// function to find all routes to end route
func FindAllRoutes(mainMap map[string]*structs.Station, start, end string) [][]string {
	allRoutes := [][]string{}
	// This is a good and fast way to check, if node is visited. Yay for maps
	visited := make(map[string]bool)
	var currentRoute []string
	//dfs stands for depth-first search. it is good for finding all routes for better info: https://en.wikipedia.org/wiki/Depth-first_search
	dfs(mainMap, visited, &allRoutes, currentRoute, start, end)
	if len(allRoutes) == 0 {
		fmt.Fprintf(os.Stderr, "..... No routes found from %v, to %v: \n", start, end)
		os.Exit(0)
	} else {
		allRoutes = removeAndOrder(allRoutes)
		fmt.Println("..... All non overlapping routes found:\n", allRoutes)
	}
	return allRoutes
}

// problem with dfs is it does not protect from loops. If there are any, we need to write some protection to it
func dfs(mainMap map[string]*structs.Station, visited map[string]bool, allRoutes *[][]string, currentRoute []string, current, end string) {
	// check if we reached end for recursion
	if current == end {
		// End station reached, add route to allRoutes
		route := append(currentRoute, current) // Add end station
		*allRoutes = append(*allRoutes, route)
		return
	}

	// Hope that this helps against loops
	if visited[current] {
		return // Avoid cycles
	}

	// Mark the current station as visited
	visited[current] = true
	// Add current station to the route
	currentRoute = append(currentRoute, current)

	// Explore each connection recursively. this is where the tree brances.
	//we go through as many times as there are brances. that we have not visited
	for _, station := range mainMap[current].Connections {
		if !visited[station.Name] {
			// Make a copy of the route for the recursive call
			// had some problem here. using array because faster and friend chat recommended it.
			newRoute := make([]string, len(currentRoute))
			copy(newRoute, currentRoute)
			dfs(mainMap, visited, allRoutes, newRoute, station.Name, end)
		}
	}

	// Unmark the current station as visited before backtracking
	visited[current] = false
}

func removeAndOrder(matrix [][]string) [][]string {
	newList := [][]string{}
	for _, a := range matrix {
		unique := true
		if newList == nil {
			newList = append(newList, a)
		} else {
		out:
			for i := 1; i < len(a)-1; i++ {
				for j, b := range newList {
					for _, c := range b {
						if a[i] == c {
							unique = false
							if len(b) > len(a) {
								newList[j] = a
							}
							break out
						}
					}
				}
			}
		}
		if unique {
			newList = append(newList, a)
		}
	}
	sort.Slice(newList, func(i, j int) bool {
		return len(newList[i]) < len(newList[j])
	})
	return newList
}

/*
*************************************************************************'
This is a older version with dijkstra. but there is a small problem with it.
dijkstra is what is called a greedy algorithm. it wants to find the shortest route to end
but in this exercise we need to find many routes. so using Depth-First Search might be better.
*********************************************************************************


func Createmap(mainMap map[string]*structs.Station, start string, end string) [][]string {
	unvisitStation, nextStation := []string{}, []string{}
	exit := true
	distance := 0
	// marking the first station to be the startstation
	currentStation := mainMap[start]
	currentStation.Distance = distance
	currentStation.Vistited = true
	unvisitStation = append(unvisitStation, currentStation.Connections...)

	fmt.Println("------------>", mainMap[start])
	for exit {
		// we dont use distance. all nodes are at equal distance, so no need for distance calculations
		distance++
		//we loop through all stations in the unvisited space
		for _, station := range unvisitStation {

			currentStation = mainMap[station]
			// we skip allready visited stations
			if currentStation.Vistited {
				continue
			}
			//we append all new connections to new list
			currentStation.Distance = distance
			currentStation.Vistited = true
			fmt.Println("------------>", mainMap[station])
			nextStation = append(nextStation, currentStation.Connections...)
			// if we find th final station we break
			if currentStation == mainMap[end] {
				fmt.Println("found end station........")
				exit = false
				break
				// also when we cannot reach final destination and have no new nodes. we break
			} else if len(unvisitStation) == 0 {
				fmt.Println("No route found to end station")
				exit = false
				break
			}
		}
		unvisitStation = nextStation
		exit = false
	}
	mainRoute := findShort(mainMap, start, end)
	return mainRoute

}

// we backtrack here and find the quickest routes from end to beginning.
func findShort(mainMap map[string]*structs.Station, start string, end string) [][]string {
	mainroute := [][]string{}
	tempRoute := []string{end}
	first := false

	//we declare the first node and no need to check the end node again
	currentStation := mainMap[end]
	currentStation.Vistited = false
	//we run the program as many times as there are ways to reach the end
	for _, station := range currentStation.Connections {
		tempsStation := mainMap[station]
		tempsStation.Vistited = false

		tempRoute = append(tempRoute, tempsStation.Name)
		compareStation := structs.Station{}
	out:
		for {
			// we check if the loop starts either from the before end node or
			//from deeper inside.
			if compareStation.Connections != nil {
				tempsStation = mainMap[compareStation.Name]
				tempsStation.Vistited = false
			}
			// we check for the current struct connections
			for _, station2 := range tempsStation.Connections {
				// first endpoint when there is no connection to first(did not trigger first)
				if first {
					fmt.Println("Did not find this route. stuck at: ", tempsStation.Name)
					tempRoute = []string{end}
					first = false
					break out
				}
				first = true
				// if the struct has been allready checked or not reached in mapping, it is ignored
				tempsStation2 := *mainMap[station2]
				if !tempsStation2.Vistited || tempsStation2.Distance == 2147483647 {
					continue
					// second endpoint when start is reached
				} else if mainMap[station2] == mainMap[start] {
					tempRoute = append(tempRoute, tempsStation2.Name)
					mainroute = append(mainroute, tempRoute)
					tempRoute = []string{end}
					first = false
					break out
					// when find first node assign to it
				} else if first {
					compareStation = tempsStation2
					first = false
					// when other node in the logic has shorter distance
				} else if tempsStation2.Distance < compareStation.Distance {
					compareStation = tempsStation2
				}
			}
		}

	}

	return mainroute
}
*/
