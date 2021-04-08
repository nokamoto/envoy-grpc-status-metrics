
all:
	go install google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc
		
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	internal/protobuf/server.proto

	go fmt ./...
	go test ./...
	go mod tidy

stats:
	@curl http://localhost:9901/ready || (echo error: run [skaffold dev --port-forward] && false)
	go run ./cmd/client
	curl -s http://localhost:9901/stats | grep Say

prom:
	@curl http://localhost:9901/ready || (echo error: run [skaffold dev --port-forward] && false)
	go run ./cmd/client
	curl -s http://localhost:9901/stats/prometheus | grep Say

datadog:
	@kubectl config current-context | grep docker-desktop
	helm install datadog -f deployments/helm-datadog-values.yaml  --set datadog.apiKey=${DATADOG_API_KEY} datadog/datadog
