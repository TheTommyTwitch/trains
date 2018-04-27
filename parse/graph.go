package parse

import (
	"container/heap"
	"errors"
)

// Graph is the structure that holds our
// graph data.
type Graph struct {
	Nodes   map[int]*Station
	Edges   map[int][]*Train
	visited map[int]bool
}

// NewGraph makes a graph from trains and stations.
func NewGraph(stations []*Station, trains []*Train) *Graph {
	g := new(Graph)
	g.Nodes = make(map[int]*Station)
	g.Edges = make(map[int][]*Train)

	for _, st := range stations {
		g.Nodes[st.ID] = st
		for _, tr := range trains {
			// Edges for that station
			if tr.DepartureStation == st.ID {
				g.Edges[st.ID] = append(g.Edges[st.ID], tr)
			}
		}
	}

	return g
}

// GetAllStations returns all stations
func (g *Graph) GetAllStations() (s []*Station) {
	for _, station := range g.Nodes {
		s = append(s, station)
	}
	return
}

// GetAllTrains returns all trains
func (g *Graph) GetAllTrains() (t []*Train) {
	for _, v := range g.Edges {
		for _, train := range v {
			t = append(t, train)
		}
	}
	return
}

// GetTrainsByStation gets trains leaving
// a given station id.
func (g *Graph) GetTrainsByStation(id int) (t []*Train) {
	for _, train := range g.Edges[id] {
		t = append(t, train)
	}
	return
}

// GetStationByID returns a station for the given id.
// returns nil if not found.
func (g *Graph) GetStationByID(id int) *Station {
	if st, ok := g.Nodes[id]; ok {
		return st
	}

	return nil
}

// GetStationByName return a station for a given name.
// returns nil if not found.
func (g *Graph) GetStationByName(name string) *Station {
	for _, st := range g.Nodes {
		if st.Name == name {
			return st
		}
	}
	return nil
}

// GetTrainsByID returns all the trains that
// leave that given station.
func (g *Graph) GetTrainsByID(id int) []*Train {
	if tr, ok := g.Edges[id]; ok {
		return tr
	}

	return nil
}

// ShortestPath algo....
func (g *Graph) ShortestPath(src, dest int) ([]int, error) {
	visited := make(map[int]bool)
	dists := make(map[int]float64)
	prev := make(map[int]int)

	dists[src] = 0
	queue := &queue{&queueItem{value: src, weight: 0, index: 0}}
	heap.Init(queue)

	for queue.Len() > 0 {
		// Done
		if visited[dest] {
			break
		}

		item := heap.Pop(queue).(*queueItem)
		n := item.value
		for _, edge := range g.Edges[item.value] {
			dest := edge.ArrivalStation
			delta, _ := edge.GetTimeDelta()
			dist := dists[n] + float64(delta)
			if tentativeDist, ok := dists[dest]; !ok || dist < tentativeDist {
				dists[dest] = dist
				prev[dest] = n
				heap.Push(queue, &queueItem{value: dest, weight: dist})
			}
		}
		visited[n] = true
	}

	if !visited[dest] {
		return nil, errors.New("no shortest path exists")
	}

	path := []int{dest}
	for next := prev[dest]; next != 0; next = prev[next] {
		path = append(path, next)
	}

	// Reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path, nil
}

//////////////////////////////
// QUEUE ////////////////////
/////////////////////////////
type queueItem struct {
	value  int
	weight float64
	index  int
}

type queue []*queueItem

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i].weight < q[j].weight
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *queue) Push(x interface{}) {
	n := len(*q)
	item := x.(*queueItem)
	item.index = n
	*q = append(*q, item)
}

func (q *queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}

func (q *queue) update(item *queueItem, weight float64) {
	item.weight = weight
	heap.Fix(q, item.index)
}
