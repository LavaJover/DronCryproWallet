FROM golang:latest

WORKDIR /app

COPY ./api-gateway/api-gateway .

COPY ./auth-service/auth-service .

COPY ./api-gateway/config ./api-gateway-config

COPY ./auth-service/config ./auth-service-config

ENV AUTH_CONFIG_PATH=./auth-service-config/config.yaml

ENV CONFIG_PATH=./api-gateway-config/local.yaml

ENTRYPOINT [ "./auth-service", "./api-gateway" ]