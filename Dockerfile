FROM golang:1.16-alpine AS build
RUN apk --no-cache add tzdata

RUN mkdir -p /data/blog
COPY ./config /data/blog
WORKDIR /data/blog
RUN GOPROXY=https://goproxy.cn,direct CGO_ENABLED=0 GOOS=linux go build -o blogserver .


###
FROM scratch as final
COPY --from=build /data/blog /data/blog
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai

WORKDIR /data/blog

EXPOSE 8000
CMD ["./blogserver"]