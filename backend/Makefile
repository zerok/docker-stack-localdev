docker: $(shell find . -name '*.go')
	docker build --no-cache -t docker-stack-backend:latest .

backend: $(shell find . -name '*.go')
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
