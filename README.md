# user-service

A gRPC-based user service built with Go and MongoDB.

## Security

This repository follows security best practices and has automated security scanning enabled. See [SECURITY.md](SECURITY.md) for details.

## How to run
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

## Repository Settings Required

The following security settings should be configured by a repository administrator:

### 1. Branch Protection
Navigate to Settings → Branches → Branch protection rules and configure:
- ✅ Require pull request reviews before merging
- ✅ Require status checks to pass before merging
- ✅ Require signed commits
- ✅ Include administrators

### 2. GitHub Advanced Security
Navigate to Settings → Code security and analysis and enable:
- ✅ Code scanning (CodeQL analysis)
- ✅ Secret scanning
- ✅ Dependency review

### 3. Security Features
- ✅ Dependabot alerts (already enabled)
- ✅ Dependabot security updates (already enabled)

These settings are required to fully comply with SLSA v1.2-rc2 security requirements.