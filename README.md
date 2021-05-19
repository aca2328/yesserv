# yesserver
A simple http echo server that accept any http URI and send back container hostname + http request detail as text

except for /metric that return prometheus format http statistics

# local build and run
```
docker build . -t yesserv:latest
docker run -ti --rm -p 8000:8000 yesserv
```

# run from dockerhub
```
docker push aca2328/yesserv:latest
docker run -ti --rm -p 8000:8000 aca2328/yesserv
```
