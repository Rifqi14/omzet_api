FROM golang:1.16-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./
COPY domain/files/id_rsa ./
ADD dump.sql /docker-entrypoint-initdb.d
# COPY .env.example ./.env

ENV APP_SERVICE_NAME=test-majoo
ENV APP_HOST=127.0.0.1:8080
ENV APP_LOCALE=en
ENV DB_HOST=test-majoo-db
ENV DB_PORT=3306
ENV DB_USER=root
ENV DB_PASS=majoo
ENV DB_NAME=test-majoo
ENV DB_MAX_CONNECTIONS=200
ENV DB_MAX_LIFETIME_CONNECTION=60
ENV DB_SSL_MODE=disable
ENV SECRET=
ENV SECRET_REFRESH_TOKEN=
ENV TOKEN_EXP_TIME=72
ENV REFRESH_TOKEN_EXP_TIME=1440
ENV JWE_PRIVATE_KEY=./id_rsa
ENV JWE_PRIVATE_KEY_PASSPHRASE=
ENV REDIS_HOST=test-majoo-redis:6379
ENV REDIS_PASSWORD=majoo

RUN go build -o test-majoo main.go

EXPOSE 8080

CMD [ "./test-majoo" ]