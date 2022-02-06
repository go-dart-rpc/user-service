# user-service

How to run
```
go run main.go
```

Proto model in use
```
string id = 1;
string first_name = 2;
string last_name = 3;
string email = 4;
string gender = 5;
```

Create
```
grpcurl -plaintext -d '{"user":{"firstName":"Bob","lastName":"Marley","email":"dontworry@be.happy","gender":"M"}}' localhost:50051 proto.UserService.CreateUser
```

Read
```
grpcurl -plaintext -d '{"userId":"61ff591f128f597086ac0c4c"}' localhost:50051 proto.UserService.ReadUser
```

List
```
grpcurl -plaintext localhost:50051 proto.UserService.ListUser
```


Update
```
grpcurl -plaintext -d '{"user":{"id":"61ff591f128f597086ac0c4c","firstName":"Bob","lastName":"Marley","email":"dontworry@be.happy","gender":"M"}}' localhost:50051 proto.UserService.UpdateUser
```

Delete
```
grpcurl -plaintext -d '{"userId":"61ff5de24f01624d063397c3"}' localhost:50051 proto.UserService.DeleteUser
```