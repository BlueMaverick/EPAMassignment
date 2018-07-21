FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 

RUN go get -d -v ./...
RUN go build -o main .
RUN go install ["C:\Users\ab107p\Desktop\Olympic\Code\goa-cellar"]
#RUN goa-cellar
RUN chmod 777 -R /app

EXPOSE '5000'

CMD ["goa-cellar"]
#ENTRYPOINT  ["/app/mainvm.go"]