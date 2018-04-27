# trains

points: 40

submit: trains.zip to dropbox

For this assignment, you are to find a solution to a graph problem by utilizing various structures that we have
discussed and used this past semester. You will be able to choose the language of your choice however I will
mostly only be able to assist if you choose to use C++. Below you will find the details of the problem, think about
what you must achieve and plan your approach BEFORE starting to write code.

For this project you will be in charge of helping travelers schedule their trips on trains which are leaving and
arriving at various stations. You are to create a program that will let a user find a path between two stations
among other features. You will be provided with two files: trains.dat which will have the schedule of trains
running between stations and stations.dat which contains the list of stations that are in your train network.
Your program will provide the following functionality:

 - Print complete train schedule for all stations
 - Print complete train schedule for a specific station
 - For a station name, look up station number
 - For a station number, look up the satation name
 - Determine if there is a direct rout from station A to station B
 - Determine if station B can be reached from station A
 - For any two stations determine the shortests amount of time it will take to go from A to B wihtout layovers(format: HH:MM) If no route exists, alert the user
 - For any two stations determine the shortests overall travel time including layovers at stations (format:HH:MM)If no route exists, alert the user
 - For any two stations determine the shortest overall travel time including layovers at stations when requesting to leave at a certain time (format: HH:MM)
 - A passenger is able to say they want to leave at 09:30 and your program will take this into account
when choosing paths
 - Whenever a user is asked for a departure and arrival station they should enter the station numbers. Theexception to this is the look up station id by name function.


Input Files

 - Your program will read in 2 input files. `stations.dat` and `trains.dat`.

 - `stations.dat` contains the mapping of station names to their unique id numbers. The file format will consist of a series of ID and name pairs. An example file would look similiar to
```
1 madison
2 brookings
3 sioux_falls
4 fargo
```

You can assume the following about the `stations.dat` file: 
 - The station id's will fall in the range of 1 to 199
 - The id's may not be sequencially in order
 - Each station name will be at most 25 characters in length and no spaces will be included
 - The file will not specify how many stations are in the file
 - This file contains information regarding the trains that will be traveling between the various stations. The information contained in this file includes the departure and arrival station id's as well as the departure and arrival
times in 24 hour time. 
 - A sample file would look similiar to
```
1 2 0830 1120
1 4 1100 1540
3 2 1200 1600
4 3 1600 1800
2 1 0900 1000
```

You can assume the following about the `trains.dat` file

 - The arrival and departure times will always be 4 digits in length
 - No trains will cross the midnight mark. In other words no train would leave the stations at 2300 and arrive at
0130 the next morning
 - The arrival and departure times will be in 24 hour time
 - This assignment is much more free form then the ones i have given you previously. There is no expected output
file to match or provided header file implementations to meet.    - You should take this as an opportunity to show
what you have learned over the past semester to implement a more ambitious program(even an extremely arbitrary one). Choose your structures wisely and this is actually fairly strait forward. 
 - The STL library or other built
in libraries for stacks/queues/vectors are open for use with the exception of any prebuilt graphing structures or algorithms.


Expectations
 - This assignment is much more free form then the ones i have given you previously. There is no expected output
file to match or provided header file implementations to meet. You should take this as an opportunity to show
what you have learned over the past semester to implement a more ambitious program(even an extremely
arbitrary one). Choose your structures wisely and this is actually fairly strait forward. The STL library or other built
in libraries for stacks/queues/vectors are open for use with the exception of any prebuilt graphing structures or
algorithms.

Extra credit
- for the find shortest riding time and shortests overall travel time functions, print the itinerary of trains the user
would take for that trip

Example Output
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
Enter option: 2
Enter station id: 1
Schedule for madison
Departure to brookings at 0830, arriving at 1120
Departure to fargo at 1100, arriving at 1540
Arrival from brookings at 2200
Enter option: 4
Enter station name: madison
madison's station id is 1
Enter option: 5
Enter departure station id: 1
Enter destination station id: 2
Service is available from madison to brookings
Enter option: 7
Enter departure station id: 2
Enter destination station id: 3
Time on train to go from brookings to sioux_falls is 7 hours and 40 minutes
Itinerary
---------
Leave from brookings at 0900, arrive at madison at 1000
Leave from madison at 1100, arrive at fargo at 1540
Leave from fargo at 1600, arrive at sioux_falls at 1800
Enter option: 9
Goodbye!
```