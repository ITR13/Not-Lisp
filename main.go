// notLisp project main.go
package main

import (
	"fmt"
	"strings"

	"github.com/ITR13/Not-Lisp/interpreter"
)

func main() {

}

func basic() {
	s := `
			()()(    *    )(
			()()(   (*)   )(
			()()(  ((*))  )(
			()()( (((*))) )(
				((*))()()
			)(
				(())
			))(
				()()()(
					()()(   ((((*))))   )(
					()()(  (((((*)))))  )(
					()()( ((((((*)))))) )(
						(*)()()
					)( 
						()()()(
							*()(
								((*))()
							)
						)
					)
					)( (((*)))   )
					)( (((*)))() )
				)
			))(
				()()()(
					()()(   (((((( ))))))   )(
					()()( (   ((((*))))()  ) )(
						()()( (((((*))))) )(
						()()(   (((((( ))))))   )(
						()()( (   ((((*))))()  ) )(
							(((((*)))))()()
						)(   ((((*))))()   )
						)(  ((((( )))))    )
						)( (((((( ))))))() )
					)(  (((((*))))) () )
					)( ((((((*))))))() )
				)
			))(
				()()( (((((((*))))))) )(
				()()(     (((*)))     )(
					(((((((*)))))))()()
				)( ( (((*)))() ) )
				)
			)
		`

	s = strings.Replace(s, "*", "(((((((((())))))))))", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, "\n", "", -1)

	bytes := []byte(s)

	data := interpreter.Parse(bytes)
	fmt.Println(data, interpreter.Count(data))
}
