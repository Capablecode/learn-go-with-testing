package main

import (

	// "learning/hex"
	// "learning/helperfunc"
	// "learning/pointer"
	// "learning/reverse"
	// "learning/article"
	// "learning/lesson"
	// "learning/punctuation"
	// "learning/algorithm"
	// animashelter "learning/AnimaShelter"
	"fmt"
	// "learning/interfaces"
	asciiartweb "learning/ascii-art-web"
	"net/http"
	// "learning/lesson"
	// "learning/sum"
	// "fmt"
)

func main() {
	// fmt.Println(lesson.Factorial_Recursion(5))
	// fmt.Println(lesson.ReverseAString("hello"))
	// fmt.Println(lesson.Hello("Eldee", "Spanish"))
	// fmt.Println(sum.SumAllTails())
	// fmt.Println(algorithm.PrintMaxNum([]int{3, 5, 7, 10, 0}))
	// fmt.Println(reverse.Reverse("hello"))
	// fmt.Println()
	// a := 20
	// b := 40
	// pointer.Swap(&a, &b)
	// fmt.Println(" a = ", a,"\n","b = ", b)
	// x := 2
	// pointer.Double(&x)
	// fmt.Println(x)
	// numSlices := []int{10, 2, 4, 1, 8}
	// fmt.Println(algorithm.SortNumInAscendingOrder(numSlices))
	// fmt.Println(algorithm.SortNumInDescendingOrder(numSlices))
	// fmt.Println(helperfunc.ToUpperCase([]string{"go", "is", "fun"}, 2, 2))
	// result := helperfunc.ToUpperCase([]string{"go", "is", "fun"}, 2, 1)
	// fmt.Println(result)
	// result2 := helperfunc.ToLowerCase([]string{"go", "is", "fun"}, 3, 2)
	// fmt.Println(result, result2)
	// text := "I am exactly how they describe me: 'awesome'"
	// fmt.Println(punctuation.FixPunctuation(text))
	// text := "hello"
	// fmt.Println(helperfunc.Cap(text))
	// text := "1af"
	// fmt.Println(hex.ConvertHexidecimal(text))
	// text := "10"
	// fmt.Println(hex.ConvertBinary(text))
	// text := "SUNDAY"
	// fmt.Println(helperfunc.Lower(text))

	// text := "Elton John said: ' I am the most well-known homosexual in the world '"
	// fmt.Println(helperfunc.FixQuotes(text))
	// text := "There it was. A amazing rock!"
	// fmt.Println(article.Article(text))

	//  import (
	// 	"fmt"
	// // 	"net/http"
	// // 	"time"
	//  )

	// func greet(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello World! %s", time.Now())
	// }

	// func main() {
	// 	http.HandleFunc("/", greet)
	// 	http.ListenAndServe(":8080", nil)
	// }

	// myCard := interfaces.CreditCard{CardNumber: "1234-5678-9012-3456"}
	// myPaypal := interfaces.Paypal{Email: "alabisunday@gmail.com"}

	// interfaces.ProcessPayment(myCard, 200.00)
	// interfaces.ProcessPayment(myPaypal, 50.0)

	// dog := animashelter.Dog{Name: "german"}
	// cat := animashelter.Cat{Breed: "persia"}

	// // animashelter.MakeThemTalk(dog)
	// // animashelter.MakeThemTalk(cat)

	// animals := []animashelter.Speaker{dog, cat}
	// for _, animal := range animals{
	// 	animashelter.MakeThemTalk(animal)
	// }
	mux := http.NewServeMux()
	mux.HandleFunc("/", asciiartweb.ServeHome)
	mux.HandleFunc("/ascii-art", asciiartweb.PrintAscii)
	fmt.Println("Server starting at :5000...")
	http.ListenAndServe(":5000", mux)
}
