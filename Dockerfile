# This file is a template, and might need editing before it works on your project.
FROM golang:1.14 AS builder

MAINTAINER Hollson Hollson@qq.com

WORKDIR /app

COPY ./tmp/ .

COPY --from=builder /usr/src/app/app .
CMD ["./app"]