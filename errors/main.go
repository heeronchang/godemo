package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	// controller()

	pkgErrorTest()
}

// func controller() {
// 	err := service()
// 	fmt.Println(err)
// 	fmt.Println("--------")

// 	fmt.Println(errors.Unwrap(err))

// }

// func service() error {
// 	err := dao()
// 	// return fmt.Errorf("third level:%w", err)
// 	return errors.Join(err, fmt.Errorf("third level"))
// }

// func dao() error {
// 	err := db()
// 	// return fmt.Errorf("second level:%w", err)
// 	return errors.Join(err, fmt.Errorf("second level"))
// }

// func db() error {
// 	return errors.New("original err")
// }

func pkgErrorTest() {
	err := baz()

	// fmt.Printf("error: %v", err)
	fmt.Printf("error: %+v", err)
}

func foo() error {
	err := errors.New("original err")
	return errors.Wrap(err, "foo err")
	// return errors.WithMessage(err, "foo err")
}

func bar() error {
	err := foo()
	// return errors.Wrap(err, "bar err")
	return errors.WithMessage(err, "bar err")
}

func baz() error {
	err := bar()
	// return errors.Wrap(err, "baz err")
	return errors.WithMessage(err, "baz err")
}
