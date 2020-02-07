FROM golang:1.12.5 AS builder
ADD . dummyBack/
WORKDIR dummyBack/
RUN go mod tidy
RUN go build -o build/dummyBack
RUN mkdir /dist && mv build/dummyBack /dist/dummyBack

FROM ubuntu:latest
COPY --from=builder /dist/dummyBack .
RUN chmod +x dummyBack
ENTRYPOINT ["/dummyBack"]