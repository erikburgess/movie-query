# movie-query

## Requirements
Program must accept a command line parameter of a name of a movie
* Query the OBDb API for a result over the network
* Output the Rotten Tomato rating in some useful way to the user at the command line
* Be “dockerized” – i.e. the work to make it into a docker container must be present

## usage

### cli
```bash
export OMDB_API_KEY=myKey
./movie-query -name "Office Space"
```

### docker
```bash
export OMDB_API_KEY=myKey
docker run -eOMDB_API_KEY=$OMDB_API_KEY docker.io/luvmysubi/movie-query -- -name "Office Space"
```
