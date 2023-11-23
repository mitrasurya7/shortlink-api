# Start from golang base image
FROM golang:1.21-alpine

# Add Maintainer info
LABEL maintainer="Mitra Surya"

# Install git, bash, and build-base
RUN apk update && apk add --no-cache git bash build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working directory inside the container
COPY . .
COPY .env .


# Download all the dependencies
RUN go download

# Build the Go app
RUN go build -o /build

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["/app/build"]
