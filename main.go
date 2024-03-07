package main

import (
	"flag"
	"fmt"
	"os"
	"stations/inOut"
	"stations/logic"
	"time"
)

func main() {
	//assess arguments
	networkMap := flag.String("networkMap", "test.txt", "path/Name for Network Map. default: test.txt")
	start := flag.String("start", "waterloo", "name of starting station. default: waterloo")
	end := flag.String("end", "st_pancras", "name of ending station. default: st_pancras")
	numTrains := flag.Int("numTrains", 1, "number of trains. default: 1")
	flag.Parse()

	startt := time.Now()

	if *numTrains <= 0 {
		fmt.Fprintf(os.Stderr, "number of trains needs to be positive intiger. You entered: %v'\n", *numTrains)
		os.Exit(1)
	}

	if *start == *end {
		fmt.Fprintf(os.Stderr, "start station and end station are the same. You entered: %v'\n", *start)
		os.Exit(1)
	}

	//build our slice of stations and the map
	fmt.Printf("\nStart Station: %v  End Station: %v  Number of Trains: %v \n", *start, *end, *numTrains)
	stations := inOut.FillStruct(*networkMap, *start, *end)
	findTime := time.Since(startt)
	shortestRoutes := logic.FindAllRoutes(stations, *start, *end)
	findRoutes := time.Since(startt)
	inOut.Printout(shortestRoutes, *numTrains)
	elapsed := time.Since(startt)
	fmt.Printf("time took loading drom file %s \n", findTime)
	fmt.Printf("time took finding path %s \n", findRoutes-findTime)
	fmt.Printf("time took printing %s", elapsed-findRoutes)
}
