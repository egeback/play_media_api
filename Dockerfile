# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

RUN go get -u github.com/swaggo/swag/cmd/swag

# Add Maintainer Info
LABEL maintainer="Marky Egebäck <marky@egeback.se>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
# RUN go build -o main internal/main.go
RUN ./cmd/build.sh

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./playapi"]
