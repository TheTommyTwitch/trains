package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thetommytwitch/trains/parse"
)

func main() {
	stationsFile := flag.String("stations", "", "a string")
	trainsFile := flag.String("trains", "", "a string")

	flag.Parse()

	if *stationsFile == "" || *trainsFile == "" {
		fmt.Println("usage: go run main.go -stations=./stations.dat -trains=./trains.dat")
		os.Exit(1)
	}

	stations, err := parse.GetStations(*stationsFile)
	if err != nil {
		panic(err)
	}

	trains, err := parse.GetTrains(*trainsFile)
	if err != nil {
		panic(err)
	}

	for _, station := range stations {
		fmt.Printf("%+v\n", station)
	}
	for _, train := range trains {
		fmt.Printf("%+v\n", train)
	}

	printOpts()
}

func printOpts() {
	fmt.Printf(
		`
========================================================================
READING RAILWAYS SCHEDULER
========================================================================
Options - (Enter the number of your selected option)
(1) - Print full schedule
(2) - Print station schedule
(3) - Look up stationd id
(4) - Look up station name
(5) - Servie available
(6) - Nonstop service available
(7) - Find route (Shortest riding time)
(8) - Find route (Shortest overall travel time)
(9) - Exit
`)
}
