# Base Image
FROM golang:1.24-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download all

# Build the Go app
# RUN go build -o main .
RUN go build -o main . && ls -l main
RUN ls -l 

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

