package interfaces

import "fmt"

type Payer interface{
	Payer (amount float32)string
}

//creating a concrete type that can implement the interface
type CreditCard struct{
	CardNumber string 
}

type Paypal struct{
	Email string
}

//Creating the method 
func (cc CreditCard) Payer(amount float32)string{
	return fmt.Sprintf("Paid $%.2f using credit card: %s", amount, cc.CardNumber)
}

//creating the method
func (pp Paypal) Payer(amount float32)string{
	return fmt.Sprintf("Paid $%f using Paypal account: %s", amount, pp.Email)
}

//using the interface
func ProcessPayment(p Payer, amount float32){
	fmt.Println(p.Payer(amount))
}

func main(){
	// myCard := CreditCard{CardNumber: "1234-5678-9012-3456"}
	// myPaypal := Paypal{Email: "alabisunday@gmail.com"}

	// ProcessPayment(myCard, 200.00)
	// ProcessPayment(myPaypal, 50.0)
	
}