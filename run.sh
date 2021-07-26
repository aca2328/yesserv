docker build . -t yesserv:latest
docker run -ti --rm -p 8000:8000 yesserv
## docker run --rm -p 8000:8000 -v $PWD:/go golang go run main.go &
