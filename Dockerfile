# Build Stage
FROM golang:alpine:3.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o bin/passwordgen

# Final Stage
FROM alpine:3.20.0 AS final

# Copy the built executable from the build stage
COPY --from=builder /app/bin/passwordgen /usr/local/bin/passwordgen

# Specify the entrypoint for the Docker container
ENTRYPOINT ["/usr/local/bin/passwordgen"]
