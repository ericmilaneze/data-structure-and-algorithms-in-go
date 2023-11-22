# Data Structures & Algorithms In Go

This repo was created to add examples of data structures and algorithms related or not to them.

It is not supposed to be totally based only on the [JavaScript Data Structures & Algorithms + LEETCODE Exercises](https://www.udemy.com/course/data-structures-algorithms-javascript) course, but this course is one of the main inspirations to start this repo. There's [this other](https://github.com/ericmilaneze/data-structures-and-algorithms-plus-leetcode-exercises) repo that is completely based on this course (all examples and exercises written in JavaScript).

## Data Structures

* [Linked Lists](./linkedlist/linkedlist.go)
    * [Find middle node](./linkedlist/find-middle-node.go)
    * [Has loop](./linkedlist/has-loop.go)
    * [Find kth node from end](./linkedlist/find-kth-node-from-end.go)
    * [Partition list](./linkedlist/partition-list.go)
    * [Remove duplicates](./linkedlist/remove-duplicates.go)
    * [Binary to decimal](./linkedlist/binary-to-decimal.go)
    * [Reverse between](./linkedlist/reverse-between.go)

## Run the project

The main file is not doing much, but there are some examples running on it.

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
- Add tags following the semantic versioning pattern (vX.X.X, e.g. v0.0.1)
- Run the following command:
```bash
GOPROXY=proxy.golang.org go list -m github.com/ericmilaneze/data-structure-and-algorithms-in-go@v0.1.0

# if using Windows add SET in front of it:
# SET GOPROXY=proxy.golang.org go list -m github.com/ericmilaneze/data-structure-and-algorithms-in-go@v0.1.0
```

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
