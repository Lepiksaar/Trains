package main

import (
	"flag"
	"fmt"
	"stations/inOut"
	"stations/logic"
)

func main() {
	//assess arguments
	networkMap := flag.String("networkMap", "test.txt", "path/Name for Network Map. default: test.txt")
	start := flag.String("start", "waterloo", "name of starting station. default: waterloo")
	end := flag.String("end", "st_pancras", "name of ending station. default: st_pancras")
	numTrains := flag.Int("numTrains", 1, "number of trains. default: 1")
	flag.Parse()

	//build our slice of stations and the map
	fmt.Printf("\nStart Station: %v  End Station: %v  Number of Trains: %v", *start, *end, *numTrains)
	stations := inOut.FillStruct(*networkMap)
	shortestRoutes := logic.FindAllRoutes(stations, *start, *end)
	inOut.Printout(shortestRoutes, *numTrains)
}
