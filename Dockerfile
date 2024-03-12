# Use an official Go runtime as a parent image
FROM golang:1.21

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /bi
COPY . /app

# Download and install any required dependencies
RUN go mod download

# Build the Go application
RUN go build -o personal-page .

# Expose port 8080 to the outside world
EXPOSE 5000

# CONFIG 
ENV config=docker

# Command to run the executable
CMD ["./personal-page"]
