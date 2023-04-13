## Step 1: making build
FROM golang:alpine3.17 as build

WORKDIR /src

COPY . .

RUN go mod download

RUN go build -o build/getcep cmd/getcep/main.go

## Step 2: creating image
FROM golang:alpine3.17

WORKDIR /src

COPY --from=build /src/build/getcep .

CMD ["sleep", "infinity"]
