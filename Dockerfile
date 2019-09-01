FROM golang:1.12.7

WORKDIR /go/src/github.com/azbshiri/parking-lot
COPY . .

RUN go install &&\
    rm -rf /go/src/github.com/azbshiri/parking-lot

ENTRYPOINT ["parking-lot"]
CMD ["parking-lot"]