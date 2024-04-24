FROM golang:1.22-alpine as builder
COPY . /dockercan
WORKDIR /dockercan
RUN cd ./cmd/dockercan && go install
CMD ["dockercan"]

FROM alpine
RUN mkdir -p /run/docker/plugins
COPY --from=builder /go/bin/dockercan .
CMD ["/dockercan"]