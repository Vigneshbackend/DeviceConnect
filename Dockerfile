
FROM alpine:latest
RUN apk update && apk add --no-cache git 
RUN apk --no-cache add ca-certificates
# CGO_ENABLED=0 GOOS=linux go build -o main
# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY  build/* .
ADD .env .     

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]