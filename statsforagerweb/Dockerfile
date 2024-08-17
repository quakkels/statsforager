FROM ubuntu:latest

WORKDIR /app

COPY bin/statsforagerweb .

EXPOSE 8000

RUN chmod +x statsforagerweb

ENTRYPOINT [ "./statsforagerweb" ]


