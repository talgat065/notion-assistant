FROM golang:1.17-alpine AS build

WORKDIR /app
COPY ./go.mod ./

RUN go mod download
COPY / ./
RUN echo $(ls -lha)
RUN go build -o bin ./cmd/cli

FROM scratch
WORKDIR /
COPY --from=build /app/bin/cli /cli
EXPOSE 8080

ENTRYPOINT ["/cli"]
