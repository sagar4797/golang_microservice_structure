# -----------------------------------------------------------------------------
# Build stage - Intermidiate  container
# -----------------------------------------------------------------------------
FROM golang:1.18-alpine3.15 AS builder

# Setup working directory
WORKDIR /app

# Copy go.mod and go.sum files in containers /app folder
COPY go.mod go.sum ./

# Downloads and verify the packages in go.mod file
RUN go mod download
RUN go mod verify
# Copy source directory in containers /app folder
COPY . .

# Build the binary
RUN go build -o main main.go

# -----------------------------------------------------------------------------
# Build main Docker image
# -----------------------------------------------------------------------------
FROM alpine:3.15

# Setup working directory
WORKDIR /app
RUN apk add --no-cache jq

# Copy the file from build stage to run stage
COPY --from=builder /app/main /app/main

COPY migrate.linux-amd64 ./migrate
COPY db/migration ./migration
COPY start.sh ./start.sh

EXPOSE ${PORT}
CMD [ "./start.sh" ]