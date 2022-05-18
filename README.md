# golang_api

Steps to execute this project:

Checkout the codebase 

Do a docker-compose up

Three models considered for testing are User, Address and Company with respective relations. The api handlers are:

GET    /user-api/user   

POST   /user-api/user


URL:http://localhost:8000/user-api/user


payload:
`{
    "name": "Keith Mc",
    "email": "test",
    "phone": "we did not",
    "address": {
        "street": "Lotus pond road",
        "city": "Amravati"
    },
    "company":{
        "name": "inmar"
    }
}`

GET    /user-api/user/:id


PUT    /user-api/user/:id


DELETE /user-api/user/:id
