FROM golang:1.17-alpine

# Working directory
WORKDIR /app

# copy go.mod and go.sum files
COPY go.mod go.sum ./

# download dependencies
RUN go mod download

# copy the source code
COPY . .

# build the binary
RUN go build -o main .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./main"]