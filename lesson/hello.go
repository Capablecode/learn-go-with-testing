package lesson

//	func Hello(name string) string{
//		return "Hello, " + name
//	}
// const englishPrefixHello = "Hello, "
// const spanishPrefixHello = "Holla, "
// const frenchPrefixHello = "Bonjour, "
// const spanish = "Spanish"

// func Hello(name string, language string) string {
// 	if name == "" {
// 		name = "World"
// 	}
// 	if language == "Spanish" {
// 		return spanishPrefixHello + name
// 	}
// 	if language == "French" {
// 		return frenchPrefixHello + name
// 	}
// 	return englishPrefixHello + name
// }

// func Hello(name string, language string) string {
// 	prefix := englishPrefixHello

// 	if name == ""{
// 		name = "World"
// 	}

// 	switch language{
// 	case "Spanish":
// 		prefix = spanishPrefixHello
// 	case "French":
// 		prefix = frenchPrefixHello
// 	}
// 	return prefix + name
// }

//Declaration 
const (
	spanish = "Spanish"
	french  = "French"

	englishPrefixHello = "Hello, "
	spanishPrefixHello = "Holla, "
	frenchPrefixHello  = "Bonjour, "
)

//Main Function definition
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

//Helper function
func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishPrefixHello
	case "French":
		prefix = frenchPrefixHello
	default:
		prefix = englishPrefixHello
	}
	return prefix
}
