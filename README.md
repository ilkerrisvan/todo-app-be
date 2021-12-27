![Go](https://img.shields.io/badge/Go-1.17.3-f21170?style=flat-square&logo=docker&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-3.3.2-f21170?style=flat-square&logo=docker&logoColor=white)

# Todo Rest API

This API only creates a task and get all tasks.

## Table of Contents:

- [Getting Started](#getting-started)
    - [Requirements](#requirements)
    - [Building with Docker](#with-docker)
- [API Endpoints and Documentations](#api-endpoints-and-documentations)
    - [GET `/todo`](#get-allkeyvaluepairs)
    - [POST `/todo`](#post-keyvaluepair)
- [Contact Information](#contact-information)
- [License](#license)

<br/>

## Getting Started

<hr/>

### Requirements:

<hr/>

- Go v1.17.3 or higher -> [Go Installation Page](https://go.dev/dl/)
- Docker v3.3.2 or higher (optional) -> [Docker Get Started Page](https://www.docker.com/get-started)
  <br/>

  Before starting the application, fork/download/clone this repo.

<hr/>

### Building with Docker

- To run the application on [localhost:8000](http://localhost:8000):

```
docker-compose up --build
```


## API Endpoints and Documentations

<hr/>

### GET `/todo`

- Description: Returns all tasks from the db.


#### Request:

```
GET Request to '/todo' endpoint. 
```


#### Reponse:

```
[
    {
        "date": "2021-12-27T00:00:00Z",
        "text": "foo"
    }
]
```




### POST `/todo`

- Description: Creates new task.

#### Request Body:

```
   {
            "text": "bar"
   }
```

#### Reponse:

```
{
    "date": "2021-12-27T16:16:47.209023+03:00",
    "text": "bar"
}
```
<hr/>





## Contact Information

<hr/>

#### Author: İlker Rişvan

#### Github: ilkerrisvan

#### Email: ilkerrisvan@outlook.com

#### Date: December, 2021

## License

<hr/>

[MIT](https://choosealicense.com/licenses/mit/)
