# base image
FROM golang:1.16

RUN apk update && apk add --no-cache git

WORKDIR /app
# copy everything to workdir
COPY . .
# COPY go.mod go.sum ./
# RUN go mod download
# obtain the package
RUN go mod tidy

# build

# RUN CGO_ENABLED=0 GOOS=linux go build -o /main
RUN go build -o binary


# # compile
# RUN go build -o main .
# start the app
CMD ["/app/binary"]