FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 

RUN go get -d -v ./...
RUN go build -o main .
RUN go install github.com/BlueMaverick/EPAMassignment
#RUN goa-cellar
#RUN chmod +x /app

EXPOSE '5000'

CMD ["goa-cellar"]
#ENTRYPOINT  ["/app/mainvm.go"]