package main

import "fmt"

type dog struct{}

type fish struct{}

type bird struct{}

func (dog) walk() string{
	return "dog walk"
}

func (fish) swim() string{
	return "fish swim"
}

func (bird) fly() string{
	return "bird fly"
}

func moveDog(p dog){
	fmt.Println(p.walk())
}

func moveFish(f fish){
	fmt.Println(f.swim())
}

func moveBir(b bird){
	fmt.Println(b.fly())
}

func main() {
	d := dog{}
	moveDog(d)
	f := fish{}
	moveFish(f)
	b := bird{}
	moveBir(b)
}