#build stage
FROM golang:1.10.3-alpine AS build-env
ENV GOPROXY=https://gocenter.io 

ADD . /src
RUN apk update \
    && apk add --virtual build-dependencies \
        git

RUN go get -u github.com/aws/aws-xray-sdk-go/...
RUN cd /src && go build -o myapp

FROM alpine
RUN apk add openssh 
WORKDIR /app
COPY --from=build-env /src/myapp /app/
ENTRYPOINT ["./myapp"]
