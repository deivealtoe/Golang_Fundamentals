### Golang - The start

[fmt package documentation](https://pkg.go.dev/fmt)

- Package ```fmt``` implements formatted I/O with functions analogous to ```C```'s ```printf``` and ```scanf```. The format 'verbs' are derived from ```C```'s but are simpler
- ```go run hello.go``` compile and executes
- We don't need to use semicolon in go ```;``` < nope
- A go app needs to have a main func, othherwise it wont compile
- We can concatenate strings using just a comma ```,```. Hence ```fmt.Println("Hello", "World!")```
- Go has type inferente. We don't need to put a type to a variable (it is good though)
  - In case of numbers, the type will be ```float64```
- To declare a variable we have two options
  - With the "normal" way
    - ```var name string = "Alex"```
  - With the simpler > Declare a variable and infer its type
    - ```name := "Alex"```
- To receive a input from the keyboard we need to
  - Declare the variable ```var choice int```
  - Get value from keyboard ```fmt.Scanf("%d%", &choice)```
    - ```%d``` is the type expected, integer of base 10. See doc
    - ```&``` specify the address of the variable in memory
- With the ```Scan``` we don't need to specify the type received. It would be like this
  - ```fmt.Scan(&choice)```
- At if statements, we can use only boolean values
- We don't need to use ```break``` in switch cases, we can use though..
- String comparissons are case sensitive
- In Go, functions can return 2 values
  - If we don't want to receive some of the values, we can ignore it using underscore ```_```. It tells Go to ignore it
  - ```
    func main() {
        name, age := ReturnNameAndAge()
  
        fmt.Println("Hello " + name + "! You are", age, "years old, right?")
    }

    func ReturnNameAndAge() (string, int) {
        name := "Silvia"
        age := 26

        return name, age
    }
    ```
- Go doesnt have ```while``` '-' it only has ```for```
  - If we need a infinit loop, eu can use ```for {  }```
  - It will run until eu ```ctrl + c``` it
- Arrays in go have a fixed size ```var statesArray [4]string```. We can assign a value like so ```statesArray[3] = "ES"```
- ```slice ``` is a dynamically-sized, flexible view into the elements of an array. Day-to-day, slice is more common than arrays
  - ```
    package main

    import "fmt"

    func main() {
        primes := [6]int{2, 3, 5, 7, 11, 13}

        var s []int = primes[1:4]
        fmt.Println(s)
    }
    ```