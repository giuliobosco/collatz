FROM golang

WORKDIR /app
COPY go.mod ./
COPY main.go ./

RUN go build -o /collatz

CMD [ "/collatz" ]
