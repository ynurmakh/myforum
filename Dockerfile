FROM ubuntu:latest

WORKDIR /APP/.

EXPOSE 9999

RUN apt update && \
    apt upgrade -y && \
    apt install golang git sqlite3 -y

CMD [ "bash", "-c", "go run ./cmd/" ]