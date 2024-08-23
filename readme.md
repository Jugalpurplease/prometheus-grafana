Here's an updated `README.md` that includes the `prometheus.yml` configuration and how to use it:

---

# Prometheus Metrics with Gin

This Go application sets up a basic web server using the Gin framework and exposes metrics for Prometheus. It includes Docker Compose configuration to orchestrate the application along with Grafana and Prometheus. A Makefile is also provided to simplify common tasks.

## Features

- **Gin Web Server**: Serves HTTP requests and exposes metrics.
- **Prometheus**: Collects and stores metrics.
- **Grafana**: Provides a web interface for visualizing metrics.

## Getting Started

### Prerequisites

- Docker
- Docker Compose
- Make (optional, for Makefile commands)

### Installation

1. **Clone the Repository**

   ```bash
   https://github.com/Jugalpurplease/prometheus-grafana.git
   cd prometheus-grafana
   ```

2. **Configure Prometheus**

   Ensure you have the following configuration in `prometheus.yml`:

   ```yaml
   global:
     scrape_interval: 15s

   scrape_configs:
     - job_name: "rest-server"
       static_configs:
         - targets: ["rest-server:9000"]
   ```

   This configuration tells Prometheus to scrape metrics from the `rest-server` service every 15 seconds.

3. **Build and Run with Docker Compose**

   Use Docker Compose to build and start the services:

   ```bash
   docker-compose up --build
   ```

   This will start the following services:
   - **Grafana**: Accessible on port `3000`.
   - **Prometheus**: Accessible on port `9090`.
   - **Go Application (rest-server)**: Accessible on port `9000`.

4. **Access the Application**

   - Open your browser and go to `http://localhost:9000` to see the Go application.
   - Access the metrics at `http://localhost:9000/metrics`.
   - Open Grafana at `http://localhost:3000` and configure your data source and dashboards.
   - Prometheus is available at `http://localhost:9090` for querying metrics.

### Docker Compose Configuration

Here’s the `docker-compose.yml` file used to configure and run the services:

```yaml
version: '3.1'

services:
  grafana:
    image: grafana/grafana:latest
    container_name: grafana
    ports:
      - "3000:3000"
    volumes:
      - grafana-storage:/var/lib/grafana

  rest-server:
    build:
      context: ./
      dockerfile: Dockerfile
    container_name: golang
    restart: always
    ports:
      - '9000:9000'

  prometheus:
    image: prom/prometheus:v2.24.0
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
      - prometheus_data:/prometheus
    ports:
      - 9090:9090
    restart: always

volumes:
  prometheus_data:
  grafana-storage:
```

### Makefile Commands

The Makefile provides shortcuts for common commands:

- **Re Build  the Docker Image and Start Services**

  ```bash
  make build
  ```

  This command builds the Docker images and starts the services.

- **Start Services**

  ```bash
  make up
  ```

  This command starts the services using existing images.

### Prometheus Configuration

The `prometheus.yml` file is configured as follows:

```yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: "rest-server"
    static_configs:
      - targets: ["rest-server:9000"]
```

This configuration tells Prometheus to scrape metrics from the `rest-server` service every 15 seconds.

### Grafana Configuration

1. **Data Source Configuration**

   In Grafana, configure Prometheus as a data source by creating a `datasources.yml` file with the following content:

   ```yaml
   apiVersion: 1
   datasources:
     - name: Prometheus
       type: prometheus
       url: http://prometheus:9090
       isDefault: true
   ```

2. **Dashboard Provisioning**

   Create a dashboard provisioning configuration file (`dashboards.yml`):

   ```yaml
   apiVersion: 1
   providers:
     - name: 'Default'
       orgId: 1
       folder: ''
       type: file
       disableDeletion: false
       updateIntervalSeconds: 10
       allowUiUpdates: true
       options:
         path: /var/lib/grafana/dashboards
         foldersFromFilesStructure: true
   ```

   Place your dashboard JSON files in the `grafana/dashboards` directory.

### Code Overview

- **Metrics Setup**: `httpRequestsTotal` and `httpRequestDuration` track HTTP request counts and durations.
- **Middleware**: Measures request duration and increments counters.
- **Endpoints**:
  - `/`: Returns a greeting message.
  - `/metrics`: Exposes metrics in Prometheus format.

