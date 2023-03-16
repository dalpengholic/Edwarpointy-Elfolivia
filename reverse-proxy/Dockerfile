# 1. Build Image
# Use an official Golang runtime as a parent image
FROM golang:1.20.2-alpine3.16 AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the necessary contents into the container at /app
COPY go.mod reverse-proxy.go README.md /app/

# Build the Golang binary
RUN go build -o proxy .


# ------------------------------
# 2. Minimum-Viable Product (MVP) of reverse proxy Image
FROM busybox
WORKDIR /opt/proxy/bin
COPY --from=builder /app/ .

# Run the proxy server when the container starts
CMD ["./proxy"]
