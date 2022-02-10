# Golang Food Data REST API 🚀 

🔥 Golang Rest Api with basic JWT Authentication and Basic Crud Operations.

## Technology
- Language (golang)
- Database (mysql)
### Libraries
- Router (gorilla/mux)
- Server (net/http)
- JWT (dgrijalva/go-jwt)
- Password Encryption (bcrypt)
- Database ORM (gorm) 
- Live Reload (cosmtrek/air)

## Features
- [x] Monlith  
- [x] Authentication
- [x] endpoint protection with middleware
- [ ] Authorization - Role based Access Control (RBAC)

## API Documentation

### Authentication
> **POST** ``/auth/login``

Login with username/email and password.

##### Body

```json
{
    "id": "abc123",
    "password": "abc123",
}
```

#### Output

```json
{
    "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
    "user": {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "name": "ABC 123",
        "username": "abc123",
        "email": "admin@abc123.io",
    }
}
```

> **POST** ``/auth/signup``

Create a new user in the database.

##### Body

```json
{
    "name": "ABC 123",
    "username": "abc123",
    "email": "admin@abc123.io",
    "password": "abc123",
}
```

#### Output

```json
{
    "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
    "user": {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "name": "ABC 123",
        "username": "abc123",
        "email": "admin@abc123.io",
    }
}
```

### Data Manipulation

All endpoints are protected, must send valid **jwt** as ``Authorization`` header with each request.

> **GET** &nbsp; ``/food/all``

Get All Food Items

#### Output

```json
[
    {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "name": "Apples",
        "quantity": 100,
        "selling_price": "100 USD",
    },
    {
        "id": "123e4567-e89b-12d3-a456-426614174000",
        "name": "Mangos",
        "quantity": 97,
        "selling_price": "120 USD",
    }
]
```

> **GET** &nbsp; ``/food/<name>``

Get single Food Item by its name. name should be lowercase (e.g /food/apples)

#### Output

```json
{
    "id": "123e4567-e89b-12d3-a456-4265674174000",
    "name": "Apples",
    "quantity": 100,
    "selling_price": "100 USD",
}
```

> **POST** &nbsp; ``/food``

Add a new food item to the database.

##### Body

```json
{
    "name": "Oranges",
    "quantity": 44,
    "selling_price": "80 USD",
}
```

#### Output

```json
{
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Oranges",
    "quantity": 44,
    "selling_price": "80 USD",
}
```
> **DELETE** &nbsp; ``/food/<name>``

Delete one Food Item from the database.

#### Output

```json
{
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Oranges",
    "quantity": 44,
    "selling_price": "80 USD",
}
```