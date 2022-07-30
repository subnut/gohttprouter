package gohttprouter

import "fmt"

func ExampleNew() {
	router := New()
	fmt.Println(router)

	print := func() {
		fmt.Println(
			router.config.RedirectEmptySegments,
			router.config.RedirectTrailingSlash,
			router.config.PreserveEmptySegments,
			router.config.PreserveTrailingSlash,
		)
	}

	print()
	router.config.PreserveTrailingSlash = true
	print()
	router.config.RedirectEmptySegments = false
	print()

	// Output:
	// &{{false true false true}}
	// true true false false
	// true true false true
	// false true false true
}
