# Linked List

## Run the project

```bash
go run .
```

## Run the tests

- All tests with coverage

```bash
go test -v -cover ./...
```

- Show coverage page

```bash
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```

## Godoc

- Install godoc

```bash
go install golang.org/x/tools/cmd/godoc@latest
```

- Add godoc to the project

```bash
go get golang.org/x/tools/cmd/godoc
```

- Run it

```bash
godoc -http=:8080
```
