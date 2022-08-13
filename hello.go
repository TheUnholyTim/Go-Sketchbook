package main // Every GO program is made of packages, so we define the main code as the main package so it can run the main function

import (
	"fmt" //fmt stands for format, it allows you to use inputs and outputs of values.
)

func main() { //  all code that'll run must be inside the main function (which is declared with the reserved word func)
	// fmt.Println("Hello, world!")
	// var input string
	// fmt.Scanln(&input) // Always use & key to return memory adress! (like C)
	// /** but if we use the Scan as it is, it'll only get the first word we type, not a whole ass phrase!**/
	// fmt.Println((input + input))
	// var anotherInput int
	// fmt.Scanln(&anotherInput) // Always use & key to return memory adress! (like C)
	// if anotherInput >= 18 {
	// 	fmt.Println("Allowed!")
	// } else { // Else must be in front of the curly bracket or it will return an error!
	// 	fmt.Println("Not allowed!")
	// }

	// Switch cases
	// /* Switch cases don´t use breaks at the end of any case. Also the case is idented in the same level as the switch statement */

	// num := 2
	// switch num {
	// case 1:
	// 	fmt.Println("One!")
	// case 2:
	// 	fmt.Println("Two!")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println(num)
	// }

	// x := 8
	// switch y := x % 2; y {
	// case 0:
	// 	x -= 1
	// case 1:
	// 	x += 1
	// }
	// fmt.Println(x)
	/*
		Loops! The only loop in GO is the for loop, it is divided in
		init, condition and post statement with no braces
	*/
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// sum := 0
	// for i := 1; i <= 3; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	//In case you're wondering "but what about while loops?" WELL, they don't exist per se, but you can sort of do then like that:
	// sum := 1
	// res := 0
	// for sum <= 1000 {
	// 	res += sum
	// 	fmt.Println(res, sum)
	// 	sum++
	// 	fmt.Println(sum)
	// }
	// fmt.Println(res)

	// x := 7
	// res := 0
	// switch {
	// case x > 8:
	// 	res++
	// case x > 3 && x < 10:
	// 	res++
	// case x > 5:
	// 	res++
	// }
	// fmt.Println(res) // the output will be 1
	// res := 0
	// for i := 0; i < 10; i += 3 {
	// 	res += i
	// 	fmt.Println(res)
	// }
	// fmt.Println(res)

	// line()
	// line()
	// fmt.Println()
	// welcome("Tim")
	// sum(32, 8)
	// whoILove("girl", "")
	// fmt.Println(sum(5, 7))
	// a, b := swap(42, 8)
	// fmt.Println(a, b)
	/* Defering functions */
	//defer welcome("Tim") //defer means that the function will only execute after all the main function is executed.
	//fmt.Println("Hey!")

	/*While defering multiple functions, It'll work first-in-last-out */
	// fmt.Println("Start!")
	// for i := 0; i < 5; i++ {
	// 	defer fmt.Println(i)
	// }
	// fmt.Println("end")
	// fmt.Println(isEven(square(12)))                            //Nested functions

	/* mars age!
			How old would you be on Mars?
	A year on Earth has 365 days, while a year on Mars has 687 days.

	Create a program that takes your age in Earth years as input, and outputs the corresponding age on Mars.

	The given program takes an integer as input and passes it to the mars_age() function as argument.
	*/

	// var age int
	// fmt.Println("Por favor, digite sua idade abaixo: ")
	// fmt.Scanln(&age)

	// mars := mars_age(age)
	// fmt.Println(mars)

	/* Pointers... */

	// var p *int // in Go we declare pointers using a *, in that case, p is a pointer to an integer value
	// x := 42
	// p = &x //now p has the x memory adress value

	// fmt.Println(p)  //it'll return 0xc0000b4000
	// fmt.Println(*p) // it'll return the value stored in the memory adress stored in p

	// *p = 8 // that way we can change the variable stored in that memory adress without messing up with that variable directly.
	// fmt.Println(x)

	// x := 42

	// change(x)
	// fmt.Println(x)

	// change_ptr(&x)
	// fmt.Println(x)
	/*
		Note that change(x) didn´t changed the variable value, that's because change function will only copy the variable into the local scope
		while change_ptr indeed changed x value by targeting xs memory adress directly.
	*/

	// y := 3.141592653589793
	// fmt.Println(y)
	// zero(y)
	// fmt.Println(y) // As you can see, it won't work with float32 and float64 values since the zero function returns an integer pointer

	/* Structs! */ //? Friendly reminder that GO does not support classes.
	// type Contact struct {
	// 	name string //field 1
	// 	age  int    //field 2
	// }
	// x := Contact{"James", 42} // This will create a new Contact
	/* //? Always good to remember that:
	We can also provide the names of the fields when creating a new struct.
	For example:
	x := Contact{name: "James", age: 42}
	*/

	//? We can access the struct fields using struct_name and a dot:

	// x.age = 8
	// fmt.Println(x.age)
	// fmt.Println(x.name)

	//? Of course we can use pointers inside constructs
	// type Contact struct {
	// 	name string
	// 	age  int
	// }
	//? similar to simple pointers, we can also make struct pointers using the & operator

	// x := Contact{"James", 42}
	// p := &x

	//	fmt.Println(*p.name) //? Can´t show struct data like that
	// fmt.Println(p.name) // Pointers to structs are automactically dereferenced, meaning we can acess the field by simply using a dot, no *
	// type Car struct {
	// 	name  string
	// 	color string
	// 	year  int
	// }
	// a := Car{"BMW", "red", 2018}
	// p2 := &a
	// fmt.Println(p2.name, p2.color, p2.year)

	//? We also can use pointers when creating a new struct.

	// type Person struct {
	// 	name string
	// 	age  int
	// }

	// p := &Person{"John", 15} //? Pointers to structs are useful, as they allow you to pass them to functions as arguments.

	// type Dog struct {
	// 	name string
	// 	age  int
	// }

	// x := &Dog{"Max", 8}
	// fmt.Println(x.age)

	//? Methods! //

	//Go to line 294 to see a method, since they can only be declared outside a function...since functions can't be nested inside other functions
	//? After defining the method, we can call it on our struct using the dot syntax, like this:
	// x := Girlfriend{"girl", 19, "taken"}
	//x.greeting()
	//welcome(x)

	/**
	*? In case we need to change the data of the struct in a method, we can use...POINTERS as method receivers! ?>
	**/
	//fmt.Println(x.age)
	//x.increase(2)
	//fmt.Println(x.age)

	/** Created a withdraw method to withdraw cash from Bank Account. **/
	//timAcc := BankAccount{"Tim Kramer", 32000.00, 0.0}
	//timAcc.withdraw(1000.50)
	//fmt.Println()
	/* Also! Formating strings in GO is EXACTLY THE SAME as in C language. */
	//fmt.Printf("Current value in your account is: %f. \n Last value withdrawn from it was: %f\n", timAcc.storedMoney, timAcc.lastTransaction)

	/** Ticking Timer Challenge!
	*? Build a Timer app that should count up to a given number.
	*? Your Program needs to take a number as input and make the TIMER tick that number of times.
	*? The code in main intiliazes a Timer and takes a number as input. Then it calls the tick() method for the Timer the given number of times.
	*? Define the Timer Struct with two fields: id and value, and define the tick() method, which should increment the value by one and output its current value.
	**/
	var y int
	fmt.Scanln(&y)

	t := Timer{"timer1", 0}

	for i := 0; i < y; i++ {
		t.tick(i)
	}

	fmt.Println("Uncomment my code to see what it does, duh!") //strings must be encapsulated by ' " ' like C language.
	// fmt.Println((2022 - 1990) > (2050 - 2022)) // This returns true.
}

