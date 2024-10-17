FROM golang:1.23-alpine AS build

WORKDIR /src/

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o snmp-proxy main.go

FROM alpine:3.20

COPY --from=build /src/snmp-proxy /bin/snmp-proxy

EXPOSE 8080

USER nobody

ENTRYPOINT ["/bin/snmp-proxy"]