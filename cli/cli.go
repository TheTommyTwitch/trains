package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/thetommytwitch/trains/parse"
)

// Cli is the entrypoint to our
// command line interface.
type Cli struct {
	graph *parse.Graph
}

// NewCli returns a new Cli
func NewCli(graph *parse.Graph) *Cli {
	return &Cli{graph: graph}
}

// Executer is the execute function.
func (c *Cli) Executer() {
	fmt.Printf("Make a selection: ")

	number, err := readInput()
	if err != nil {
		panic(err)
	}

	switch number {
	case 1:
		c.printFullSchedule()
	case 2:
		c.printStationSchedule()
	case 3:
		c.lookupStationID()
	case 4:
		c.lookupStationName()
	case 5:
		c.serviceAvailable()
	case 6:
		c.nonstopServiceAvailable()
	case 7:
		c.findShortestRidingTime()
	case 8:
		c.findShortestTravelTime()
	case 9:
		fmt.Println("GoodBye!")
		os.Exit(0)
	default:
		fmt.Println("Choose 1 though 9...")
	}
}

// 1
func (c *Cli) printFullSchedule() {
	for _, train := range c.graph.GetAllTrains() {
		dep := c.graph.GetStationByID(train.DepartureStation)
		arr := c.graph.GetStationByID(train.ArrivalStation)
		if dep != nil && arr != nil {
			fmt.Printf("Train leaving %s at %s headed to %s\n", dep.Name, train.DepartureTime, arr.Name)
		}
	}
}

// 2
func (c *Cli) printStationSchedule() {
	fmt.Printf("Enter station id: ")
	station, err := readInput()
	if err != nil {
		return
	}

	for _, train := range c.graph.GetTrainsByStation(station) {
		dep := c.graph.GetStationByID(train.DepartureStation)
		arr := c.graph.GetStationByID(train.ArrivalStation)
		if dep != nil && arr != nil {
			fmt.Printf("Train leaving %s at %s headed to %s\n", dep.Name, train.DepartureTime, arr.Name)
		}
	}
}

// 3
func (c *Cli) lookupStationID() {
	fmt.Printf("Enter station id: ")
	id, err := readInput()
	if err != nil {
		return
	}
	station := c.graph.GetStationByID(id)
	if station != nil {
		fmt.Printf("The name of the station is %s.\n", station.Name)
	} else {
		fmt.Println("Station not found.")
	}
}

// 4
func (c *Cli) lookupStationName() {
	fmt.Printf("Enter station name: ")
	name, err := readStringInput()
	if err != nil {
		return
	}
	station := c.graph.GetStationByName(name)
	if station != nil {
		fmt.Printf("The id of the station is %v.\n", station.ID)
	} else {
		fmt.Println("Station not found.")
	}
}

// 5
func (c *Cli) serviceAvailable() {
	fmt.Printf("Enter first station id: ")
	firstID, err := readInput()
	if err != nil {
		return
	}

	fmt.Printf("Enter second station id: ")
	secondID, err := readInput()
	if err != nil {
		return
	}

	stations, err := c.graph.ShortestPath(firstID, secondID)
	if err == nil && len(stations) > 0 {
		fmt.Println("Service is available.")
	} else {
		fmt.Println("Service not available.")
	}
}

// 6
func (c *Cli) nonstopServiceAvailable() {
	fmt.Printf("Enter first station id: ")
	firstID, err := readInput()
	if err != nil {
		return
	}

	fmt.Printf("Enter second station id: ")
	secondID, err := readInput()
	if err != nil {
		return
	}

	stations, err := c.graph.ShortestPath(firstID, secondID)
	if err == nil && len(stations) == 2 {
		if stations[0] == firstID && stations[1] == secondID {
			fmt.Println("Nonstop service is available.")
		}
	} else {
		fmt.Println("Service not available.")
	}
}

// 7
func (c *Cli) findShortestRidingTime() {
	fmt.Printf("Enter first station id: ")
	firstID, err := readInput()
	if err != nil {
		return
	}
	firstStation := c.graph.GetStationByID(firstID)
	if firstStation == nil {
		fmt.Println("Station one not found.")
		return
	}

	fmt.Printf("Enter second station id: ")
	secondID, err := readInput()
	if err != nil {
		return
	}
	secondStation := c.graph.GetStationByID(secondID)
	if secondStation == nil {
		fmt.Println("Station two not found.")
		return
	}

	var totalTime time.Duration

	stations, err := c.graph.ShortestPath(firstID, secondID)
	if err == nil && len(stations) > 0 {
		fmt.Printf("Path from %s to %s: \n", firstStation.Name, secondStation.Name)
		for i, v := range stations {
			sta := c.graph.GetStationByID(v)
			if sta != nil {
				fmt.Printf("%s", sta.Name)
				if i < len(stations)-1 {
					tr := c.graph.GetTrain(v, stations[i+1])
					if tr != nil {
						delta, _ := tr.GetTimeDelta()
						totalTime += delta
						fmt.Printf(" --- %v ---> ", fmtDuration(delta))
					}
				}
			}
		}
		fmt.Printf("   total time: %v", fmtDuration(totalTime))
		fmt.Println()
	}
}

// 8
func (c *Cli) findShortestTravelTime() {
	fmt.Println("8 Not implemented")
}

func readInput() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return -1, err
	}

	text = strings.TrimSpace(text)

	num, err := strconv.Atoi(text)
	if err != nil {
		return -1, err
	}

	return num, nil
}

func readStringInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	text = strings.TrimSpace(text)

	return text, nil
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}

// PrintOpts prints the cli options
func (c *Cli) PrintOpts() {
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
