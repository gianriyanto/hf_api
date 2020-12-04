# Hello Fresh Take Home Programming Challenge.
by Gian Riyanto

### Description
This is a take home task I attempted for Hello Fresh's Software Engineering Intern role. The context, description, and requirements are detailed [here](https://github.com/hello-abhishek/hf-take-home-programming-challenges/blob/main/SOFTWARE-ENGINEER.md)

Tech Stack Used: Golang (Language), Beego v1.12.3 Web Framework and ORM, Postgres (Database), Postman (E2E Test), Postico (SQL Browser)

### Remarks
For the most part, this task was a new and challenging learning experience for me. I had the opportunity to tinker with a new programming language, web framework, and various other tools, as well as explored new back-end focused concepts and design architectures. I was able to complete the following:

- Planned and designed the desired data model.

- Created a ORM with beego and connected postgres database.

- Routers, Models.go and Controller.go for Menu, Recipe, User, Ingredient.

- Create REST APIs for Menu, Recipe, User, Ingredient using an MVC architecture.

- CRUD operations for Menu, Recipe, User, Ingredient.

- E2E test with Postman.

Listed are achievement at the given timeframe. Several requirements remain incomplete, solution design has a lot of room for improvement, and more comprehensive testing should be implemented.

To Be Done:

- Finish connecting data models (particularly the 1:M and M:M) according to the designed entity relationship diagram (see [DESIGN.md](https://github.com/gianriyanto/hf_api/blob/master/DESIGN.md))

- Create Unit tests and test runner.

Nonetheless, the current progres at this task was a challenging yet valuable learning experience. Check it out. 


## Setup
Install [go](https://golang.org/doc/install)

install [beego](https://beego.me/docs/install/)

```bash
go get github.com/astaxie/beego
go get github.com/beego/bee
```

## Run

```bash
bee run
```

## Generate Swagger Documentation
To auto generate documentation of the API, use this command:
```bash
bee run -gendoc=true -downdoc=true
```
Now go to the browser
```
http://localhost:8080/swagger/
```
