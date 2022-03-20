dl(Debug x Log) - The instant logger package for debug
======

[![Go Reference](https://pkg.go.dev/badge/github.com/task4233/dl.svg)](https://pkg.go.dev/github.com/task4233/dl)
[![.github/workflows/ci.yml](https://github.com/task4233/dl/actions/workflows/ci.yml/badge.svg)](https://github.com/task4233/dl/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/task4233/dl)](https://goreportcard.com/report/github.com/task4233/dl)
[![codecov](https://codecov.io/gh/task4233/dl/branch/main/graph/badge.svg?token=xrhysp4Tzf)](https://codecov.io/gh/task4233/dl)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)


![delog.gif](https://user-images.githubusercontent.com/29667656/159164178-f72aede7-f825-438a-add6-aa3deedf8c4c.gif)

## Description
Who doesn't write wrong codes? No one.  
Then, programs don't work well, and developers write logs for debugging to understand what happens.

However, some developers forget to delete their logs after resolving the problem and push their codes. In the worse case, the logs might be released.

**dl** is developed to resolve their problems.

## Features
- **Logging package for debug in Go**
- **Command for Sweeping all functions of this package**
- **Command for installing git pre-commit hook**

## Installation
### Go1.18

```bash
$ go install github.com/task4233/dl/cmd/dl@latest
```

## Use Case
### Debug

[Playground](https://go.dev/play/p/GRucgd6JhPk)
```go
package main

import (
	"os"

	"github.com/task4233/dl"
)

type U[T any] []T

func (t U[T]) append(v T) {
	t = append(t, v)
	dl.Info(t)
}

func (t U[T]) change(v T) {
	t[0] = v
	dl.FInfo(os.Stdout, t)
}

func main() {
	t := U[int]([]int{1, 3})
	t.append(5)
	t.change(5)
}


// Output:
// [DeLog] info: main.U[int]{1, 3, 5} (main.U[int]) prog.go:13
// [DeLog] info: main.U[int]{5, 3} (main.U[int]) prog.go:18
```

### Adds dl into pre-commit of Git
1. Please run commands below to install dl in your Git repository.

```bash
$ dl init .
```

2. Just commit

- `delog` is used in the file.

```bash
$ cat main.go 
package main

import (
	"fmt"
	
	"github.com/task4233/dl"
)

func SayHi[T any](v T) {
	dl.Printf("Type: %T, v: %v\n", v, v) // This statement can be removed by `$ dl clean main.go`
	fmt.Println("Hi, ", v)
}

func main() {
    SayHi("hoge")
}
```

- invoke `$ git commit`

```bash
$ git add main.go
$ git commit -m "feat: add main.go"
remove dl from main.go # automatically removed
[master 975ecf9] feat: add main.go
 1 file changed, 12 insertions(+), 21 deletions(-)
 rewrite main.go (91%)
```

- `delog` is removed automatically

```bash
$ git diff HEAD^
diff --git a/main.go b/main.go
index 90a78bd..0e28e8a 100644
--- a/main.go
+++ b/main.go
@@ -1,21 +1,12 @@
 package main

+import (
+       "fmt"
+)
 
+func SayHi[T any](v T) {
+       fmt.Println("Hi, ", v)
+}

 func main() {
+       SayHi("hoge")
 }
```

### Remove dl from pre-commit of Git

```bash
$ dl remove .
```

## Contribution
Please feel free to make [issues](https://github.com/task4233/dl/issues/new/choose) and pull requests.

## Author
[task4233](https://task4233.dev)
