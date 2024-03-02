package main

import (
	"flag"
	"fmt"
	inout "stations/inOut"
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
	stations := inout.FillStruct(*networkMap)
	shortestRoutes := logic.Createmap(stations, *start, *end)
	inout.Printout(shortestRoutes, *numTrains)

	fmt.Printf("\nstart: %v  end: %v  numTrains: %v\n", *start, *end, *numTrains)
	fmt.Println("stations: ")
	fmt.Println(stations)
	fmt.Println("this is shortest route", shortestRoutes)
}
