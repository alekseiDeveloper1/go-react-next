FROM golang:1.24.3

WORKDIR /app
ENV AIR_USE_POLLING=true
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin


CMD ["air"]
