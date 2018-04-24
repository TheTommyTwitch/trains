package parse

import (
	"fmt"
	"testing"
)

func TestGetStations(t *testing.T) {
	path := "../stations.dat"
	stations, err := GetStations(path)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if len(stations) != 4 {
		fmt.Println("Length of stations was: ", len(stations))
		t.Fail()
	}
}

func TestGetTrains(t *testing.T) {
	path := "../trains.dat"
	trains, err := GetTrains(path)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if len(trains) != 5 {
		fmt.Println("Length of stations was: ", len(trains))
		t.Fail()
	}
}

func TestReadLines(t *testing.T) {
	filePath := "../stations.dat"
	testLines := []string{
		"1 madison",
		"2 brookings",
		"3 sioux_falls",
		"4 fargo",
	}

	lines, err := readLines(filePath)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	for k := range lines {
		if testLines[k] != lines[k] {
			fmt.Printf("testLine %v does not match\n", k)
			fmt.Printf("testLine %v is: %s\n", k, testLines[k])
			fmt.Printf("lines %v is: %s\n", k, lines[k])
			t.Fail()
		}
	}
}

func TestNewStation(t *testing.T) {
	line := "1 madison"

	station := NewStation(line)
	if station == nil {
		fmt.Println("station was nil!")
		t.Fail()
	}

	if station.ID != 1 {
		fmt.Println("station id was: ", station.ID)
		t.Fail()
	}

	if station.Name != "madison" {
		fmt.Println("station name was: ", station.Name)
		t.Fail()
	}
}

func TestNewTrain(t *testing.T) {
	line := "1 2 0830 1120"

	train := NewTrain(line)
	if train == nil {
		fmt.Println("train was nil!")
		t.Fail()
	}

	if train.DepartureStation != 1 {
		fmt.Println("train dep station was: ", train.DepartureStation)
		t.Fail()
	}

	if train.ArrivalStation != 2 {
		fmt.Println("train arr station was: ", train.ArrivalStation)
		t.Fail()
	}

	if train.DepartureTime != "0830" {
		fmt.Println("train dep time was: ", train.DepartureTime)
		t.Fail()
	}

	if train.ArrivalTime != "1120" {
		fmt.Println("train arr time was: ", train.ArrivalTime)
		t.Fail()
	}

}
