FROM ubuntu:latest
LABEL authors="tecs"

ENTRYPOINT ["top", "-b"]