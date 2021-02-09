FROM golang:1.14 as build

# Set the Current Working Directory inside the container
WORKDIR /tmp/app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go get -d -v ./...

COPY . .

RUN find . -name "*.go"

RUN CGO_ENABLED=0 go build -o ./out/prace ./*.go

FROM alpine

COPY --from=build /tmp/app/out/prace /app/prace

CMD ["/app/prace"]