package main

import (
	"fmt"

	"github.com/mrbns/valgo/lib/v"
)

func main() {
	schema := v.NewPipesBuilder(
		v.Entry("email").StringPipe("nazmul", v.IsEmail()),
		v.Entry("passowrd").StringPipe("A strong password is good"),
	)

	v.NewPipesMap(map[string]v.Pipe{
		"email": v.NewStringPipe("is this a valid email"),
	})

	err := schema.Validate()
	if err != nil {
		fmt.Printf("%v", err)
	}

}
