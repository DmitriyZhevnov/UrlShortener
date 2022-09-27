# UrlShortener

### Example of request for generate short link:

```
POST
URL: http://localhost:8080
Body:
{
    "url": "https://www.google.ru/example"
}
```
Response should be like that:
```
"https://www.shortener.de/UgGKi8tJoB"
```

### Example or request for get long link:
GET
URL: http://localhost:8080/UgGKi8tJoB
Response should be like that:
```
"https://www.google.ru/example"
```

## Build & Run (Locally)
1. Rename `.env example` file to `.env`
2. Run command for build and up application:
```
make drun
```
3. Run command for initialize Postrges database:
```
make migrate
```

## REST API
Endpoint | Method | Response codes | Description
--- | --- | --- | ---
*/ | POST | 200, 400, 500 | Generate short link
*/:uri | GET | 200, 404, 500 | Generate short link
