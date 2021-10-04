# API Documentation

## Rest API

### Movie [/movies]

#### Fetch Movie[GET]

* Example
```
endpoint : http://localhost:3030?searchword=Batman&pagination=1

```

* Response 200 (application/json)
```
{
    "meta": {
        "success":true,
        "message":"success"
    }
	"data":{
        "Response":true,
        "Search:[
            {
                "Poster": "https://m.media-amazon.com/images/M/MV5BNDdjYmFiYWEtYzBhZS00YTZkLWFlODgtY2I5MDE0NzZmMDljXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg",
                "Title": "Batman Forever",
                "Type": "movie",
                "Year": "1995",
                "imdbID": "tt0112462"
            }....
        ]
    }
}
```

#### Detail Movie[Get]

Get access token use phone number and password

* Example
```
    endpoint : http://localhost:3030?title=batman&movie_id=imdbID
```

* Response 200 (application/json)
```
	{
	    "meta": {
            "success": true,
            "message": "success"
        },
        "data": {
            "Actors": "Michael Keaton, Jack Nicholson, Kim Basinger",
            "Awards": "Won 1 Oscar. 9 wins & 26 nominations total",
            ....
        }
	}
```
## GRPC API

### Port [:3031]

#### Function [FetchMovie]

* Request
```
{
    "searchword": "batman",
    "pagination": "2"
}
```
* Success Response 

```
	{
        "search": [
            {
                "title": "Batman: The Killing Joke",
                "year": "2016",
                "imdbID": "tt4853102",
                "type": "movie",
                "poster": "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"
            }, ...
  ]
```

#### Function [DetailMovie]

* Request
```
{
  "title": "batman",
  "movieId": ""
}
```
* Success Response 

```
	{
        "title": "Batman",
        "year": "1989",
        "rated": "PG-13",
        "released": "23 Jun 1989",
        "runtime": "126 min",
        "genre": "Action, Adventure",
        "director": "Tim Burton",
        ...
    }
```