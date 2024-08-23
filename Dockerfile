# Start with the official Go base image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 9000

# Command to run the application
CMD ["./main"]
