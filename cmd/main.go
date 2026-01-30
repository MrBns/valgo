package main

import (
	"fmt"

	"github.com/mrbns/valgo/lib/v"
)

func main() {

	schema := v.NewPipesMap(map[string]v.PipeFace{
		"email":    v.StringPipe("hi@naz.io", v.IsEmail(), v.MaxLength(7, v.ErrMsg("{VALUE} cannot be more than 7 character"))),
		"password": v.StringPipe("HELLOSAbinaYesminIloveYouBabe", v.IsAlpha()),
		"age":      v.IntPipe(10, v.Min(18, v.ErrMsg("must be adult to attend. and {VALUE} is not enough"))),
	})

	err := schema.Validate()

	// for _, e := range err {
	// 	fmt.Println(e.Err)
	// }

	fmt.Printf("%#v", err)

}
