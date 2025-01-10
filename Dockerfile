##### Stage 1 #####
FROM golang:1.23-alpine as builder

LABEL maintainer="nguyenmanh180102@gmail.com"
LABEL description="Dockerfile for Go Application with Multi-stage Build"

ENV APP_PATH=/project

RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH

# Copy Go application dependency files
COPY go.mod . 
COPY go.sum .

# Download Go application module dependencies
RUN go mod download

# Copy actual source code for building the application
COPY . .

ENV CGO_ENABLED=0

RUN go build -o app cmd/main.go

##### Stage 2 #####
FROM scratch

ENV DIST_PATH=/dist
WORKDIR $DIST_PATH

# Copy the built application
COPY --from=builder /project/app .
COPY --from=builder /project/docserror.md ./docserror.md

# Copy the configuration file
COPY --from=builder /project/conf/service.env ./conf/service.env


# Set the command to run the application with the configuration file
CMD ["./app", "-config=./conf/service.env"]
