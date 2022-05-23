# GO_APP

DB Schema: https://dbdiagram.io/d/628b853bf040f104c17c271a

![](https://github.com/KSahu1705/GO_APP/blob/master/db.png)

- User Payment have not implemented it but logic will be same as that of address one.
- To keep the work flow simple didn't implement hashing on password and credential.
- Have provided the ID manually for now.
- Done with API testing using Postman.

### TO DO:
- UNIT Testing Using MockGen.

### API
```Golang
// Routing for handling the projects
a.Get("/users", a.GetAllUser)
a.Get("/users/{id}", a.GetUser)
a.Get("/users/{id}/address", a.GetUserAddress)
a.Post("/users", a.CreateUser)
a.Post("/users/{id}/add_address", a.CreateUserAddress)
a.Put("/users/{id}/update_user", a.UpdateUser)
a.Put("/users/{id}/{addr_id}/update_address", a.UpdateUserAddress)
a.Put("/users/{id}/disable", a.DisableUser)
a.Put("/users/{id}/enable", a.EnableUser)
a.Delete("/users/{id}", a.DeleteUser)
a.Delete("/users/{id}/{addr_id}", a.DeleteUserAddress)
```

### RUN:

`go run main.go`