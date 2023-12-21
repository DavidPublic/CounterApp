# Start from a lightweight version of the Go image
FROM golang:alpine as builder

# Set the working directory outside of $GOPATH to enable Go modules
WORKDIR /app

# Copy the go.mod and go.sum files first to cache dependencies
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Use a minimal alpine image to run the app
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /app/main /main

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["/main"]

