FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 

RUN go get -d -v ./...
RUN go build -o main .
RUN go install github.com/BlueMaverick/EPAMassignment
RUN chmod 755 /app

EXPOSE '5000'

CMD ["EPAMassignment"]