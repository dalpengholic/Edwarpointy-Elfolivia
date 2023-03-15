# Use an official Golang runtime as a parent image
FROM golang:1.20.2-alpine3.16

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Golang binary
RUN go build -o main .

# Set Env variables
ENV EnvPort=8000
ENV EnvMessage=hello

# Run the webserver when the container starts
CMD ./main -c "$EnvMessage" -p "$EnvPort"

