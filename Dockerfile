# Stage 1: Build stage
FROM golang:1.23.2-bullseye AS build

# Create a non-root user for running the application
RUN useradd -u 1001 nonroot

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o movies-recommendation-poc/server/ ./app/main.go

WORKDIR /app/movies-recommendation-poc

RUN mkdir /config
RUN mkdir /docs

WORKDIR /app

COPY ./config/app_config.json /movies-recommendation-poc/config
COPY ./docs/ /movies-recommendation-poc/docs

# Change the ownership of the files to the non-root user
RUN chown -R nonroot ../movies-recommendation-poc

# Stage 2: Final stage
FROM scratch

WORKDIR /app

EXPOSE 8080

# Copy /etc/passwd for non-root user information
COPY --from=build /etc/passwd /etc/passwd

# Copy the binary from the build stage
COPY --from=build /app/movies-recommendation-poc/ ./

# Use the non-root user created in the build stage
USER nonroot

# Copy the configuration files
COPY --from=build /app/config/ ./config/

# Set the entrypoint command
ENTRYPOINT ["./server/main"]