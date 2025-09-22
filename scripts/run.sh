#!/bin/bash
# Start producer and consumer
gnome-terminal -- bash -c "cd cmd/producer && go run main.go"
gnome-terminal -- bash -c "cd cmd/consumer && go run main.go"
