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
		// because tracks are set from shortest to longest, we start from the longest and check, if we can
		for i := len(tracks) - 1; i >= 0; {
			//our formula to see if it is good idea to send it from long way
			if len(tracks[i])-len(tracks[0]) < numTrains-currentTrain+1 {
				tempTrack := &trackSlice[i]
				tempTrack.Train = append([]int{currentTrain}, tempTrack.Train...)
				currentTrain++
				i--
				if numTrains < currentTrain {
					break
				}
			} else {
				tempTrack := &trackSlice[0]
				tempTrack.Train = append([]int{currentTrain}, tempTrack.Train...)
				currentTrain++
				i--
				break
			}
		}
		for last, printTrack := range trackSlice {
			// because we started checking from longest, the algo also prefers longer routes.
			//so if time is same, we have to continue from our first route
			if numTrains < currentTrain && last == 0 {
				continue
			}
			// when there is only one train. Or some tracks are not used, then it skipps them also
			if printTrack.Train == nil {
				continue
			}
			// we append the train and what route it takes(first stop) into a []string
			printStr = append(printStr, "T"+strconv.Itoa(printTrack.Train[0])+"-"+printTrack.Name[0])
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
