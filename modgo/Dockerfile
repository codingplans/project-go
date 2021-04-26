
# 主要测试 docker 本地部署
FROM alpine

RUN mkdir /src && mkdir /ssd
COPY main /src/

RUN echo 123 >> /ssd/aa.log

# 环境变量
ENV NODE_VERSION 1234
# 只执行一次 必执行
ENTRYPOINT exec /src/main >> /ssd/test.log 2>&1

