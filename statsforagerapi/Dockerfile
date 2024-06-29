FROM ubuntu:latest

WORKDIR /app

COPY bin/statsforagerapi .

EXPOSE 8000

RUN chmod +x statsforagerapi

ENTRYPOINT [ "./statsforagerapi" ]


