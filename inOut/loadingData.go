package inOut

import (
	"bufio"
	"fmt"
	"os"
	"stations/structs"
	"strings"
)

var mainStations = make(map[string]*structs.Station)

// function that reads the info into a struct
func FillStruct(filePath string) map[string]*structs.Station {

	mapFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer mapFile.Close()

	scanner := bufio.NewScanner(mapFile)
	place := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// fast checks to continue, if lines are not correct.
		if line == "stations:" || strings.HasPrefix(line, "#") || len(line) < 1 {
			continue
		}
		if line == "connections:" {
			//switch to building connections
			place = false
			continue
		}
		if place {
			// we run the function to recieve required struct
			tempStation := makeStation(line)
			// we add the required struct to a map
			mainStations[tempStation.Name] = &tempStation
		} else {
			addConnection(line)
		}
	}
	fmt.Println("\n..... Successfully loaded Stations and connections")
	return mainStations
}
func makeStation(line string) structs.Station {
	parts := strings.Split(line, ",")

	if len(parts) < 3 {
		fmt.Println("Not enough arguments to station format:", line)
		return structs.Station{}
	}
	name := strings.TrimSpace(parts[0])

	/*
		It seems we dont need coordinates

		//removing whitespace and ignoring evrithing after #
		uglyX := strings.TrimSpace(ignoreHash(&parts[1]))
		uglyY := strings.TrimSpace(ignoreHash(&parts[2]))
		//assigning cordinates
		x, err1 := strconv.Atoi(uglyX)
		y, err2 := strconv.Atoi(uglyY)
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid Coordinates: " + name)
			return structs.Station{}
		} */
	//we return the struct that we want to add to map
	return structs.Station{
		Name: name,
		//X:        x, it seems we dont need it
		//Y:        y, it seems we dont need it
		Vistited: false,
		//Distance: 2147483647, needed it for dijkstra. dont need it anymore
	}
}

// adds a connection to the RailMap struct from line of input
func addConnection(line string) {

	// Parse the line and update the connections map
	// Example parsing, adjust according to your input format and requirements
	stops := strings.Split(line, "-")
	if len(stops) != 2 {
		fmt.Println("wrong number to connections format:", line)
		return
	}
	//we strip all whitespace from text
	stop := strings.TrimSpace(stops[0])
	stop1 := strings.TrimSpace(stops[1])

	// we put correct map keys to correct variables
	station, ok := mainStations[stop]
	station1, ok1 := mainStations[stop1]
	if !ok || !ok1 {
		fmt.Println("You tried to append to a station that does not exist: ", line)
		return
	}
	// we add stations to struct.stations.connections
	station.Connections = append(station.Connections, station1)
	//biderictional connections for roundabouts.
	station1.Connections = append(station1.Connections, station)

}

// small helper function to remove after #
/* func ignoreHash(input *string) string {
	parts := strings.Split(*input, "#")
	if len(parts) > 0 {
		return parts[0]
	}
	return *input
} */
