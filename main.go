package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// represents a station in our map
type Station struct {
	Name     string
	Distance int
	Parent   string
	Visited  bool
	X        int
	Y        int
}

// a map containing the stations and their connections
type RailMap struct {
	Stations    []*Station
	Connections map[*Station][]*Station
}

// displays program usage
func displayHelp() {
	fmt.Printf("train scheduler usage:\ngo run . [path to file containing network map] [start station] [end station] [number of trains]")
}

// builds stations array
func buildStations(filePath string) ([]Station, RailMap) {
	var stations []Station
	var connections RailMap
	mapFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return stations, connections // Return an empty slice or handle the error
	}
	defer mapFile.Close()
	scanner := bufio.NewScanner(mapFile)
	for scanner.Scan() {
		place := true
		//build stations until "connections:" is hit
		line := scanner.Text()
		if line == "stations:" {
			continue
		}
		if line == "connections:" {
			//switch to building connections
			place = false
			continue
		}
		line = trimLine(line)
		if line == "" {
			continue
		}
		if place {
			station := makeStation(line)
			stations = append(stations, station)
			connections.Stations = append(connections.Stations, &station)
		} else {
			connections = addConnection(line, stations, connections)
		}
	}
	return stations, connections
}

// makes a Station struct from line of input
func makeStation(line string) Station {
	parts := strings.Split(line, ",")
	name := parts[0]
	x, err1 := strconv.Atoi(parts[1])
	y, err2 := strconv.Atoi(parts[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid Coordinates: " + name)
		//optionally quit with error
	}
	station := Station{
		Name:     parts[0],
		X:        x,
		Y:        y,
		Distance: 1 << 20,
		Parent:   "",
		Visited:  false,
	}
	return station
}

// adds a connection to the RailMap struct from line of input
func addConnection(line string, stations []Station, connections RailMap) RailMap {
	stops := strings.Split(line, "-")
	stop1 := stationLookup(stops[0], stations)
	stop2 := stationLookup(stops[1], stations)
	if stop1 == nil || stop2 == nil {
		fmt.Printf("\nInvalid connection: %v\n", line)
		//optionally quit with error
		return connections
	}
	//possibly check here for redundant or reverse connections
	connections.Connections[stop1] = append(connections.Connections[stop1], stop2)
	connections.Connections[stop2] = append(connections.Connections[stop2], stop1)
	return connections
}

// looks up station by name string, returns pointer or nil if not found
func stationLookup(name string, stations []Station) *Station {
	for _, station := range stations {
		if station.Name == name {
			return &station
		}
	}
	return nil
}

// processes raw lines from network map
func trimLine(line string) string {
	parts := strings.Split(line, "#")
	line = strings.ReplaceAll(parts[0], " ", "")
	return line
}

func main() {
	//assess arguments
	args := os.Args
	if len(args) != 5 {
		displayHelp()
		os.Exit(0)
	}
	networkMap, start, end, numTrains := args[1], args[2], args[3], args[4]

	//build our slice of stations and the map
	stations, connections := buildStations(networkMap)

}
