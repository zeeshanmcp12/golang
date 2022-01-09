# Notes

Notes are written for my own understanding and not for others

## 01Into

- Create directory named 01Into
- Open this in "Integrated terminal"
- Execute "go mod init hello" -> it will initialize a module named "hello" as well as create go.mod file.
- Create main.go file.
- Write "package main" (on top of file)
- to print the hello world:
  - initialize function with:
        func main(){ fmt.Println("Hello World")}
  - fmt is a package in golang
  - Println is a method
  - adding fmt in function will import "fmt" package (in golang, we don't need to import packages manually)
- go help -> for go documentation and find out the command etc.
  - go run help -> it will specifically come up with the help related to "run" command
  - go env GOPATH -> to know more about env variable for go.
  - environment variable path for windows: C:\Users\%USERPROFILE%\go
- lexer
  - is who knows the grammer of language like where to put semicolon and where it should not etc.
  - this is something who can add/remove semicolon at the end of the line. for example:
    - we can put semicolon at the end of line 6 (before curely braces) in main.go file but lexer will remove that semicolon because we've installed intellisence from go language.
- types
  - case insensitive - almost everything in go language
  - Everything is type in golang.
  - first initial letter of type gives idea about whether the type is public or private. for example, Println function.
  - const LoginName string -> this capital 'L' in variable (LoginName) declares that it is public and available across the files and program.
  - This Println function was written in a way that it can be used publically. Some other functions that help to this Println are not exposed publically.
  - Similar to other languages, go also has types for example, string, bool, integr, float, complex
  - some advance types: array, slices, maps, structs, pointers
- User Input
  - to take an input from user (or from keyboard), in golang we need to have some packages which are bufio and os.
  - buffer i/o will help to buffer the i/o and os package will help to get the input from user (or from keyboard)
  - syntax:
    - reader := bufio.NewReader(os.Stdin)
      - reader -> variable
      - := -> operator for assignment and declaration
      - bufio -> go package
      - NewReader -> method of bufio package -> public because of first letter is capital
      - os -> go package
      - Stdin -> method of os package.
    - As soon as we save the file with above syntax, it will import both packages (bufio and os)
    - till above syntax, a process is going on to read the input (is syntax tak koi na koi process read kar raha hai.)
  - comma ok syntax || err err
  - syntax
    - input, _ := reader.ReadString('\n')
      - input -> whatever an input will be given
      - _ -> (underscore) means, if any mistake/error occurs during input so this is something try catch.
      - reader -> is a variable we declared above
      - .ReadString -> is a method which will wait (or read) for string we gonna input.
      - ('\n') -> represents new line so reader.ReadString will read for string till we press "Enter".
