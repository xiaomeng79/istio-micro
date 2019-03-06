# Metrics

[来源](https://github.com/hb-go/echo-web/blob/master/middleware/metrics/README.md)

## Grafana
浏览器:localhost:3000
```bash
docker run --name grafana -d -p 3000:3000 grafana/grafana
```

## Prometheus
浏览器:localhost:9090
```bash
docker run -d --name prometheus -p 9090:9090 -v ~/tmp/prometheus.yml:/etc/prometheus/prometheus.yml \
              prom/prometheus
              
# Dashboard JSON
# metrics/prometheus/grafana.json
```

配置文件deployments/configs/prometheus.yml
```bash
global:
  scrape_interval:     15s # By default, scrape targets every 15 seconds.
  evaluation_interval: 15s # Evaluate rules every 15 seconds.

  # Attach these extra labels to all timeseries collected by this Prometheus instance.
  external_labels:
    monitor: 'codelab-monitor'

rule_files:
  - 'prometheus.rules.yml'

scrape_configs:
  - job_name: 'prometheus'

    # Override the global default and scrape targets from this job every 5 seconds.
    scrape_interval: 5s

    static_configs:
      - targets: ['localhost:9090']

  - job_name:       'select_web'

    # Override the global default and scrape targets from this job every 5 seconds.
    scrape_interval: 5s

    static_configs:
      - targets: ['www.localhost.com']
        labels:
          group: 'production'

```

## Push模式
### Graphite
浏览器:localhost:8090
```sh
# 登录账户名:root，密码:root
docker run -d --name graphite --restart=always -p 8090:80 -p 2003-2004:2003-2004 -p 2023-2024:2023-2024 -p 8125:8125/udp -p 8126:8126 hopsoft/graphite-statsd

# Dashboard JSON
# middleware/metrics/grafana_graphite.json
```

### InfluxDB
```bash

```
