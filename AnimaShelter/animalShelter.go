package animashelter

import "fmt"

type Speaker interface{
	Speak() string
}

type Dog struct{
	Name string
}

type Cat struct{
	Breed string
}

func(dd Dog) Speak()string{
	return fmt.Sprint("Woof! My name is ", dd.Name)
}

func (cc Cat) Speak()string{
	return fmt.Sprint("Meow! I am a ", cc.Breed)
}

func MakeThemTalk(s Speaker){
	fmt.Println(s.Speak())
}
// animal := make(Speaker{dog, cat})
