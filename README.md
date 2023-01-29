
# wm-back-end

This project is the bridge between MongoAtlas and my front web Wallmart01Front


## Deployment

To deploy this project run

```bash
  docker build --tag wm-back-end .
```

To run this project:

```bash
  docker run -d -p 8080:8080 --name wm-back-end wm-back-end
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`MONGO_CONNECTION_SCHEME`
`MONGO_HOST`
`MONGO_USER`
`MONGO_PASS`
`MONGO_DATABASE`


## API Reference

#### Get Products

```http
  GET /search/searchString
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `searchString` | `string` | returns a list of products in base of parameter. |

#### Get all products

```http
  GET /search/
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `none`    | `string` | Returns the first 10 items from the database. |




## Usage/Examples

Postman
```
http://localhost:8080/search/asdfdsa
```
When its connected to the database it should return an array of 17 objects


## Links for the project
 Live url:

  - [https://back-data01.fly.dev/search/](https://back-data01.fly.dev/)

Github Project:
 - [https://github.com/HansBukerG/wm-back-end](https://github.com/HansBukerG/wm-back-end)
 Docker Project:

 - [https://hub.docker.com/repository/docker/hansbukerg/wm-back-end/](https://hub.docker.com/repository/docker/hansbukerg/wm-back-end/)

 SonarCloud Test:

  - [https://sonarcloud.io/project/overview?id=HansBukerG_wm-back-end](https://sonarcloud.io/project/overview?id=HansBukerG_wm-back-end)



## Authors

- [@HansBukerG](https://www.github.com/HansBukerG)

