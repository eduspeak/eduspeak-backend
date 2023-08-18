# Use an official Golang runtime as the base image
FROM golang:1.20.4

# Set the working directory
WORKDIR /app

# Copy the Go Modules files
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download
# Copy the source code into the container
COPY . ./

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Expose the port that the application runs on
EXPOSE 3000

# Command to run the executable
CMD ["./app"]
