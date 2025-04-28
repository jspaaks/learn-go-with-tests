package helloworld

const (
	dutch   = "Dutch"
	french  = "French"
	spanish = "Spanish"

	dutchHelloPrefix   = "Hallo"
	englishHelloPrefix = "Hello"
	frenchHelloPrefix  = "Bonjour"
	spanishHelloPrefix = "Hola"
)

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case dutch:
		prefix = dutchHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func SayHello(name string, language string) (message string) {
	if name == "" {
		name = "world"
	}
	var prefix string = getGreetingPrefix(language)
	message = prefix + ", " + name
	return
}
