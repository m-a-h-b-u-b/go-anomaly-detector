package detector

import "log"

func Process(data []byte) {
    log.Printf("Processing packet of length %d bytes\n", len(data))
}
