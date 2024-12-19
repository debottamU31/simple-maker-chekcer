
FROM golang:1.23-bullseye AS builder


LABEL maintainer="debottam.upadhyaya@gmail.com"
LABEL description="Sample maker-checker application"
LABEL version="1.0.0"

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN sqlc generate


RUN useradd -u 1001 -m appuser


# FOR Linux Distros
RUN CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -ldflags='-linkmode external -extldflags "-static"' \
    -o server ./internal/main.go

# FOR M! MAC
# RUN CGO_ENABLED=1 \
#     GOOS=linux \
#     GOARCH=arm64 \
#     go build -ldflags='-linkmode external -extldflags "-static"' \
#     -o server ./internal/main.go


FROM alpine:3.17.5


LABEL maintainer="debottam.upadhyaya@gmail.com"
LABEL description="Sample maker-checker application"
LABEL version="1.0.0"

WORKDIR /app


RUN apk add --no-cache sqlite \
    && apk upgrade --no-cache


RUN adduser -D -u 1001 appuser


RUN chown -R appuser:appuser /app
COPY --from=builder --chown=appuser:appuser /app/server .
COPY --chown=appuser:appuser schema.sql .
RUN sqlite3 messages.db < schema.sql && chown appuser:appuser messages.db


USER appuser


EXPOSE 8080


ENTRYPOINT ["./server"]
CMD []