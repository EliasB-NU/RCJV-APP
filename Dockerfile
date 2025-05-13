# Build backend
FROM golang:1.24.1-alpine AS builder-backend

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY backend/ ./backend/

RUN go build ./backend/main.go

# Build adminsite
FROM node:22.13.0-alpine AS builder-adminsite

WORKDIR /app

COPY adminsite/ ./

RUN npm install

RUN npm run build

# Build webview
FROM node:22.13.0-alpine AS builder-webview

WORKDIR /app

COPY webview/ ./

RUN npm install

RUN npm run build

# Final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates wget

WORKDIR /app

COPY ods/ ./ods

COPY --from=builder-backend /app/main .

COPY --from=builder-adminsite /app/dist ./adminsite/dist
COPY --from=builder-webview /app/dist ./webview/dist

EXPOSE 3006

CMD ["./main", "prod"]