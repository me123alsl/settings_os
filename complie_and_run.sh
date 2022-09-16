#!/bin/sh

# Compile the program golang with linux os prograrm arm64
GOOS=linux GOARCH=arm64 go build -o main main.go