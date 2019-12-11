FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY . /Users/darren/go/src/github.com/pwcong/app/main.go
WORKDIR /Users/darren/go/src/github.com/pwcong/app/main.go

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
    then \
    app; \
    else \
    go get github.com/pilu/fresh && \

    fresh; \
    fi

EXPOSE 3000