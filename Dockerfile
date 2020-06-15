FROM golang:latest AS build

RUN mkdir -p /usr/src/micro-bastion
WORKDIR /usr/src/micro-bastion
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:latest

WORKDIR /srv/
COPY --from=build /usr/src/micro-bastion/micro-bastion .

CMD ["/srv/micro-bastion"]

