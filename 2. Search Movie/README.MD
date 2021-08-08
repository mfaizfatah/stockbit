# API SEARCH MOVIE

In this service have 2 transport: REST JSON HTTP and GRPC

Before you start, set environment variable for 'apikey'
```
export OMDB_APIKEY="your key"
```

Run the server:
```
go run main.go
```

### REST JSON HTTP

Request:
```
curl --request GET \
  --url 'http://localhost:80/?searchword=Batman&pagination=2'
```

Response:
```{
  "Search": [
    {
      "Title": "string",
      "Year": "string",
      "imdbID": "string",
      "Type": "string",
      "Poster": "string"
    },
    .....
  ],
  "totalResults": "string",
  "Response": "string"
}
```

### GRPC

Request
```
gRPC localhost:81/SearchService/SearchMovie
```

Response:
```
Same as REST JSON HTTP
```