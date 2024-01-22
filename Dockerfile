##
# Builder
##
FROM golang:latest AS builder

# Install necessary dependencies
RUN apt-get update && apt-get install -y git

# Copy src
COPY . /go/src/dyndns-agent

# Enter the agent directory
WORKDIR /go/src/dyndns-agent

# Build
RUN go generate
RUN go mod tidy
RUN go build -o dyndns-agent


##
# Main image
##
FROM debian:stable-slim
RUN apt-get update && apt-get upgrade
RUN apt install -y ca-certificates curl 

# Copy build
COPY --from=builder /go/src/dyndns-agent/dyndns-agent /app/dyndns-agent
RUN chmod +x /app/dyndns-agent

WORKDIR /app

CMD ["/app/dyndns-agent"]