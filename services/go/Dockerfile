# Use an official Golang runtime as a parent image
FROM golang:1.19-alpine

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the current directory contents into the container at /go/src/app
COPY . /go/src/app

# Build the Go application
RUN go build -o goapp .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./goapp"]