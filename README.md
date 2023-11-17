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

## Documentation

### Public documentation

The public documentation can be found [here](https://pkg.go.dev/github.com/ericmilaneze/data-structure-and-algorithms-in-go).

#### How to generate the public documentation

- Add a LICENSE to the repo
- Add tags following the semantic versioning pattern (vX.X.X)
- Run the following command:
```bash
GOPROXY=proxy.golang.org go list -m github.com/ericmilaneze/data-structure-and-algorithms-in-go@v0.1.0

# if using Windows add SET in front of it:
# SET GOPROXY=proxy.golang.org go list -m github.com/ericmilaneze/data-structure-and-algorithms-in-go@v0.1.0
```
If using Windows

### Godoc (local)

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
