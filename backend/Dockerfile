FROM golang:1.9
RUN mkdir -p /go/src/github.com/zerok/docker-stack-localdev/backend
COPY vendor/ /go/src/github.com/zerok/docker-stack-localdev/backend/vendor/
COPY main.go /go/src/github.com/zerok/docker-stack-localdev/backend/
WORKDIR /go/src/github.com/zerok/docker-stack-localdev/backend
RUN CGO_ENABLED=0 GOOS=linux go build -a
RUN ls -l /go/src/github.com/zerok/docker-stack-localdev/backend/

FROM alpine:3.1
RUN mkdir /app
WORKDIR /app/
COPY --from=0 /go/src/github.com/zerok/docker-stack-localdev/backend/backend /app/
RUN ls -l
ENTRYPOINT ["/app/backend"]
EXPOSE 80
