# UrlShortener

`/generate` 

Example of body request:
```
{
    "url": "https://www.google.ru/234ertert5454"
}
```


### Build & Run (Locally)
1. Rename `.env example` file to `.env`
2. Run command for build and up application:
```
make drun
```
3. Run command for initialize Postrges database:
```
make migrate
```

### REST API
Endpoint | Method | Response codes | Description
--- | --- | --- | ---
*/generate | POST | 200, 400, 500 | Generate short link
