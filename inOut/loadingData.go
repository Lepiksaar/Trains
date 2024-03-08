package inOut

import (
	"bufio"
	"fmt"
	"os"
	"stations/structs"
	"strconv"
	"strings"
)

var mainStations = make(map[string]*structs.Station)

// function that reads the info into a struct
func FillStruct(filePath, start, end string) map[string]*structs.Station {
	numStations := 0
	place := true
	foundStatline := false

	mapFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error:", err)
	}
	defer mapFile.Close()

	scanner := bufio.NewScanner(mapFile)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "stations:" {
			foundStatline = true
			continue
		}
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
			tempStation := makeStation(line, &numStations)
			// we add the required struct to a map
			mainStations[tempStation.Name] = &tempStation
		} else {
			addConnection(line)
		}
	}
	// a lot of error managment
	_, ok := mainStations[start]
	if !ok {
		fmt.Fprintf(os.Stderr, "Start station not found within stations (%v) \n ", start)
		os.Exit(1)
	}

	_, ok2 := mainStations[end]
	if !ok2 {
		fmt.Fprintf(os.Stderr, "Start station not found within stations (%v) \n ", end)
		os.Exit(1)
	}

	if !foundStatline {
		fmt.Fprintf(os.Stderr, "stations: not found \n")
		os.Exit(1)
	}

	if place {
		fmt.Fprintf(os.Stderr, "connections: not found \n")
		os.Exit(1)
	}

	fmt.Println("\n..... Successfully loaded Stations and connections")
	return mainStations
}

func makeStation(line string, nuStat *int) structs.Station {
	parts := strings.Split(line, ",")

	if len(parts) < 3 {
		fmt.Fprintf(os.Stderr, "Invalid station on %v:", line)
		return structs.Station{}
	}
	name := strings.TrimSpace(parts[0])

	//checking if the station allready exists.
	_, ok := mainStations[name]
	if ok {
		fmt.Fprintf(os.Stderr, "Duplicate station found (%v) \n ", name)
		return structs.Station{}
	}

	//increment the number of stations
	if *nuStat > 10000 {
		fmt.Fprintf(os.Stderr, "reached maximum number of stations(10000)")
		return structs.Station{}
	}
	*nuStat += 1

	//removing whitespace and ignoring evrithing after #
	uglyX := strings.TrimSpace(ignoreHash(&parts[1]))
	uglyY := strings.TrimSpace(ignoreHash(&parts[2]))
	//assigning cordinates
	x, err1 := strconv.Atoi(uglyX)
	y, err2 := strconv.Atoi(uglyY)
	if err1 != nil || err2 != nil || x < 0 || y < 0 {
		fmt.Fprintf(os.Stderr, "Invalid Coordinates: %v ", name)
		return structs.Station{}
	}
	for _, structure := range mainStations {
		if structure.X == x && structure.Y == y {
			fmt.Fprintf(os.Stderr, "Falure to add (%v) to the same coordinates as (%v) \n", name, structure.Name)
			return structs.Station{}
		}

	}
	//we return the struct that we want to add to map
	return structs.Station{
		Name:     name,
		X:        x,
		Y:        y,
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
	if !ok {
		fmt.Fprintf(os.Stderr, "Failure to add connection, station %v does not exist on line: %v\n", stop, line)
		return
	} else if !ok1 {
		fmt.Fprintf(os.Stderr, "Failure to add connection, station %v does not exist on line: %v\n", stop1, line)
		return
	}
	// we add stations to struct.stations.connections
	for _, i := range station.Connections {
		if i == station1 {
			fmt.Fprintf(os.Stderr, "duplicate connection found between %v and: %v\n", station1.Name, station.Name)
			return
		}
	}
	// we add stations to struct.stations.connections
	station.Connections = append(station.Connections, station1)
	//biderictional connections for roundabouts.
	station1.Connections = append(station1.Connections, station)

}

// small helper function to remove after #
func ignoreHash(input *string) string {
	parts := strings.Split(*input, "#")
	if len(parts) > 0 {
		return parts[0]
	}
	return *input
}
