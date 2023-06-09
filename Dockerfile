FROM golang:1.20
WORKDIR /go/src/fiber
COPY . .
RUN make build
EXPOSE 3000
RUN cd /go/src/fiber
CMD ["./out/fiber"]