# Developer notes

## Requirements

- Go
- `pkgsite` (for offline viewing of documentation)

## Viewing documentation

```console
# start the documentation server
$ cd <project root>
$ pkgsite -open . &
# use browser to navigate to
# http://localhost:8080/github.com/jspaaks/learn-go-with-tests
```

## Running the tests

```console
$ go test ./helloworld/pkg/helloworld -v
=== RUN   TestSayHello
=== RUN   TestSayHello/saying_hello_to_people_in_dutch
=== RUN   TestSayHello/saying_hello_to_people_in_english
=== RUN   TestSayHello/saying_hello_to_people_in_french
=== RUN   TestSayHello/saying_hello_to_people_in_spanish
=== RUN   TestSayHello/saying_hello_to_people_in_unknown_language
=== RUN   TestSayHello/saying_'Hello,_world'_when_an_empty_string_is_suplied
--- PASS: TestSayHello (0.00s)
    --- PASS: TestSayHello/saying_hello_to_people_in_dutch (0.00s)
    --- PASS: TestSayHello/saying_hello_to_people_in_english (0.00s)
    --- PASS: TestSayHello/saying_hello_to_people_in_french (0.00s)
    --- PASS: TestSayHello/saying_hello_to_people_in_spanish (0.00s)
    --- PASS: TestSayHello/saying_hello_to_people_in_unknown_language (0.00s)
    --- PASS: TestSayHello/saying_'Hello,_world'_when_an_empty_string_is_suplied (0.00s)
PASS
ok      github.com/jspaaks/learn-go-with-tests/helloworld/pkg/helloworld        0.001s
# or,
$ cd helloworld/pkg/helloworld
$ go test -v
=== RUN   TestSayHello
=== RUN   TestSayHello/saying_hello_to_people_in_dutch
=== RUN   TestSayHello/saying_hello_to_people_in_english
=== RUN   TestSayHello/saying_hello_to_people_in_french
=== RUN   TestSayHello/saying_hello_to_people_in_spanish
=== RUN   TestSayHello/saying_hello_to_people_in_unknown_language
=== RUN   TestSayHello/saying_'Hello,_world'_when_an_empty_string_is_suplied
--- PASS: TestSayHello (0.00s)
    --- PASS: TestSayHello/saying_hello_to_people_in_dutch (0.00s)
    --- PASS: TestSayHello/saying_hello_to_people_in_english (0.00s)
    --- PASS: TestSayHello/saying_hello_to_people_in_french (0.00s)
    --- PASS: TestSayHello/saying_hello_to_people_in_spanish (0.00s)
    --- PASS: TestSayHello/saying_hello_to_people_in_unknown_language (0.00s)
    --- PASS: TestSayHello/saying_'Hello,_world'_when_an_empty_string_is_suplied (0.00s)
PASS
ok      github.com/jspaaks/learn-go-with-tests/helloworld/pkg/helloworld        0.001s
```

## Running an executable

```console
$ cd <project root>
$ go run ./helloworld/cmd/helloworld/main.go
Hello, John
# or,
$ cd ./helloworld/cmd/helloworld
$ go run main.go
Hello, John
```
