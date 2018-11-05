FROM golang:1.11-alpine3.7
LABEL authors="John Nolette <john@neetgroup.net>"

ENV LC_ALL=C.UTF-8 \
    LANG=C.UTF-8 \
    APP_LOC=/go/src/github.com/neetjn/go-cars

# install dep package manager
RUN apk update; apk add git
RUN go get -u github.com/golang/dep/cmd/dep

ADD . $(APP_LOC)
WORKDIR $(APP_LOC)
# install dependencies
RUN dep ensure

# run application
CMD go run cars.go
