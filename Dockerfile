# build stage
# golang:1.17.7-alpine3.5
FROM golang@sha256:d030a987c28ca403007a69af28ba419fca00fc15f08e7801fc8edee77c00b8ee AS builder

# Create appuser
ENV USER=appuser
ENV UID=10001

RUN adduser \    
  --disabled-password \    
  --gecos "" \    
  --home "/appuser" \    
  --shell "/sbin/nologin" \       
  --uid "${UID}" \    
  "${USER}"

WORKDIR /go/src/app

RUN apk --update add ca-certificates

# sownload deps
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

# copy code to build
COPY . .

# build static version, without cgo and debug symbols for the relevant arch
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o api-seed -v ./cmd

# final stage 0MB image
FROM scratch

# copy user data
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder --chown=appuser:appuser /appuser /appuser
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR         /usr/src/app

# copy app's binary 
COPY --from=builder /go/src/app/api-seed ./api-seed

# Use an unprivileged user.
USER appuser:appuser

EXPOSE 8002

CMD ["./api-seed"]
