package main

import (
    "log"
    "time"
    "github.com/segmentio/kafka-go"
)

func main() {
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "network-data",
    })
    defer writer.Close()

    for i := 0; i < 100; i++ {
        msg := kafka.Message{
            Key:   []byte("key"),
            Value: []byte("sample network packet data"),
        }
        err := writer.WriteMessages(nil, msg)
        if err != nil {
            log.Println("Failed to write message:", err)
        }
        time.Sleep(1 * time.Second)
    }
}
