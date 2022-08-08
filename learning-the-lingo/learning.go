// Always start with package declaration
package learning_the_lingo

// Then do some imports
import (
	"fmt"
	"math/rand"
)

// Because it should all begin here
func learning_the_lingo() {
	fmt.Println("Hello World")
	language_characteristics()
	primitive_variables()
	good_at()
	pointers()
	constants()
	slices()
	structs()
	maps()
}

/*
Longer comments. Ahhhhhhh, I've missed this style of comment blocks.
*/
func language_characteristics() {
	creators := "Google Developers"
	benefits := [6]string{"concurrent by default", "garbage collected", "type safety", "rapid compilation", "fully compiled", "ease of use"}
	random_benefit := rand.Intn(6)
	fmt.Println(creators, " developed Go because it was better at ", benefits[random_benefit])
}

func good_at() {
	applications := [5]string{"web services", "web applications", "task automation", "gui/thick-client", "machine learning"}
	best_applications := [2]string{applications[0], applications[1]}
	fmt.Println("Go is best at ", best_applications)
}

func primitive_variables() {

	// heavy handed
	var i int
	i = 42

	// faster
	var f float32 = 3.14

	// even faster
	// Implicit initialization syntax
	woah := "woah"

	// bools
	boolean := true

	// complex
	complex := complex(3, 4)

	// cool few things here; variable unpacking and the real command
	real, imaginary := real(complex), imag(complex)

	fmt.Println(i, f, woah, boolean, complex, real, imaginary)

}

// Oh hey pointers
func pointers() {

	// Create a pointer to a string
	var me *string = new(string)
	// dereference the pointer to set it
	*me = "Jared Sulzdorf"
	// print out the pointer location
	fmt.Println(me)
	// print out the value
	fmt.Println(*me)

}

func constants() {
	const pi = 3.14159
	fmt.Println(pi)
}

// Slices
func slices() {
	arr := [3]int{1, 2, 3}
	// whatup python syntax
	slice := arr[:]
	fmt.Println(slice)

	// Compiler is now going to manage the size of the array
	new_slice := []int{1, 2, 3}
	// now append - not really effecient, but okay
	new_slice = append(new_slice, 3, 4, 5, 6)
	fmt.Println(new_slice)

	// Yeah, this is all Pythonesque
	small_slice := new_slice[1:2]
	fmt.Println(small_slice)

}

// Maps - nothing special
func maps() {
	// map strings to ints
	mapping := map[string]int{"foo": 42}
	fmt.Println(mapping)
}

func structs() {
	type user struct {
		ID    int
		first string
		last  string
	}
	me := user{ID: 1, first: "Jared", last: "Sulzdorf"}
	fmt.Println(me)
}
