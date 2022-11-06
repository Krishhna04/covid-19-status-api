ARG GOLANG_VERSION=1.16.6
FROM golang:${GOLANG_VERSION} as builder

ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Kolkata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

#Configure timezone
RUN apt-get update && apt-get install -y --no-install-recommends tzdata && \
    ln -fs /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone && \
    dpkg-reconfigure --frontend noninteractive tzdata 

WORKDIR /app
COPY . .
RUN go get -d &&\
    go install github.com/swaggo/swag/cmd/swag@latest && \
    swag init && \
    go build -o main .


COPY --from=builder /app/.env ./
COPY --from=builder /app/main ./
COPY --from=builder /app/docs/ ./docs/



CMD ["./main"]

