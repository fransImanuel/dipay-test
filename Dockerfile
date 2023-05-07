FROM golang:1.19

ENV MONGO_URI=host 
    

WORKDIR /app


COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go mod tidy


RUN go build -o dipaytest .

EXPOSE 8080

CMD ./dipaytest -MONGO_URI=$MONGO_URI 
