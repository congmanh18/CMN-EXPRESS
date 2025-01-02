##### Stage 1 #####
FROM golang:1.23-alpine as builder

RUN mkdir -p /project
WORKDIR /project

### Copy Go application dependency files
COPY go.mod .
COPY go.sum .

### Download Go application module dependencies
RUN go mod download

### Copy actual source code for building the application
COPY . .

ENV CGO_ENABLED=0

RUN go build -o app cmd/main.go

##### Stage 2 #####
FROM scratch

WORKDIR /dist

### Copy the built application
COPY --from=builder /project/app .

### Copy the configuration file
COPY --from=builder /project/conf/service.env ./conf/service.env

### Set the command to run the application with the configuration file
CMD ["./app", "-config=./conf/service.env"]
