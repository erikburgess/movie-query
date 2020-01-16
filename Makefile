all: push

build:
	docker build . -t luvmysubi/movie-query:latest

push: build
	docker push luvmysubi/movie-query:latest