/*Functions and stuff */
func line() {
	fmt.Print("----------------")
}

func sum(a, b int) int {
	return a + b //For you to be able to use Return you must specify what sort of data it'll be returning when declaring the function
}

func whoILove(name, gender string) {
	switch gender {
	case "female":
		fmt.Println("The girl I love names's " + name)
	case "male":
		fmt.Println("The boy I love name's " + name)
	default:
		fmt.Println("The person I love name's " + name)
	}
}

func swap(x, y int) (int, int) { //must declare the data type of every value will be receiving back from the function
	return y, x
}

func square(num int) int {
	return num * num
}
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func double(num int) {
	fmt.Println(num * 2)
}
func mars_age(age int) int {
	days := age * 365
	return days / 687
}

/* We can even pass pointers as function args! */
func change(val int) {
	val = 8
}

func change_ptr(ptr *int) {
	*ptr = 8
}

// This function will change any integer value to 0
func zero(x *int) {
	*x = 0
}

func (greet Girlfriend) greeting() {
	fmt.Println("Ohayo, da-arling!")
}

func welcome(x Girlfriend) {
	fmt.Println(x.name)
	fmt.Println(x.age)
	fmt.Println(x.status)
}

/**
*? In the code above, welcome() is a function that takes a struct as its argument.
*
**/

func (ptr *Girlfriend) increase(val int) {
	ptr.age += val
}

func (ptr *BankAccount) withdraw(val float32) {
	ptr.lastTransaction = -val
	ptr.storedMoney -= val

}

func (t *Timer) tick(value int) {
	for i := 0; i <= value; i++ {
		t.value = i + 1
	}
	fmt.Println(t.value)
}

/* Structs */

//? Structs can also be created outisde functions (like the main function) || You should always create them outside the main function!

type Girlfriend struct {
	name   string
	age    int
	status string
}

type BankAccount struct {
	name            string
	storedMoney     float32
	lastTransaction float32
}

type Timer struct {
	id    string
	value int
}
