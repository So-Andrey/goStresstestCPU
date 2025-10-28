# goStresstestCPU

A simple Go program that loads all CPU cores to 100% using concurrent goroutines.

## Description
The program detects the number of CPU cores and runs one goroutine per core.  
Each goroutine performs continuous math operations to generate CPU load and prints a message every 5 seconds.

## Usage
```bash
git clone https://github.com/So-Andrey/goStresstestCPU.git
cd goStresstestCPU
go run main.go
