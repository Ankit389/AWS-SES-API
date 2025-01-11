# AWS SES Mock API

This is a mock implementation of the AWS SES API for testing purposes.

## Features

- Mock email sending endpoints
- Rate limiting
- Request logging
- Statistics tracking
- Docker support

## Setup

### Local Development

1. Install Go 1.21 or later
2. Clone the repository
3. Run `go mod download`
4. Run `go run cmd/main.go`

### Docker

1. Build the image:
```bash
docker build -t aws-ses-mock .
```

2. Run the container:
```bash
docker run -p 8080:8080 aws-ses-mock
```

## API Endpoints

- POST /v1/email/send - Send email
- POST /v1/email/send-raw - Send raw email
- POST /v1/email/send-templated - Send templated email
- GET /v1/email/quota - Get sending quota
- GET /v1/email/statistics - Get sending statistics

## Example Usage

Send an email:
```bash
curl -X POST http://localhost:8080/v1/email/send \
  -H "Content-Type: application/json" \
  -d '{
    "source": "sender@example.com",
    "destination": ["recipient@example.com"],
    "message": {
      "subject": "Test Email",
      "body": "This is a test email"
    }
  }'
```