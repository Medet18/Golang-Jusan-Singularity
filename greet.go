package greet

import (
	"fmt"
)

/*func main() {
/*fmt.Println("Hello world")*/
/*greet("Bender", 4)*/
/*text := greet("Bender", 4)
	fmt.Println(text)
}*/

func Greet(name string, table int) string {
	/*fmt.Printf("Hello Bender!, Your table is S\n")
	fmt.Printf("Hello %s!, Your table is %d\n", name, table)
	fmt.Sprintf()*/
	return fmt.Sprintf("Hello %s! Youur table is %d", name, table)
}
