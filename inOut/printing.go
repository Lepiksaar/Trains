package inOut

import (
	"fmt"
	"stations/structs"
	"strconv"
	"strings"
)

func Printout(tracks [][]string, numTrains int) {
	// arrange tracks from longest to shortest
	end := true
	currentTrain := 1
	trackSlice := loadNames(tracks)
	printStr := []string{}
	for end {
		for i := len(tracks) - 1; i >= 0; i-- {
			if len(tracks[i])-len(tracks[0]) <= numTrains-currentTrain-len(tracks[0])+2 {
				tempTrack := &trackSlice[i]
				tempTrack.Train = append([]int{currentTrain}, tempTrack.Train...)
				currentTrain++
				//i--
				if numTrains <= currentTrain {
					break
				}
			} else {
				tempTrack := &trackSlice[0]
				tempTrack.Train = append([]int{currentTrain}, tempTrack.Train...)
				currentTrain++
				//i--
				break
			}
		}
		fmt.Println(trackSlice)
		for _, printTrack := range trackSlice {
			printStr = append(printStr, "T"+strconv.Itoa(printTrack.Train[0])+"-"+printTrack.Name[0])
		}
		moveList(&printStr, &trackSlice)
		fmt.Println(printStr)
		if currentTrain > numTrains {
			end = false
		}
		fmt.Println()
		//actualPrint(&trackSlice)
	}
}
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
	fmt.Println(list)
	for index, v := range *list {
		if v == "" {
			continue // Skip iteration if v is nil
		}
		parts := strings.Split(v, "-")
	out:
		for _, n := range *name {
			for i, statName := range n.Name {
				fmt.Println(parts[0])
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
			fmt.Println(parts[0], parts[1])
			newList[index] = parts[0] + "-" + parts[1]
			//fmt.Println(newList)
			*list = newList
		}
	}
}
