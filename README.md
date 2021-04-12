# envoy-grpc-status-metrics

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/nokamoto/envoy-grpc-status-metrics)

```bash
$ make datadog DATADOG_API_KEY=<YOUR-DATADOG-API-KEY> # (optional) datadog-cluster-agent
$ skaffold dev --port-forward
```

```bash
$ make stats
LIVE
go run ./cmd/client
2021/04/07 15:42:30 res=, err=<nil>
2021/04/07 15:42:30 res=<nil>, err=rpc error: code = AlreadyExists desc = code=6
2021/04/07 15:42:30 res=<nil>, err=rpc error: code = Internal desc = code=13
2021/04/07 15:42:30 res=<nil>, err=rpc error: code = InvalidArgument desc = code=3
curl -s http://localhost:9901/stats | grep Say
cluster.backend_service.grpc.Server.Say.0: 1
cluster.backend_service.grpc.Server.Say.13: 1
cluster.backend_service.grpc.Server.Say.3: 1
cluster.backend_service.grpc.Server.Say.6: 1
cluster.backend_service.grpc.Server.Say.failure: 3
cluster.backend_service.grpc.Server.Say.request_message_count: 4
cluster.backend_service.grpc.Server.Say.response_message_count: 1
cluster.backend_service.grpc.Server.Say.success: 1
cluster.backend_service.grpc.Server.Say.total: 4
```

```bash
$ make prom
LIVE
go run ./cmd/client
2021/04/07 15:41:51 res=, err=<nil>
2021/04/07 15:41:51 res=<nil>, err=rpc error: code = AlreadyExists desc = code=6
2021/04/07 15:41:51 res=<nil>, err=rpc error: code = Internal desc = code=13
2021/04/07 15:41:51 res=<nil>, err=rpc error: code = InvalidArgument desc = code=3
curl -s http://localhost:9901/stats/prometheus | grep Say
envoy_cluster_grpc_0{envoy_grpc_bridge_method="Say",envoy_grpc_bridge_service="Server",envoy_cluster_name="backend_service"} 1
envoy_cluster_grpc_13{envoy_grpc_bridge_method="Say",envoy_grpc_bridge_service="Server",envoy_cluster_name="backend_service"} 1
envoy_cluster_grpc_3{envoy_grpc_bridge_method="Say",envoy_grpc_bridge_service="Server",envoy_cluster_name="backend_service"} 1
envoy_cluster_grpc_6{envoy_grpc_bridge_method="Say",envoy_grpc_bridge_service="Server",envoy_cluster_name="backend_service"} 1
envoy_cluster_grpc_failure{envoy_grpc_bridge_method="Say",envoy_grpc_bridge_service="Server",envoy_cluster_name="backend_service"} 3
envoy_cluster_grpc_request_message_count{envoy_grpc_bridge_method="Say",envoy_grpc_bridge_service="Server",envoy_cluster_name="backend_service"} 4
envoy_cluster_grpc_response_message_count{envoy_grpc_bridge_method="Say",envoy_grpc_bridge_service="Server",envoy_cluster_name="backend_service"} 1
envoy_cluster_grpc_success{envoy_grpc_bridge_method="Say",envoy_grpc_bridge_service="Server",envoy_cluster_name="backend_service"} 1
envoy_cluster_grpc_total{envoy_grpc_bridge_method="Say",envoy_grpc_bridge_service="Server",envoy_cluster_name="backend_service"} 4
```
