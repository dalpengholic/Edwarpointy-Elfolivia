# 1. Build Image
# Use an official Golang runtime as a parent image
FROM golang:1.20.2-alpine3.16 AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the necessary contents into the container at /app
COPY go.mod webserver.go README.md /app/

# Build the Golang binary
RUN go build -o main .


# ------------------------------

# 2. Minimum-Viable Product (MVP) Image
FROM busybox
WORKDIR /opt/greet/bin
COPY --from=builder /app/ .

# Set Env variables
ENV EnvPort=8000
ENV EnvMessage=hello

# Run the webserver when the container starts
CMD ./main -c "$EnvMessage" -p "$EnvPort"

