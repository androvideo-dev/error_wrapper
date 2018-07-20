# Error Wrapper
![](https://travis-ci.org/daniel-androvideo/error_wrapper.svg?branch=master)

> error wrapper for golang

## Install

```bash
go get -u github.com/daniel-androvideo/error_wrapper
```

## How to use

```golang
package main

import (
	"github.com/daniel-androvideo/error_wrapper"

	"errors"
	"fmt"
)

func main() {
	w := error_wrapper.ErrorWrapper{}
	w.Do(func() error {
		fmt.Println("task 1")
		return nil
	})
	w.Do(func() error {
		fmt.Println("task 2")
		return errors.New("oops")
	})
	w.Do(func() error {
		fmt.Println("task 3")
		return nil
	})

	fmt.Println(w.GetErr())
}
// result
// task 1
// task 2
// oops
```
