# parking-lot [![Build Status](https://travis-ci.com/azbshiri/parking-lot.svg?token=1dbZM2CNEGcT8cVBF1Eg&branch=master)](https://travis-ci.com/azbshiri/parking-lot)
This is a solution for Gojek parking lot test


# Installation
To build this locally you should either have `go` or `docker` installed

```shell script
$ cd parking-lot
$ go install
$ parking-lot -h
Usage of parking-lot:
  -f string
        Execute parking lot instructions from given file
  -i    Execute parking lot instructions from interactive shell

```

### Running `parking-lot` binary inside a Docker container:

```shell script
$ cd parking-lot
$ docker build -t parking-lot
$ docker run -ti parking-lot -h
Usage of parking-lot:
  -f string
        Execute parking lot instructions from given file
  -i    Execute parking lot instructions from interactive shell
```



### Read from a file
```shell script
$ go build
$ ./parking-lot -f file_inputs.txt
Created a parking lot with 6 slots
Allocated slot number: 1
Allocated slot number: 2
Allocated slot number: 3
Allocated slot number: 5
Allocated slot number: 5
Allocated slot number: 6
Slot number 4 is free
Slot No.        Registeration No        Colour
1               KA-01-HH-1234           White
2               KA-01-HH-9999           White
3               KA-01-BB-0001           Black
5               KA-01-HH-2701           Blue
6               KA-01-HH-3141           Black
Allocated slot number: 4
Sorry, parking lot is full
KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333
1, 2, 4
6
Not found

```
## Read from interactive shell
```shell script
$ go build
$ ./parking-lot -i
parking-lot> create_parking_lot 6
Created a parking lot with 6 slots
parking-lot> park KA-01-HH-1234 White
Allocated slot number: 1
parking-lot> park KA-01-HH-1234 White
Allocated slot number: 2
parking-lot> park KA-01-HH-1234 White
Allocated slot number: 3
parking-lot> park KA-01-HH-1234 White
Allocated slot number: 4
parking-lot> park KA-01-HH-1234 White
Allocated slot number: 5
parking-lot> park KA-01-HH-1234 White
Allocated slot number: 6
parking-lot> leave 4
Slot number 4 is free
parking-lot> status
Slot No.        Registeration No        Colour
1               KA-01-HH-1234           White
2               KA-01-HH-9999           White
3               KA-01-BB-0001           Black
5               KA-01-HH-2701           Blue
6               KA-01-HH-3141           Black
parking-lot> park KA-01-HH-1234 White
Allocated slot number: 4
parking-lot> park KA-01-HH-1234 White
Sorry, parking lot is full
parking-lot> registeration_numbers_for_cars_with_colour White
KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333
parking-lot> slot_numbers_for_cars_with_colour White
1, 2, 4
parking-lot> slot_number_for_registeration_number KA-01-HH-1234
6
parking-lot> slot_number_for_registeration_number MA-01-HH-1234
Not found
```


# CLI usage
```shell script
Usage of parking-lot:
  -f string
        Execute parking lot instructions from given file
  -i    Execute parking lot instructions from interactive shell
````

