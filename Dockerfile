FROM golang:1.16 AS build
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -buildid="

FROM gcr.io/distroless/static

COPY --from=build /build/dump-server /dump-server
CMD [ "/dump-server" ]
