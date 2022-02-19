// Golang program to illustrate the usage of

// Including the main package
package main

// Importing fmt
import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Calling main
func main() {

	// Declaring some variables
	var timeLeave string
	var secondReq, hours, minute, secondResp int
	var timeHours, timeMinute, timeSecond, timeHoursFinal, timeMinuteFinal, timeSecondFinal, totalTimeHours, totalTimeMinute, totalTimeSecond int

	// Calling Scanf() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Printf("Masukkan Jam Berangkat : ")
	fmt.Scanf("%s", &timeLeave)
	fmt.Printf("Masukkan Lama tempuh (Dalam Detik) : ")
	fmt.Scanf("%d", &secondReq)

	// save process result into variabel
	hours = secondReq / 3600
	modMinuteSec := secondReq % 3600
	minute = modMinuteSec / 60
	secondResp = modMinuteSec % 60
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

	diffMinute := timeMinute
	diffSecond := timeSecond
	// final time
	totalTimeHours = timeHours + hours
	totalTimeMinute = timeMinute + minute
	totalTimeSecond = timeSecond + secondResp
	if totalTimeSecond > 60 {
		diffSecond = totalTimeSecond - 60
		timeMinuteFinal = totalTimeMinute + int(math.Floor(float64(totalTimeSecond)/60))
		timeSecondFinal = diffSecond
	}

	if totalTimeMinute > 60 {
		diffMinute = timeMinuteFinal - 60
		timeHoursFinal = totalTimeHours + int(math.Floor(float64(totalTimeMinute)/60))
		timeMinuteFinal = diffMinute
	}

	timeHoursFinalStr := strconv.Itoa(timeHoursFinal)
	timeMinuteFinalStr := strconv.Itoa(timeMinuteFinal)
	timeSecondFinalStr := strconv.Itoa(timeSecondFinal)

	// Print Time arrive
	if timeHoursFinal < 10 {
		timeHoursFinalStr = "0" + strconv.Itoa(timeHoursFinal)
	}
	if timeMinuteFinal < 10 {
		timeMinuteFinalStr = "0" + strconv.Itoa(timeMinuteFinal)
	}
	if timeSecondFinal < 10 {
		timeSecondFinalStr = "0" + strconv.Itoa(timeSecondFinal)
	}
	fmt.Println(fmt.Sprintf("Akan Sampai Pada %s:%s:%s ", timeHoursFinalStr, timeMinuteFinalStr, timeSecondFinalStr))

}
