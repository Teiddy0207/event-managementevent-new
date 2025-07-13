FROM golang:1.23-alpine

WORKDIR /app

# Copy module files và tải dependencies
COPY ./be-event/go.mod ./be-event/go.sum ./
RUN go mod download

# Copy toàn bộ source code Go
COPY ./be-event ./

# Build ứng dụng Go
RUN go mod tidy 
RUN go build -o main .

# Chạy ứng dụng
EXPOSE 8080
CMD ["./main"]
