# Use an official Golang runtime as a parent image
# It is a good idea to specify a version to not be suprised in the future.
FROM golang:1.23.0

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the local Go app into the container
COPY . .

# Build the Go app
RUN go build -o hello .

# Run the executable
CMD ["./hello"]

