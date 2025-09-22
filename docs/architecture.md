# System Architecture – Go Anomaly Detector

This document describes the **architecture of `go-anomaly-detector`**, a real-time network behavioral anomaly detection system built with **Go** and **Kafka**. It covers key components, data flow, deployment strategies, and scalability considerations.

---

## High-Level Architecture Diagram

Below is a simplified ASCII diagram showing the system flow:

```
   ┌─────────────┐       ┌─────────────┐
   │ Network     │       │ Test Data   │
   │ Traffic     │       │ (Optional)  │
   └─────┬───────┘       └─────┬───────┘
         │                     │
         ▼                     ▼
   ┌─────────────┐
   │ Kafka       │
   │ Broker      │
   └─────┬───────┘
         │
         ▼
   ┌─────────────┐
   │ Go Consumer │
   │ & Detector  │
   └─────┬───────┘
         │
         ▼
   ┌─────────────┐
   │ Logging &   │
   │ Metrics     │
   └─────┬───────┘
         │
         ▼
   ┌─────────────┐
   │ Alerting /  │
   │ Dashboard   │
   └─────────────┘
```

> **Optional:** For a polished visual, create a diagram using **Draw.io** or **Lucidchart**.

---

## Components

### 1. Kafka Producer
- Captures network traffic (live or simulated)
- Publishes messages to Kafka topic `network-data`
- Supports batching and retries for high throughput

### 2. Kafka Broker
- Decouples producer and consumer
- Provides message durability and replication
- Handles high-throughput message streams

### 3. Go Consumer & Anomaly Detector
- Subscribes to Kafka topic
- Processes messages in real-time
- Applies **rule-based** or **ML-based** detection
- Logs anomalies, updates metrics

### 4. Logging & Metrics
- Structured logging with `zap` or `logrus`
- Prometheus metrics endpoint (`/metrics`)
- Integrates with Grafana dashboards for visualization

### 5. Alerting
- Optional notifications via Slack, email, or webhooks
- Sends alerts for anomalies above threshold

---

## Data Flow

1. Network traffic is captured or test data is generated  
2. Producer serializes events and sends them to Kafka  
3. Kafka stores messages across partitions  
4. Consumer reads messages and applies anomaly detection  
5. Detected anomalies are logged and metrics updated  
6. Alerts sent to external systems (optional)  

---

## Deployment Architecture

### Local / Development
- Single Kafka instance
- Producer + Consumer run locally using **Docker Compose**

### Production
- Multiple Kafka brokers with replication
- Consumer scaled with Kubernetes Deployments
- Detection rules and thresholds configured via **ConfigMaps**
- Observability via Prometheus + Grafana

---

## Scaling & Reliability Considerations
- Horizontal scaling of consumers for higher throughput
- Kafka partitions enable parallel processing
- Graceful shutdown and message commit handling
- Docker/Kubernetes deployment ensures resilience

---

## Optional Enhancements
- Integrate ML models for advanced anomaly detection
- Multi-cluster Kafka for geo-redundancy
- Historical data storage for trend analysis and reporting
