# microtrader

docker run --rm -ti -v ~/GitCode/microtrader:/app golang:1.14-buster /bin/bash

## Adding RabbitMQ for the golang producer

```
root@99a2fde70f34:/app/components/market-data# go get github.com/streadway/amqp
```

## Start the project

```
docker-compose build
docker-compose up
docker-compose down
```