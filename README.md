# How to start the server.

### 1. Install Goose migration tool:
```
$  go install github.com/pressly/goose/v3/cmd/goose@latest
```
### 3. Run database migration command:
```
$  make migrate-up
```
### 3. Run the golang server:
```
$  make start
```

------------

#### Software used in this project
1. Golang
2. SQLite
3. Gorm
4. Make
5. Bootstrap framework.
6. Auth0 for authentication.
