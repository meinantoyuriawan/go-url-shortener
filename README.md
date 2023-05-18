# Golang url shortener
Url shortener using golang and redis

to start using docker, do it by:
1. docker build -t meinanto/redis-test:latest -f Dockerfile .
2. docker compose up

makesure that port 6379 (for redis) and 8080 (for go) is not already used. 

To test the api, can do it by :
1. http://localhost:8080/create-short-url and using :
    {
        "long_url": (the url desired),
        "user_id" : (user id)
    }
    it will return the short-url
2. http://localhost:8080/go-to/{shortUrl} by using shor-url obtained before