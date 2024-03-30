package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const numStations = 10000
const maxConnections = 10

func main() {

	// Open a new file for writing
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Write stations
	_, err = writer.WriteString("stations:\n")
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		return
	}

	for i := 1; i <= numStations; i++ {
		_, err := writer.WriteString(fmt.Sprintf("station_%d,%d,%d\n", i, rand.Intn(10000), rand.Intn(10000)))
		if err != nil {
			fmt.Printf("Error writing to file: %s\n", err)
			return
		}
	}

	_, err = writer.WriteString("\nconnections:\n")
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		return
	}

	// For each station, create 1 to 4 connections to random other stations
	for i := 1; i <= numStations; i++ {
		numConnections := rand.Intn(maxConnections) + 1 // Ensure at least 1 connection
		connections := make(map[int]bool)               // To keep track of unique connections

		for len(connections) < numConnections {
			targetStation := rand.Intn(numStations) + 1 // Pick a random station to connect to

			// Avoid self-connections and duplicate connections
			if targetStation != i && !connections[targetStation] {
				connections[targetStation] = true
				_, err := writer.WriteString(fmt.Sprintf("station_%d-station_%d\n", i, targetStation))
				if err != nil {
					fmt.Printf("Error writing to file: %s\n", err)
					return
				}
			}
		}
	}

	// need to flush the writer. or else some of the lines in the end wont be written.
	err = writer.Flush()
	if err != nil {
		fmt.Printf("Error flushing to file: %s\n", err)
		return
	}

	fmt.Printf("test.txt file has been successfully created with %d stations and up to  %d random connections.\n", numStations, maxConnections)
}
