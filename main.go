package main

import (
	"fmt"
	"products/router"
	"regexp"
)

func main() {
	router.StartRouter()

	fmt.Println(regexp.QuoteMeta(`SELECT id, email, password from users WHERE email = $1`))
}
