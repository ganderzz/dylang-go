Basic compiler written to learn Golang.

Try it out: `go run *.go ./test/test.dyl`


Currently supports:
  - Tokenizing:
    - Alpha words
    - Numbers
    - Assignment
    - "let" keyword 
  - Parsing
  - Generating to valid javascript 
    - `let apples = 42` -> `(function() {var apples=42})()`