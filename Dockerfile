# Start from golang base image
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Rodericus Ifo Krista"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base && apk add --no-cache util-linux && apk add make

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install github.com/google/wire/cmd/wire@latest
RUN go install -v ./...

# Expose port 8081
EXPOSE 8081

# Run the executable
CMD [ "make", "run-build-docker" ]