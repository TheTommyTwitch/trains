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
}
