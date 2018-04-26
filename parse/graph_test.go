package parse

import (
	"fmt"
	"testing"
)

func setUp() ([]*Station, []*Train) {
	path1 := "../stations.dat"
	stations, _ := GetStations(path1)

	path2 := "../trains.dat"
	trains, _ := GetTrains(path2)

	return stations, trains
}

func getTestGraph() *Graph {
	st, tr := setUp()
	return NewGraph(st, tr)
}

func TestNewGraph(t *testing.T) {
	st, tr := setUp()
	graph := NewGraph(st, tr)
	if graph == nil {
		fmt.Println("graph is nil")
		t.Fail()
	}
	if len(graph.Nodes) != 4 {
		fmt.Println("graph len should be 4 but was: ", len(graph.Nodes))
		t.Fail()
	}
}

func TestEdges(t *testing.T) {
	graph := getTestGraph()

	edges := graph.Edges[1]
	if len(edges) != 2 {
		fmt.Println("edges len should be 2 but was: ", len(edges))
		for _, edge := range edges {
			fmt.Printf("%+v\n", edge)
		}
		t.Fail()
	}

	edges = graph.Edges[2]
	if len(edges) != 1 {
		fmt.Println("edges len should be 2 but was: ", len(edges))
		t.Fail()
	}
}

func TestGetStationByID(t *testing.T) {
	graph := getTestGraph()

	st := graph.GetStationByID(1)
	if st == nil {
		fmt.Println("station is nil")
		t.Fail()
	}
	if st.Name != "madison" {
		fmt.Printf("station name is %s, but should be madison\n", st.Name)
		t.Fail()
	}

	// Check for missing id
	st = graph.GetStationByID(245)
	if st != nil {
		fmt.Println("station should be nil but was not nil")
		fmt.Println(st)
		t.Fail()
	}
}

func TestGetStationByName(t *testing.T) {
	graph := getTestGraph()

	st := graph.GetStationByName("madison")
	if st == nil {
		fmt.Println("station is nil")
		t.Fail()
	}
	if st.Name != "madison" {
		fmt.Printf("station name is %s, but should be madison\n", st.Name)
		t.Fail()
	}

	// Check for missing name
	st = graph.GetStationByName("dfdafaf")
	if st != nil {
		fmt.Println("station should be nil but was not nil")
		fmt.Println(st)
		t.Fail()
	}
}

func TestGetTrainsByID(t *testing.T) {
	graph := getTestGraph()

	trains := graph.GetTrainsByID(1)
	if trains == nil {
		fmt.Println("trains is nil")
		t.Fail()
	}

	if len(trains) != 2 {
		fmt.Println("trains len should be 2 but was: ", len(trains))
	}
}
