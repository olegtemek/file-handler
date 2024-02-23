FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ["./go.mod", "/go.sum", "./"]
RUN go mod download

#BUILD
COPY . .
RUN go build -o ./bin/file-handler cmd/file-handler/main.go


FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/file-handler /
COPY --from=builder /usr/local/src/.env /

CMD ["/file-handler"]
EXPOSE 8000