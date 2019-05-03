### USER API


### Architecture Design
```
https://www.lucidchart.com/documents/edit/9ef14b7a-53ea-44da-bb61-fb7b1038d453/0
```
### TechStack
```
Go
Docker
```
### API URLs
```


```


### Instructions
```

```

### Go Read ME
```
Http Request
method route version


```

### Links
```
Rest API Important Link https://blog.restcase.com/5-basic-rest-api-design-guidelines/
CQRS https://www.bouvet.no/bouvet-deler/utbrudd/a-simple-todo-application-a-comparison-on-traditional-vs-cqrs-es-architecture
```

	router.HandleFunc("/users", handlers.DeleteUserHandler).Methods("DELETE")
	//Below routers are not into the business they are for testing
	router.HandleFunc("/test", handlers.TestHandler).Methods("POST")
	router.HandleFunc("/ping", handlers.PingHandler).Methods("GET")
  
### User Service API Schema

####     GET /ping  
    Ping the userapi service endpoint; only for testing purposes  
    
    Response:
    - 200 Success
    - 404 Not Found
</br>

####     GET /test  
    Only for testing purposes  
    
    Response:
    - 200 Success
    - 404 Not Found
</br>

#### POST /users  
    Reister a new user 
    Accept: application/json

    Body: {}

    Response:
    - 201 Created
    - 400 Invalid Request
</br>

#### POST /user  
    User log in 
    Accept: application/json

    Body: {}

    Response:
    - 200 Success
    - 400 Invalid Request
</br>

#### GET /user  
    Confirm user registration 
    Accept: application/json

    Body: {}

    Response:
    - 200 Success
    - 400 Invalid Request
</br>

#### POST /users  
    Resend confirmation 
    Accept: application/json

    Body: {}

    Response:
    - 200 Success
    - 400 Invalid Request
</br>

#### GET /users/{userid} 
    Handles forgot password 
    Accept: application/json

    Body: {}

    Response:
    - 204 No Content
    - 400 Invalid Request

</br>

#### DELETE /users  
    Delete user recordsfrom database  
    Accept: application/json

    Body: {"userid": <user_id>}

    Response:
    - 204 No Content
    - 404 Not Found
