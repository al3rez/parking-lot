# parking-lot [![Build Status](https://travis-ci.com/azbshiri/parking-lot.svg?token=1dbZM2CNEGcT8cVBF1Eg&branch=master)](https://travis-ci.com/azbshiri/parking-lot)
This is a solution for Gojek parking lot test

![](https://www.ardanlabs.com/images/gopher-kart/gopher-blue-updated.gif)


# Installation
To build this locally you should either have `go` or `docker` installed:
https://docs.docker.com/install/
https://golang.org/doc/install

```shell script
$ cd parking-lot
$ go install
$ parking-lot -h
Usage of parking-lot:
  -f string
        Execute parking lot instructions from given file
  -i    Execute parking lot instructions from interactive shell

```

### Building `parking-lot` binary inside a Docker container:

```shell script
$ cd parking-lot
$ docker build -t parking-lot
$ docker run -ti parking-lot -h
Usage of parking-lot:
  -f string
        Execute parking lot instructions from given file
  -i    Execute parking lot instructions from interactive shell
```



# CLI Usage
```shell script
Usage of parking-lot:
  -f string
        Execute parking lot instructions from given file
  -i    Execute parking lot instructions from interactive shell
  -stdin
        Execute parking lot instructions from standard streams
````

### Read instructions from a file (locally)
```shell script
$ parking-lot -f file_inputs.txt
Created a parking lot with 6 slots
Allocated slot number: 1
Allocated slot number: 2
Allocated slot number: 3
Allocated slot number: 5
Allocated slot number: 5
Allocated slot number: 6
...
```

### Read instructions from a file (it's actually standard stream) (Docker)
```shell script
$ cat file_inputs.txt | docker run -i parking-lot -i -stdin
created a parking lot with 6 slots
allocated slot number: 1
allocated slot number: 2
allocated slot number: 3
allocated slot number: 5
allocated slot number: 5
allocated slot number: 6
...
```

```shell script
$ cat file_inputs.txt | parking-lot -i -stdin
created a parking lot with 6 slots
allocated slot number: 1
allocated slot number: 2
allocated slot number: 3
allocated slot number: 5
allocated slot number: 5
allocated slot number: 6
...
```

### Read instructions from interactive shell (locally/Docker)
```shell script
$ docker run -ti parking-lot -i
parking-lot> create_parking_lot 6
Created a parking lot with 6 slots
parking-lot> park KA-01-HH-1234 White
...
```

```shell script
$ parking-lot -i
parking-lot> create_parking_lot 6
Created a parking lot with 6 slots
parking-lot> park KA-01-HH-1234 White
Allocated slot number: 1
parking-lot> park KA-01-HH-1234 White
Allocated slot number: 2
...
```

## Running tests

```shell script
$ cd parking-lot
$ go test -v ./...
````
