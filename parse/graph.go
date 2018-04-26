package parse

// Graph is the structure that holds our
// graph data.
type Graph struct {
	Nodes map[int]*Station
	Edges map[int][]*Train
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
