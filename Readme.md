# How to run Go scripts?

Just type the following command: ```go run {file_name}.go```

# How to generate a Go scripts build?

Just type the following command: ```go build {file_name}.go```

# How to declare a variable in Go?

- Basic way to declare a variable in go is var + variable name + variable type = start value, example: ```var delta int = 10```
- Another way is to use the ```:=``` character. This will generate a variable with inferred type based on the value, SO YOU CANNOT DECLARE THIS TYPE OF VARIABLE WITHOUT A INITIAL VALUE.
- **IMPORTANT NOTE**: the var type can be used outside function but variables declared with ```:=``` can't!
- Declaring multiple variables ```var a, b, c, d int = 1, 3, 5, 7``` or using 
```
var (
    a int 
    b int = 1
    c string = "hello"
)
```

# How to declare an array in Go?

- Basic way to declare an array var + array_name + = + [length]datatype{values}, example ```var letters = [10]int{1, 2, 3, 4, 5}```

# How to make an IF/else statement?

```
    if x > y {
        fmt.Println("x is greater than y")
    } else if x < y {
        fmt.Println("x is greater than y")
    } else {
        fmt.Println("x and y are equal")
    }
```

# How to make an IF/else statement?

```
day := 4

switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  default:
    fmt.Println("Sunday")
}
```

# How to make a loop in Go?

```
adj := [2]string{"big", "tasty"}
fruits := [3]string{"apple", "orange", "banana"}

for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
}
```

# How to make a function in Go?

- func + {func_name} ({parameters}) (returns)

```
func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}
```

# Go variable naming rules:

- A variable name must start with a letter or an underscore character (_)
- A variable name cannot start with a digit
- A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
- Variable names are case-sensitive (age, Age and AGE are three different variables)
- There is no limit on the length of the variable name
- A variable name cannot contain spaces
The variable name cannot be any Go keywords

# How to add comments in Go?

Go just use the common character to add comments. The "//" for single lines and "/**/" for multi line comments

# Go language major characteristics
- Go is strong typed language and the common types are, int, float32, string and bool.
- Go is a compiled language not interpratated
- Variables with no initial values like ```var a string``` will have a default value, for strings will be "", for int will be 0 and for bool will be false

# Some interesting points of Go language

- Go like C# need to have a main function that will be the first executed method
- If you have unused libraries Go lang removes than once you save the file
- Unused vaiable triggers syntax error and prevent you to run the code
- Go add semicolons (;) by defaults so it not necessary to add then
- None line in go can start with the character "{", since it throws a syntax error on your code. So method declaration cannot be like this: 
``` 
func main() 
{

}
```

it need to be like this:

``` 
func main() {
    
}
```