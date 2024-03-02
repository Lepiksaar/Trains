package logic

import (
	"fmt"
	"stations/structs"
)

// here we map the distances with the reference of beginning point.
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
			// first endpoint when
			for _, station2 := range tempsStation.Connections {
				if first {
					fmt.Println("Did not find this route. stuck at: ", tempsStation.Name)
					tempRoute = []string{end}
					first = false
					break out
				}
				first = true
				tempsStation2 := *mainMap[station2]
				if !tempsStation2.Vistited || tempsStation2.Distance == 2147483647 {
					continue
				} else if mainMap[station2] == mainMap[start] {
					tempRoute = append(tempRoute, tempsStation2.Name)
					mainroute = append(mainroute, tempRoute)
					tempRoute = []string{end}
					first = false
					break out
				} else if first {
					compareStation = tempsStation2
					first = false
				} else if tempsStation2.Distance < compareStation.Distance {
					compareStation = tempsStation2
				}
			}
		}

	}

	return mainroute
}
