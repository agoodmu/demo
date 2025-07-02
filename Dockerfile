# Deploy-Stage
FROM amazon/aws-cli:latest

COPY go.mod .

# Command to run the application
ENTRYPOINT ["sleep"]
CMD ["infinity"]