FROM golang:1.16.3
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64
ENV APP_HOME /var/www/html/klikacc
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
COPY . .
# RUN go mod download \
#     go mod vendor \
#     go mod verify \
RUN go build -o /var/www/html/klikacc/app ./cmd/main.go 
ENTRYPOINT [ "/var/www/html/klikacc/app" ]
# EXPOSE 8080
