FROM golang:1.15

WORKDIR /code
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["lookitup.baby"]