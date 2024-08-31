# Use the official Golang image with the desired version
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy go.mod and go.sum files and download dependencies
# This step helps in caching dependencies and optimizing build times
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the rest of the application source code
COPY . .

# Build the Go application
# Output binary is placed in /usr/local/bin/app
RUN go build -v -o /usr/local/bin/app ./...

# Use a smaller base image for the final stage to reduce image size
FROM debian:bullseye-slim

# Set the working directory in the final stage
WORKDIR /usr/src/app

# Copy the built Go binary from the builder stage
COPY --from=builder /usr/local/bin/app /usr/local/bin/app

# Set the entrypoint for the container
CMD ["app"]
