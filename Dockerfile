FROM golang:1.20-bullseye
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o server .
EXPOSE 8080
CMD [ "/app/server" ]
