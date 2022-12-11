# fridgiGO

> This is the backend of fridgiGO. This backend was developed in the programming language Go (Golang).

## Below is the list of developed API endpoints

### **_Ping_**

| Route          | HTTP Request | Body | Description     |
| :------------- | :----------: | :--: | :-------------- |
| `/v1/api/ping` |    `GET`     |      | Ping the server |

### **_Users_**

| Route                            | HTTP Request |                            Body                             | Description                                        |
| :------------------------------- | :----------: | :---------------------------------------------------------: | :------------------------------------------------- |
| `/v1/api/users/login`            |    `POST`    |                     `{email, password}`                     | Authenticate user (login function)                 |
| `/v1/api/users/register`         |    `POST`    | `{first_name, last_name, email, password, password_repeat}` | Register user (register function)                  |
| `/v1/api/users/register-confirm` |    `POST`    |               `{email, confirmation_number}`                | Confirm user (user register confirmation function) |
| `/v1/api/users/user-info`        |    `GET`     |                          `{token}`                          | Get user info                                      |
