# jarjarbinks

![Jarjar](./dist/logo.jpg)

This service has two different endpoints that are only used to save cache entry and find the saved entry with the relevant key.
The cache structure is built on FIFO, but the policies are not extended.


CURL(s)

# POST
curl --request POST \
--url http://localhost:8080/cache/store/v1 \
--header 'Content-Type: application/json' \
--data '{
"key" : "2",
"value": {
"world": 1
},
"expireAt": 20000
}'

# GET
curl --request GET \
--url http://localhost:8080/cache/store/v1/2