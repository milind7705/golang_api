FROM golang:bullseye as DEV

WORKDIR /opt
RUN apt update \
    && apt install -y --no-install-recommends \
        curl \
        # psql
        postgresql-client

COPY ./ ./
RUN go get \
    # ensure that the entrypoint and scripts are executable
    && chmod +x entrypoint.sh

EXPOSE 8000
ENTRYPOINT ["./entrypoint.sh"]
CMD ["go", "run", "main.go"]
