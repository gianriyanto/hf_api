# Hello Fresh Take Home Programming Challenge.
by Gian Riyanto

For the most part, this task was a new and challenging learning experience for me. I had the chance to tinker with a new programming language, web framework, and various other tools, as well as explored new back-end focused concepts and design architectures. Despite the learning curve and some technical issues that arise, I was able to complete part of the requirement outlined below:

- Planned and designed the desired data model.

- Created an ORM and connected postgres database.

- Models.go and Controller.go for Menu, Recipe, User.

- Create REST APIs for Menu, Recipe, User

- CRUD operations for Menu, Recipe, User

- E2E test using Postman

The project still has much room for design/architecture requiremnets as well as several missing requirements that are incomplete at the given timeframe. 

### Description
This is a take home task I attempted for Hello Fresh's Software Engineering Intern role. The context, description, and requirements are detailed [here](https://github.com/hello-abhishek/hf-take-home-programming-challenges/blob/main/SOFTWARE-ENGINEER.md)

Tech Stack Used: Golang (Language), Beego v1.12.3 (Framework and ORM), Postgres (Database), Postman (E2E Test)

## Installation

Install [go](https://golang.org/doc/install)

install [beego](https://beego.me/docs/install/)

```bash
go get github.com/astaxie/beego
go get github.com/beego/bee
```

## To Run

```bash
bee run
```

## To Generate Swagger Documentation
To auto generate documentation of the API, use this command:
```bash
bee run -gendoc=true -downdoc=true
```
Now go to the browser
```
http://localhost:8080/swagger/
```
