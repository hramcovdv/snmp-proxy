FROM golang:1.23-alpine AS build
WORKDIR /src/
COPY . .
RUN go mod download \
&& go mod verify \
&& go build -v -o snmp-proxy main.go

FROM alpine:3.20
COPY --from=build /src/snmp-proxy /bin/snmp-proxy
EXPOSE 8080
ENTRYPOINT ["/bin/snmp-proxy"]