# drain-my-spot
Service draining the k8s worker node in case of spot instances related event occurrence  

## HOWTOs 

### Build
```bash
➜  drain-my-spot git:(plan) ✗ make build
```

### Build test lint
```bash
➜  drain-my-spot git:(plan) ✗ make 
CGO_ENABLED=0 go mod download
CGO_ENABLED=0 go install github.com/golangci/golangci-lint/cmd/golangci-lint
CGO_ENABLED=0 go test -v ./...
?       github.com/bgzzz/drain-my-spot  [no test files]
CGO_ENABLED=0 golangci-lint run
CGO_ENABLED=0 go build -o ./bin/drain-my-spot .
```

*NOTE: make sure go/bin is in your path to be able to run golangci-lint*

# TODO

1. Input params
2. Querrier (+ emulation)
3. Drainer
4. K8s config 
5. Kind tests
6. k8s tests 
7. Unit and contract testing
