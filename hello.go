package main // Every GO program is made of packages, so we define the main code as the main package so it can run the main function

import (
	"fmt" //fmt stands for format, it allows you to use inputs and outputs of values.
	"time"
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
	// var y int
	// fmt.Scanln(&y)

	// t := Timer{"timer1", 0}

	// for i := 0; i < y; i++ {
	// 	t.tick(i)
	// }

	/** Arrays! **/

	// var a [5]int //Creating but not initiliazing the array (by default, the elements of the array are initilized to the zero value of the given type)
	// a = {0, 2, 3, 4, 5} //This won´t work
	// a[0] = 8 //This is how you populate an non initilized array
	// a[1] = 42
	// arr := [5]int{0, 2, 4, 6, 8} //creating and initializing the array
	//? An Array has a fixed size. Which means that once declared, you cannot increase or decrease its size.
	//? Luckily GO provides the *slice*, which is a dynamically-sized view into the elements of an array.
	//? A slice is based on an array and is defined by specifying two indices, a low and high, separated by a clon:
	// var s []int = arr[1:3] //This code selects the elements with index 1 to 3 from the a array, including the first given idnex, but excluding the last

	// fmt.Println(arr, a[0])
	// fmt.Println(s) // so, the slice s will now include the values [2,4]
	/**
	*?You can omit the low or the high bound. Omitting the low will take the value 0, while omitting the high, will take the arr lenght.
	*? For example, arr[:3] will take the first 3 elements of the array
	**/
	//We can also access a slice value the same way we did with the common array:
	// fmt.Println(s[1])
	//And since the slice don't store any value, just describe a section of an underlying array, if you try to, you'll change the original array
	// fmt.Println(arr[1]) // returns 2
	// s[0] = 3
	// fmt.Println(arr[1]) // returns 3
	/**Now, Go provides a make() function to create slices. This is how you create dynamically-sized array.**/

	// brr := make([]int, 5) //The make function creates an array of the given type and size, and returns a slice of that array
	//after creating the slice, we can add append new elements to it using the append() function
	// brr = append(brr, 8)
	// fmt.Println(brr) // It'll return [0, 0, 0, 0, 0, 8] since it'll return the old slice + the new one
	// x := make([]int, 2)
	// x = append(x, 3, 6, 2)
	// fmt.Println(x[4])                                          // It'll return 2
	// fmt.Println(x)                                             // It'll return [0 0 3 6 2]
	/** Range! **/
	//? the range form of the for loop allows you to iterate over a slice:

	// a := make([]int, 5)
	// a[1] = 2
	// a[2] = 3

	// for i, v := range a {
	// 	fmt.Println(i, v) // it'll return the index and the value stored inside.
	//So...it'll return 0 0, 1 2, 2 3, 3 0, 4 0
	// }
	//? Now, sometimes you'll want only the value, in that case, skip the index using an underscore
	// for _, v := range a {
	// 	fmt.Println(v)
	// }
	// The range can also be used to iterate over the charactersof a string.
	// x := "Hi, mom"
	// for _, c := range x {
	// 	fmt.Println(c) //Prints unicode code points of the character.
	// }
	// fmt.Println("")
	// for _, c := range x {
	// 	fmt.Printf("%c \n", c) //Prints the character.
	// }
	// res := 0
	// nums := [3]int{2, 4, 6}

	// for i, v := range nums {
	// 	if i%2 == 0 {
	// 		res += v
	// 	}
	// }
	// fmt.Println(res)
	/** MAP MAP MAP MAP MAP MAP MAAAAAAAAAAAAN **/
	//m := make(map[string]int)
	// We can create a map using the make() function, similar to arrays.
	//Maps are used to store key:value pairs, the key is always unique
	// m["James"] = 42
	// m["Amy"] = 24

	// fmt.Println(m["James"])
	/* We can also use this syntax to define a map */
	// m := map[string]int{
	// 	"James": 42,
	// 	"Amy":   24}
	// fmt.Println(m["Amy"])
	// we can use delete to remove an element from the map!
	// delete(m, "Amy")
	//to assign a new element to the map do this:
	// m["Tim"] = 23
	//? Maps are printed in the form map[key:value key:value] when output with fmt.Println().
	// m := map[int]int{
	// 	8: 42,
	// 	2: 6,
	// 	4: 9,
	// 	5: 3}
	// delete(m, 2)
	// fmt.Println(m[4] - m[5])                                   // it'll return  6
	/** Variadic Functions **/
	//? Variadic functions are the ones that can be called with any number of arguments, like Println
	// neoSum(3, 4, 5, 1, 2, 12)
	// neoSum(2, 2)
	// neoSum(-2, 5)

	// v := []int{8, 5, 3}
	// f(v...)

	// s := []int{1, 2, 4, 6, 8}
	// s[2] = s[1]
	// s[3] = s[2] + s[0]
	// fmt.Println(s[4] % s[3])

	/** Match Results **/
	/*
		? You are making a program to analyze sport match results and calculate the points of the given team.
		? the match results are stored in an Array called results.
		? Each match has one of the following results:
		*? "w" - won
		*? "l" - lost
		*? "d" - draw

		? A win adds *three* points, a draw adds one and a lost match does not add any points.
		? Your program needs to take the last match results as input and append it to the results array.
		? After that, calculate and output the total points the team gained from the results.
	*/

	// results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	// var i string
	// fmt.Scanln(&i)
	// results = append(results, i)

	// matchResults(results)

	/* Concurrency :
	?> Concurrency means multiple computations are happening at the same time.
	?> It is used when your program has multiple things to do.
	*Concurrency is about creating multiple processes excecuting independently.
	?> In order to use concurrency, the program is broken into parts, which are then executed separately.
	*/
	//?> To achieve this, GO provides Goroutines. It is pretty much like a cpu thread to accomplish multiple taks, but it consumes way less
	//?> resources than a normal cpu thread created by the OS.
	// out(0, 5)
	// out(6, 10)
	/**
	the out() function simply outputs numbers in the given range. We use a time.Sleep() to emulate work being done between outputs
	It simply waits for the provided time (50ms) before continuindo each iteration
	Now! If we call the function in main two times, the first call will execute first follow by the second call.
	This will generate the output of 0 to 5, then 6 to 10
	*/
	//?> This is called a sequential program, as the statements are executed one after the other. The first call needs to be complete
	//?> before the second call starts.

	//?> When running concurrent programs, do not often want to wait for one task to finish before starting a new one.
	//?> To achieve concurrency, let's start the function calls as Goroutines, using the go keyword:
	// go out(0, 5)
	// go out(6, 10)

	/**


	//if you run that part of the program, you'll notice something amazing about the outputs! (there's no output.) but...why?
	//?> The output happens like that because the main() function exits before the Goroutines complete.
	//*Our program has 3 virtual threads, the 2 function call and a main(). Our 2 funcs get executed concurrently, and main() doest not wait
	//* for them to finish.
														//	Channels
	//?> a channel is like a pipe, allowing you to send and receive data from goroutines, and enbaling them to communicate and synchronize.
	**/
	//?> To use a channel, we first need to make one using the make() function:
	//ch := make(chan int) // The type after the *chan* keyword indicates the type of the data we'll send through the channel.
	//ch <- 8              // We can send data to the channel using this syntax.
	//* Similarly, we can receive data from the channel using the following syntax:
	//value := <-ch
	//* and if we do not need the value as a variable, we can simply use:
	//<-ch //* data flows in the direction of the arrow.
	//* Now we can even use our out function without the time.Sleep in main()

	// ch := make(chan bool)

	// go out(0, 5, ch)
	// go out(6, 10, ch)
	// <-ch

	//* The receive operation blocks the code until and unless some data is sent by the send operation.
	//* If nodata is receive, a deadlock will occur, blocking the code from executing.
	//* Now we'll send data from a function to the main function through the channel
	// evenCh := make(chan int)
	// sqCh := make(chan int)

	// go evenSum(0, 100, evenCh)
	// go squareSum(0, 100, sqCh)

	// fmt.Println(<-evenCh + <-sqCh) //* If you do not need to send data to a channel anymore, you can close it using close(ch)
	//*Where ch is the name of the channel. This is done in the sender.

	/** Select **/
	//?> The select statement is used to wait on multiple channel operations
	//* The syntax is similar to switch except that each of the case statements will be a channel operation.

	// evenCh := make(chan int)
	// sqCh := make(chan int)

	// go evenSum(0, 100, evenCh)
	// go squareSum(0, 100, sqCh)

	// select {
	// case x := <-evenCh:
	// 	fmt.Println(x)
	// case y := <-sqCh:
	// 	fmt.Println(y)
	// }
	/**image.png
	**Combining goroutines and channels with select is a powerful feature of Go. Imagine a program that needs to execute some code whenever
	**one of the concurrent operations completes -- this can be achieved using select.
	**/
	//?> a select can have a default case, which will execute when no channel is ready. For example, we could have an infinite for loop,
	//?> waiting for one of the channels to receive data:
	// for {
	// 	select {
	// 	case x := <-evenCh:
	// 		fmt.Println(x)
	// 		return
	// 	case y := <-sqCh:
	// 		fmt.Println(y)
	// 		return
	// 	default:
	// 		fmt.Println("Nothing avaiable")
	// 		time.Sleep(50 * time.Millisecond)

	// 	}
	// }
	/**
	**A select statement blocks until at least one of its cases can procced.
	** The default case is useful in preventing deadlocks -- without it the select would wait for a channel forever, crashing the program
	** if none of the channels received data.
	**/
	ch := make(chan int)
	go chTest(ch)
	fmt.Println(<-ch / 2)
	fmt.Println("Uncomment my code to see what it does, duh! (NOT AT ONCE, PLEASE, IT'LL CONFLICT ALL THE WAY DOWN)") //strings must be encapsulated by ' " ' like C language.
	// fmt.Println((2022 - 1990) > (2050 - 2022)) // This returns true.

	// time.Sleep(500 * time.Millisecond) //Using this to observe that each coroutine worked independent and concurrently.
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

func out(from, to int, ch chan bool) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
	}
	ch <- true
}

func evenSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i
		}
	}
	ch <- result
}

func squareSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i * i
		}
		ch <- result
	}
}

func chTest(ch chan int) {
	ch <- 4
}

/** Variadic Functions **/

func neoSum(nums ...int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)
}

func f(v ...int) {
	res := 20
	for _, a := range v {
		res -= a
	}
	fmt.Println(res)
}

func matchResults(matches []string) {
	res := 0
	for _, v := range matches {
		switch v {
		case "w":
			res += 3
		case "d":
			res += 1
		default:
			res += 0
		}
	}
	fmt.Println(res)
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
