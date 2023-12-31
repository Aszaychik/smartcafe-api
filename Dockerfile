# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files into the container
COPY . .

# Create a directory for uploads
RUN mkdir /app/uploads

# Build the Go application
RUN make build

# Expose the port that your API will run on
EXPOSE 8080

# Run the application
CMD ["./smartcafe.exe"]