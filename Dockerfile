FROM golang:1.19-alpine

WORKDIR /app

COPY . .

EXPOSE 8080

# Command to run the application
CMD ["go", "run", "main.go"]
