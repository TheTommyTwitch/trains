package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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
	fmt.Println("Make a selection: ")

	number, err := readInput()
	if err != nil {
		panic(err)
	}

	switch number {
	case 1:
		fmt.Println("1 Not implemented")
	case 2:
		fmt.Println("2 Not implemented")
	case 3:
		fmt.Println("3 Not implemented")
	case 4:
		fmt.Println("4 Not implemented")
	case 5:
		fmt.Println("5 Not implemented")
	case 6:
		fmt.Println("6 Not implemented")
	case 7:
		fmt.Println("7 Not implemented")
	case 8:
		fmt.Println("8 Not implemented")
	case 9:
		fmt.Println("GoodBye...")
		os.Exit(0)
	default:
		fmt.Println("Choose 1 though 9...")
	}
}

// 1
func (c *Cli) printFullSchedule() {

}

// 2
func (c *Cli) printStationSchedule() {

}

// 3
func (c *Cli) lookupStationID() {

}

// 4
func (c *Cli) lookupStationName() {

}

// 5
func (c *Cli) serviceAvailable() {

}

// 6
func (c *Cli) nonstopServiceAvailable() {

}

// 7
func (c *Cli) findShortestRidingTime() {

}

// 8
func (c *Cli) findShortestTravelTime() {

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
