# Advanced Golang Mini Course by Alterra

## About

<pre>
Event organizer : <a target="_blank" href="https://www.alterra.id/">Alterra</a>
Mentor          : <a target="_blank" href="https://www.linkedin.com/in/azka-fadhli-ramadhan-9b070313a/">Azka Fadhli Ramadhan</a>
Participant     : <a target="_blank" href="https://www.linkedin.com/in/irvan-ahmad-prasetya-6306a8115/">Irvan Ahmad P.</a>
</pre>

## Day 1

**Objectives**

- Create Postman collection (JSON file)
- Create Postman environment
- Implement HTTP Method (GET, POST, UPDATE, DELETE)

**Tasks**

Do request to the following API target by using Postman environment, save the result using “Save Response” (Save as example), then export collection.

## Day 2

**Objective**

- Organize code using MVC pattern
- Create static CRUD API using Echo
- Create dynamic CRUD API that connect to database using Echo & Gorm

**Tasks**

Organize your code using MVC (create config, controllers, lib, models, routes)
Create static CRUD API that meets requirements.

## Day 3

**Objective**

- Log Implementation
- Implementing JWT Auth for Protecting API

**Tasks**

- Implement **log middleware** to Day 2 task
- Implement **JWT Auth middleware** based on reqirements

## Day 4

**Objective**

Implement Integration Testing

**Tasks**

- Using project previous day, implement integration testing on controller
- Minimumtest case per function is 2
  - Valid test case, e.g. http response 200, 201
  - Invalid test case, e.g. http response 400, 401, 500
- No minimum coverage percentage but higher is better

## Day 5

Lerning about DDD (Domain Driven Development) & Fundamental Clean Architecture

## Day 6

**Objective**

- Implement hexagonal architecture using Echo framework
- Use GORM and middleware, such as logging and authorization, in hexagonal architecture project

**Tasks**

- Refactor previous MVC project to hexagonal architecture, keep GORM, logging, and authorization middleware also
- Optional: at least one endpoint do create, read, update, or delete operation to NoSQL database

## Day 7

**Tasks**

Dockerize previous project.

## Day 8

Deploy previous project on EC2 (AWS) or Heroku.

## Day 9

Learning about System Design.

## Day 10

**Objective**

- Deployment with CI/CD

**Tasks**

- Create CI/CD (use github actions) to deploy the project.
