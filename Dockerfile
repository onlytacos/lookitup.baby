FROM golang:1.19

WORKDIR /code
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["lookitup.baby"]
