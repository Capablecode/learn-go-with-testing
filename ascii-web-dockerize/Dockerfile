# Use official Golang image
FROM golang:1.24

# Create app directory inside container
WORKDIR /app

# Copy go mod files first
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy all project files
COPY . .

# Build the Go application
RUN go build -o ascii-web-dockerize .

# Expose the port your app runs on
EXPOSE 8080

# Run the executable
CMD ["./ascii-web-dockerize"]