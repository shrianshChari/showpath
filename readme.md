# showpath
Simple way to view the paths in your `$PATH` written in Go.

### Why?
If you enter in your terminal:
```bash
echo $PATH
```

then you'll get a bunch of folder paths separated by colons, which isn't the easiest to read. This program will separate these folder paths and list them out for you to easily be able to view them.

### Usage
You should have [Go](https://go.dev/) installed on your system, and should be in the project's root directory.
```bash
# Just execute the Go file
go run main.go

# OR build the executable and execute that
go build
./showpath
```

You can also alter its behavior with different flags:
`-r` - Colors each path to make each one more visually distinguishable
`-n` - Numbers each path so keep track of how many paths are in $PATH.

### Limitations
Doesn't take into account paths that have a colon in it; will have to do testing to see how that works.
