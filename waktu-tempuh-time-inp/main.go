// Golang program to illustrate the usage of

// Including the main package
package main

// Importing fmt
import (
	"fmt"
	"strconv"
	"strings"
)

// Calling main
func main() {

	// Declaring some variables
	var timeLeave, timeArrive string
	var timeHours, timeMinute, timeSecond, timeHoursArrive, timeMinuteArrive, timeSecondArrive int
	var timeHoursFinal, timeMinuteFinal, timeSecondFinal, totalTimeHours, totalTimeMinute, totalTimeSecond int

	// Calling Scanf() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Printf("Masukkan Jam Berangkat : ")
	fmt.Scanf("%s", &timeLeave)
	fmt.Printf("Masukkan Jam Datang : ")
	fmt.Scanf("%s", &timeArrive)

	// save process result into variabel
	strTime := []string{}
	//convert to time
	// leave at 08:52:45
	if len(strings.Split(timeLeave, ":")) > 0 {
		for i := range strings.Split(timeLeave, ":") {
			convInt, _ := strconv.Atoi(strings.Split(timeLeave, ":")[i])
			if convInt < 10 {
				strings.Split(timeLeave, ":")[i] = strings.TrimLeft(strings.Split(timeLeave, ":")[i], "0")
			}
			strTime = append(strTime, strings.Split(timeLeave, ":")[i])
		}
		timeHours, _ = strconv.Atoi(strTime[0])
		timeMinute, _ = strconv.Atoi(strTime[1])
		timeSecond, _ = strconv.Atoi(strTime[2])
	}

	strTimeArrive := []string{}
	//convert to time
	// arrive at 12:15:10
	if len(strings.Split(timeArrive, ":")) > 0 {
		for i := range strings.Split(timeArrive, ":") {
			convIntArrive, _ := strconv.Atoi(strings.Split(timeArrive, ":")[i])
			if convIntArrive < 10 {
				strings.Split(timeArrive, ":")[i] = strings.TrimLeft(strings.Split(timeArrive, ":")[i], "0")
			}
			strTimeArrive = append(strTimeArrive, strings.Split(timeArrive, ":")[i])
		}
		timeHoursArrive, _ = strconv.Atoi(strTimeArrive[0])
		timeMinuteArrive, _ = strconv.Atoi(strTimeArrive[1])
		timeSecondArrive, _ = strconv.Atoi(strTimeArrive[2])
	}

	// final time
	totalTimeHours = timeHoursArrive - timeHours
	totalTimeMinute = timeMinuteArrive - timeMinute
	totalTimeSecond = timeSecondArrive - timeSecond
	timeSecondFinal = totalTimeSecond
	timeMinuteFinal = totalTimeMinute
	timeHoursFinal = totalTimeHours
	if totalTimeSecond < 0 {
		timeSecondFinal = (60 + timeSecondArrive) - timeSecond
	}

	if totalTimeMinute < 0 {
		timeMinuteFinal = (60 + timeMinuteArrive) - timeMinute
	}

	if totalTimeHours < 0 {
		timeHoursFinal = (60 + timeHoursArrive) - timeHours
	}

	// Print Time arrive
	fmt.Println(fmt.Sprintf("Lama Waktu Yang Dihabiskan Adalah %d Jam %d Menit %d Detik ", timeHoursFinal, timeMinuteFinal, timeSecondFinal))

}
