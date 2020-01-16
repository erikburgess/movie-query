FROM golang:alpine AS build

RUN mkdir /src

ADD main.go /src
RUN cd /src && \
    go build -o movie-query&& \
    ls -r /src/


FROM alpine

WORKDIR /bin

COPY --from=build /src/movie-query /bin

ENTRYPOINT ["./movie-query"]
CMD ["-h"]