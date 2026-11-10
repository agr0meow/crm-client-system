FROM golang:1.23-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /bin/crm ./cmd/crm

FROM alpine:3.20
COPY --from=build /bin/crm /bin/crm
ENTRYPOINT ["/bin/crm"]
