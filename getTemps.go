package main

import (
	"encoding/csv"
	"fmt"
	"github.com/yryz/ds18b20"
	"os"
	"strconv"
	"time"
)

func main() {
	loc, e := time.LoadLocation("America/Los_Angeles")
	CheckError(e)

	myTime := time.Now().In(loc)
	timestamp := myTime.Format(time.RFC3339)
	const numSensors int = 2 //hardcoding this - a bit of a hack

	sensors, e := ds18b20.Sensors()
	CheckError(e)

	//fmt.Printf("sensor IDs: %v\n", sensors)

	var data [][]string

	for _, sensor := range sensors {
		temp, err := ds18b20.Temperature(sensor)
		if err == nil {
			//		fmt.Printf("sensor: %s temperature: %.2f°C\n", sensor, temp)
			e := []string{timestamp, sensor, strconv.FormatFloat(temp, 'f', 3, 32)}

			//		fmt.Printf("e: %v\n", e)
			data = append(data, e)
		}
	}

	//fmt.Printf("data: %v\n", data)

	file, e := os.Create("./data/" + timestamp + ".csv")
	CheckError(e)

	writer := csv.NewWriter(file)

	writer.WriteAll(data)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}
