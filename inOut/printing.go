package inOut

import "fmt"

func Printout(tracks [][]string, numTrains int) {
	// arrange tracks from longest to shortest
	currentTrain := 1
	numberOftracks := len(tracks)
	fmt.Println(numberOftracks)
	for numTrains >= currentTrain {
		for i := len(tracks) - 1; i >= 0; {
			if len(tracks[i])-len(tracks[0]) <= numTrains-currentTrain-len(tracks[0])+1 {
				fmt.Printf("T%v-%v ", currentTrain, tracks[i][1])
				currentTrain++
				i--
				if numTrains <= currentTrain {
					break
				}
			} else {
				fmt.Printf("T%v-%v", currentTrain, tracks[0][1])
				currentTrain++
				i--
				if numTrains <= currentTrain {
					break
				}
			}
		}
		fmt.Println()
		//turn starts
		//check if longest route beats the shortest route+trains left.

		//send train that way.give info train is in first stop
		//no --- continue
		//... check other n-routes if they beat like first
		//send train that way. give info train is in first stop
		//no--- continue
		// send train to shortest route
		//if first turn end turn. give info train is in first stop
		// no trains to send end loop
	}
	//if first turn end turn

}
