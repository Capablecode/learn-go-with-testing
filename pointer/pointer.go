package pointer


func Swap(a *int, b *int){
	temp := *a
	*a = *b
	*b = temp
}


func Double(n *int){
	temp := *n
	temp *= 2
	*n = temp 
}