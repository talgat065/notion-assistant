FROM golang:1.19-alpine AS build

WORKDIR /app
COPY ./go.mod ./

RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o bin ./cmd/web

FROM scratch
WORKDIR /app
COPY --from=build /app/bin/web /app/web
EXPOSE 443

ENTRYPOINT ["/app/web"]
# Build: docker build -t notionassistant .
# Run: docker run notionassistant -p 443:443
