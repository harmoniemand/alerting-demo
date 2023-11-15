
## Build golang
FROM golang:1.21.1 AS build_go

WORKDIR /app
RUN cd /app
COPY ./ ./
RUN go mod download
RUN go build -o /alertingdemo ./cmd/alerting-demo/main.go

## DEPLOYMENT
FROM golang:1.21.1

WORKDIR /app
COPY --from=build_go /alertingdemo /app/alertingdemo
EXPOSE 3000
CMD [ "/app/alertingdemo" ]