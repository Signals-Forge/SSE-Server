# Base Image
FROM golang:alpine  

# Work Directory
WORKDIR /app

# Copy source code 
COPY . /app

# Build your executable 
RUN go build -o sse-server 

# Expose the port your SSE server uses
EXPOSE 8080

# Define the startup command
CMD ["./sse-server"] 
