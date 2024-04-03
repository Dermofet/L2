# Stage 1: Build stage
FROM golang:1.21.5 AS builder

WORKDIR /backend/

COPY develop/dev11/cmd /backend/develop/dev11/cmd
COPY develop/dev11/internal /backend/develop/dev11/internal
COPY develop/dev11/dev /backend/develop/dev11/dev

COPY go.mod /backend/go.mod
COPY go.sum /backend/go.sum

RUN go build -o /backend/build /backend/develop/dev11/cmd/L2

# Stage 2: Final stage
FROM ubuntu:22.04

WORKDIR /backend/

COPY --from=builder /backend/build /backend/build
COPY --from=builder /backend/develop/dev11/dev/.env /backend/dev/.env
COPY --from=builder /backend/develop/dev11/internal/app/migrations /backend/internal/app/migrations

CMD [ "/backend/build" ]
