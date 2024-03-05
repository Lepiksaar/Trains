package inOut

import (
	"fmt"
	"sort"
	"stations/structs"
	"strconv"
	"strings"
)

func Printout(tracks [][]string, numTrains int) {
	// tracks are arranged from shortest to longest
	end := true
	currentTrain := 1
	trackSlice := loadNames(tracks)
	printStr := []string{}
	for end {
		// logic that sets start points to each routes.
		//if it is good idea to send from there it appends train number as struct.train location 0
		for i := len(tracks) - 1; i >= 0; i-- {
			//our formula to see if it is good idea to send it from long way starting from longest
			if len(tracks[i])-len(tracks[0]) < numTrains-currentTrain+1 {
				tempTrack := &trackSlice[i]
				tempTrack.Train = append([]int{currentTrain}, tempTrack.Train...)
				tempTrack.InUse = true
				currentTrain++
				if numTrains < currentTrain {
					break
				}
				//if none of the first ones match we give the train to shortest one.
			} else if i == 0 {
				tempTrack := &trackSlice[0]
				tempTrack.Train = append([]int{currentTrain}, tempTrack.Train...)
				tempTrack.InUse = true
				currentTrain++
				break
			}
		}
		for i := 0; i < len(trackSlice); i++ {
			// Directly access each printTrack by index to modify the original slice
			printTrack := &trackSlice[i]

			if !printTrack.InUse {
				continue
			}
			if printTrack.Train == nil {
				continue
			}

			// Append the train and what route it takes (first stop) into printStr
			printStr = append(printStr, "T"+strconv.Itoa(printTrack.Train[0])+"-"+printTrack.Name[0])

			// Set InUse to false. This change will now be reflected in trackSlice
			printTrack.InUse = false
		}
		moveList(&printStr, &trackSlice)
		if currentTrain > numTrains {
			end = false
		}
		//actualPrint(&trackSlice)
	}
	actualPrinting(printStr)
}

// loading station names from list to struct
func loadNames(matrix [][]string) []structs.Track {
	tracks := []structs.Track{}
	for _, row := range matrix {
		trackName := row[1:]
		tracks = append(tracks, structs.Track{
			Name: trackName,
		})
	}
	return tracks
}

// just a double loop function that moves all stations 1 further
func moveList(list *[]string, name *[]structs.Track) {
	newList := make([]string, len(*list))
	b := *list
	actualPrinting(b)
	for index, v := range *list {
		if v == "" {
			continue // Skip iteration if v is nil
		}
		parts := strings.Split(v, "-")
	out:
		for _, n := range *name {
			for i, statName := range n.Name {
				//fmt.Println("Tere")
				if parts[1] == n.Name[(len(n.Name))-1] {
					parts = nil
					break out
				} else if parts[1] == statName {
					parts[1] = n.Name[i+1]
					break out
				}
			}
		}
		if parts != nil {
			newList[index] = parts[0] + "-" + parts[1]
			//fmt.Println(newList)
			*list = newList
		}
	}
}

// function that actually prints the movement. removing emty spaces
func actualPrinting(b []string) {
	//sorting to smallest first
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	//printing with ignoring the empty slices
	for _, a := range b {
		if a == "" {
			continue
		} else {
			fmt.Print(a, " ")
		}
	}
	fmt.Println()

}
