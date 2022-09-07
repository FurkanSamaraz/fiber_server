# Start from the latest golang base image
FROM golang:latest


# Set the Current Working Directory inside the container
ADD . /app


# Copy go mod and sum files
WORKDIR /app


# Build the Go app
RUN go build -o main .



# Command to run the executable
CMD ["./main"]

