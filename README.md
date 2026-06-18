# Kubernetes Autoscaler

Custom Kubernetes HPA with GPU utilization metrics.

## Features
- GPU utilization-based scaling (not just CPU/memory)
- Custom Prometheus metrics support
- Predictive scaling with time-series forecasting
- Scale-to-zero for batch workloads

## Metrics
```yaml
metrics:
  - type: Pods
    pods:
      metric:
        name: gpu_utilization
      target:
        type: AverageValue
        averageValue: "70"
```

## License
MIT
