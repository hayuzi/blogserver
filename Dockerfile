FROM golang:1.16-alpine AS build
RUN apk --no-cache add tzdata

RUN mkdir -p /data/blog
COPY . /data/blog
WORKDIR /data/blog
RUN GOPROXY=https://goproxy.cn,direct CGO_ENABLED=0 GOOS=linux go build -o blogserver .


###
FROM scratch as final
COPY --from=build /data/blog/blogserver /data/blog/blogserver
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY ./config/config.docker.yaml /data/blog/config/config.yaml
ENV TZ=Asia/Shanghai

WORKDIR /data/blog

EXPOSE 8081
CMD ["./blogserver", "server"]