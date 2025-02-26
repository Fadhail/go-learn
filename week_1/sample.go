package main

import "fmt"

// func main() {
// 	fmt.Println("Hello, World!")
// }

// Variable
// func main() {
// 	var a string
// 	var b int
// 	var c bool

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// Variable 2
// func main() {
// 	var student string
// 	student = "John"

// 	fmt.Println(student)
// }

// Variable 3
// func main() {
// 	var student1, student2, student3 string = "John", "Doe", "Smith"

// 	fmt.Println(student1)
// 	fmt.Println(student2)
// 	fmt.Println(student3)
// }

// Variable 4
// func main(){
// 	var i = 12
// 	var j = "John"

// 	fmt.Printf("Value i : %v and Type i : %T\n", j,j)
// 	fmt.Printf("Value j : %v and Type j : %T\n", i, i)
// }

// Variable 5
func main() {
	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)   /* Prints the value in the default format */
	fmt.Printf("%#v\n", i)  /* Prints the value in Go-syntax format */
	fmt.Printf("%v%%\n", i) /* Prints the % sign */
	fmt.Printf("%T\n", i)   /* Prints the type of the value */

	fmt.Printf("%v\n", txt)  /* Prints the value in the default format */
	fmt.Printf("%#v\n", txt) /* Prints the value in Go-syntax format */
	fmt.Printf("%T\n", txt)  /* Prints the type of the value */

	fmt.Printf("%b\n", i)   /* Base 2 */
	fmt.Printf("%d\n", i)   /* Base 10 */
	fmt.Printf("%+d\n", i)  /* Base 10 and always show sign */
	fmt.Printf("%o\n", i)   /* Base 8 */
	fmt.Printf("%O\n", i)   /* Base 8, with leading 0o */
	fmt.Printf("%x\n", i)   /* Base 16, lowercase */
	fmt.Printf("%X\n", i)   /* Base 16, uppercase */
	fmt.Printf("%#x\n", i)  /* Base 16, with leading 0x */
	fmt.Printf("%4d\n", i)  /* Pad with spaces (width 4, right justified) */
	fmt.Printf("%-4d\n", i) /* Pad with spaces (width 4, left justified) */
	fmt.Printf("%04d\n", i) /* Pad with zeroes (width 4) */
}
