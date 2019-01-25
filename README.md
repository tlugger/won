# weather-or-not
A small CLI tool for getting live weather forecasts straight to your terminal

## Installation
Clone this repo
```
git clone https://github.com/tlugger/weather-or-not won
cd won
```

Copy weather-or-not env file
```
cp .wonenv.example .wonenv
```

Copy your Dark Sky API Key, Latitude, and Longitude into .wonenv

## Run
Get current weather summary and Temperature

_Compile then run_
```
go build won.go
./won
```

_Only run_
```
go run won.go
```

## Install
_Add GOBIN to your path_
```
export PATH=$PATH:$GOPATH/bin
```

_Install won to your GOPATH_
```
go install
```

_Run won!_
```
$ won
Current condition: Flurries
Current temperature: 29.97°F
```