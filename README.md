# Solver

Think the "official" word list is missing a valid four letter word? Create an issue/merge request!  

## How does this work?
- Creates a bi-directional graph with the words list (either provided or the default set) by linking words that are one letter off from one another
    - With list: 'head', 'hear', 'heat', 'mead', 'meat', this graph would be created:
    ```
    head --- meat
      |       |
    head --- mead
      |
    hear
    ```
- Once the graph is created, the shortest path between `from` and `to` is found. This is done by the `github.com/yourbasic/graph` package

### IRIFWMTM
(Interesting Reads I Found While Making This Module) 
- https://blog.golang.org/package-names
- https://blog.golang.org/publishing-go-modules

### Learnings
- When you make a module, `go mod init` when you first make the project, and then `go mod tidy` every once in while when you have removed/added/changed your deps, this auto updates it.
- After that, run `go build` to create a go sum
- Run `go test` to run the tests

### Art Gallery
```
       ______  _____   __     _   _    ____   ____
      / ----,' / _  / / /    / / / / / ___/ /  o  )
      ---,  / / // / / /__  / /_/ / / __/  / ___ i
     /_____/ /____/ /____/ |_____/ /____/ /_/  /_/
     
```
