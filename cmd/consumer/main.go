package main

import (
    "log"
    "github.com/segmentio/kafka-go"
    "golang-behavioral-anomaly-detector/pkg/detector"
)

func main() {
    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "network-data",
        GroupID: "anomaly-detector",
    })
    defer reader.Close()

    for {
        msg, err := reader.ReadMessage(nil)
        if err != nil {
            log.Println("Error reading message:", err)
            continue
        }
        detector.Process(msg.Value)
    }
}
