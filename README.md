# Go Anomaly Detector

**go-anomaly-detector** is a high-performance, distributed streaming analytics system **built in Go for real-time detection of behavioral anomalies in network traffic**. It leverages Kafka for scalable ingestion, supports modular producer and consumer services, and is configurable for rule-based or ML-based detection. Ideal for security engineers, network analysts, and developers, it provides Docker and Kubernetes deployment support, monitoring metrics, and logs for production-grade observability.

---

## Features
- Real-time anomaly detection on network traffic
- Kafka-based streaming for high-throughput environments
- Modular Go services: producer and consumer
- Configurable detection rules (rule-based or ML-based)
- Docker and Kubernetes deployment-ready
- Logging, metrics, and monitoring support

---

## System Requirements
- **Operating System:** Linux or macOS
- **Go:** 1.20+
- **Kafka:** 3.x+ cluster running
- **Disk Space:** 20GB+
- **RAM:** 8GB+ recommended
- **Dependencies:** `go mod`, Docker (optional), Kubernetes CLI (optional)

---

## Folder Structure
```
go-anomaly-detector/
├── cmd/                 # Main entry points (producer/consumer)
├── pkg/                 # Shared libraries (Kafka client, detector, utils)
├── configs/             # Configuration files
├── scripts/             # Scripts for running services
├── deployments/         # Docker and Kubernetes manifests
├── docs/                # Architecture, usage, and contribution guides
├── testdata/            # Sample network traffic for testing
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

---

## Quick Start

### Clone the repository
```bash
git clone https://github.com/m-a-h-b-u-b/go-anomaly-detector.git
cd go-anomaly-detector
```

### Install dependencies
```bash
go mod tidy
```

### Run Kafka Producer
```bash
cd cmd/producer
go run main.go
```

### Run Kafka Consumer & Anomaly Detector
```bash
cd cmd/consumer
go run main.go
```

### Run with Docker Compose
```bash
cd deployments/docker
docker-compose up --build
```

---

## Configuration
All configurable parameters are in `configs/config.yaml`:
```yaml
kafka:
  brokers:
    - "localhost:9092"
  topic: "network-data"
  group_id: "anomaly-detector"

detector:
  anomaly_threshold: 0.8
  detection_mode: rule-based
```
- Easily override with **environment variables** for production.

---

## Metrics & Monitoring
- Supports Prometheus metrics endpoint (`/metrics`)
- Logs processed messages, detected anomalies, and errors
- Integration with Grafana dashboards recommended for visual monitoring

---

## Example Use Cases
- Detect unusual network behavior in real-time
- Deploy a lightweight behavioral threat detection system in an enterprise network
- Experiment with streaming analytics and anomaly detection algorithms
- Security monitoring for IoT devices or embedded networks

---

## Contributing
We welcome contributions! Please submit **issues** or **pull requests** for:
- Bug fixes
- New detection algorithms
- Kafka performance enhancements
- Documentation improvements

See `docs/contributing.md` for detailed guidelines.

---

## License
This project is licensed under the **Apache 2.0**.
