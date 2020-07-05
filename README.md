# Go

To install Go in windows go to [the official page](https://golang.org/dl/)
To install in ubuntu run

```bash
sudo snap install go --classic
```

---

## Commands

To compile the file run

```bash
go build FILE.go
```

To compile and execute run

```bash
go run FILE.go
```

To create and export one module

```bash
go mod
```

---

To access at the memory address

```go
&x
```

To access at the value of the memory address

```go
*x
```

# Go modules

Create a module of your code of go and put it in github

---

## To create the go module

Create the project with the name of the package equals to the name of the repository of github

Then run the code for create the file of dependencies for the project (change USERNAME and REPOSITORY for your username and the name of the repository of github)

```bash
go mod init github.com/USERNAME/REPOSITORY
```

Init the repository and then create a tag for the module with (change VERSION for the version of the module)

```bash
git tag VERSION
```

Make a push of the tags with

```bash
git push --tags
```

---

## To use the module

To use the method of the module have to has upper case initial

Create a project with

```go
package main

func main() {

}
```

Then run the command in bash (replace PACKAGE_NAME for the name of the package that you want)

```bash
go mod init PACKAGE_NAME
```

Open the go.mod file and add

```
require github.com/USERNAME/REPOSITORY VERSION
```

Then run in bash

```bash
go run main.go
```

Then from the file of go import the module

```go
import "github.com/USERNAME/REPOSITORY"
```

And it's ready to use
