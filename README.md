# Trains

### Running options:

If you have go installed:
```
$ go get github.com/thetommytwitch/trains
$ trains -stations=./stations.dat -trains=./trains.dat
```

Without go installed you can use the pre built binary in the repo. Pick your operating system.
```
$ make run_linux
```
```
$ make run_mac
```

### Usage Example

```
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
Make a selection: 1
Train leaving madison at 0830 headed to brookings
Train leaving madison at 1100 headed to fargo
Train leaving brookings at 0900 headed to madison
Train leaving sioux_falls at 1200 headed to brookings
Train leaving fargo at 1600 headed to sioux_falls

Make a selection: 2
Enter station id: 1
Train leaving madison at 0830 headed to brookings
Train leaving madison at 1100 headed to fargo

Make a selection: 3
Enter station id: 1
The name of the station is madison.

Make a selection: 4
Enter station name: madison
The id of the station is 1.

Make a selection: 5
Enter first station id: 1
Enter second station id: 2
Service is available.

Make a selection: 6
Enter first station id: 1
Enter second station id: 3
Service not available.

Make a selection: 7
Enter first station id: 1
Enter second station id: 3
Path from madison to sioux_falls: 
madison -> fargo -> sioux_falls

Make a selection: 8
8 Not implemented

Make a selection: 9
GoodBye!

```
