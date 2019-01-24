# weather-or-not
A small CLI tool for getting live weather forecasts straight to your terminal

## Installation
Clone this repo and set GOPATH
```
git clone https://github.com/tlugger/weather-or-not
cd weather-or-not
export GOPATH=$PWD
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
