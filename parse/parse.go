// Package parse parses the .dat files
// and returns a []string for each
// row.
package parse

import (
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// GetStations takes in the path to a stations.dat
// file and returns a []*Station.
func GetStations(path string) ([]*Station, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}

	stations := []*Station{}

	for _, v := range lines {
		station := NewStation(v)
		if station != nil {
			stations = append(stations, station)
		}
	}

	return stations, nil
}

// GetTrains takes in the path to a trains.dat
// file and returns a []*Train.
func GetTrains(path string) ([]*Train, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}

	trains := []*Train{}

	for _, v := range lines {
		train := NewTrain(v)
		if train != nil {
			trains = append(trains, train)
		}
	}

	return trains, nil
}

// Station is the struct containing info
// about each train station
type Station struct {
	ID   int
	Name string
}

// NewStation takes a string and converts
// it to a new station.
func NewStation(line string) *Station {
	data := strings.Split(line, " ")
	if len(data) != 2 {
		return nil
	}

	// change to int
	id, err := strconv.Atoi(data[0])
	if err != nil {
		return nil
	}

	return &Station{ID: id, Name: data[1]}
}

// Train is the struct containing info
// about each train.
type Train struct {
	DepartureStation int
	ArrivalStation   int
	DepartureTime    string
	ArrivalTime      string
}

// NewTrain takes a string and converts
// it to a new train.
func NewTrain(line string) *Train {
	data := strings.Split(line, " ")
	if len(data) != 4 {
		return nil
	}

	// change to int
	ds, err := strconv.Atoi(data[0])
	if err != nil {
		return nil
	}

	as, err := strconv.Atoi(data[1])
	if err != nil {
		return nil
	}

	return &Train{
		DepartureStation: ds,
		ArrivalStation:   as,
		DepartureTime:    data[2],
		ArrivalTime:      data[3],
	}
}

// GetTimeDelta will return the difference between the
// arrival and departure times.
func (t *Train) GetTimeDelta() (time.Duration, error) {
	depHour, err := strconv.Atoi(t.DepartureTime[0:2])
	depMin, err := strconv.Atoi(t.DepartureTime[3:4])
	arrHour, err := strconv.Atoi(t.ArrivalTime[0:2])
	arrMin, err := strconv.Atoi(t.ArrivalTime[3:4])

	t1 := time.Date(1984, time.November, 3, arrHour, arrMin, 0, 0, time.UTC)
	t2 := time.Date(1984, time.November, 3, depHour, depMin, 0, 0, time.UTC)

	return t1.Sub(t2), err
}

// ReadLines parses .dat files.
func readLines(filename string) ([]string, error) {
	var lines []string
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return lines, err
	}
	buf := bytes.NewBuffer(file)
	for {
		line, err := buf.ReadString('\n')
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}
				return lines, err
			}
		}
		line = strings.TrimSpace(line)
		lines = append(lines, line)
		if err != nil && err != io.EOF {
			return lines, err
		}
	}
	return lines, nil
}
