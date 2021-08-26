
## Simple url shortner service
It generates short-code on bases of keys. The range of keys is provided to each server.Using fastify framework for performance and maintainability.   
-------------------------------------------------------

## Arch.
![Alt text](url-short-block.png?raw=true "Title")
-------------------------------------------------------
1. The default range is { min : 0, max : 9999}. The key service which is not part of this
project can be responsible for provided one when api server's range of keys exhaust.
The maximum key can go upto 2 to the power of 64.
2. The same url created anonymously will get same token
## Start up:

1. `npm install` //Install packages

2. `npm run dev` //for development run

3. `npm start` //for prod run

4. Runs at `http://localhost:3000/`

## Take a look
![Alt text](example.gif?raw=true "Title")
-------------------------------------------------------
1. `POST : /anonymous/create` Create shortlink request as anonymous user
    `body : {url : "https://github.com/bron10/url-shortner"}`

2. `GET : /:token` Request redirection to actual url
This newly generated token is recieved in each request

