# Stage 1: Build the Go application
FROM golang:1.18.3-alpine3.16 AS builder

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o server .

# Stage 2: Create the lightweight image using scratch
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/server /server

# Command to run the application
CMD ["/server"]
