package iteration

const nRepeat int = 5

func Repeat(s string) string {	
	var repeated string = ""
	for i := 0; i < nRepeat; i++ {
		repeated += s
	}
	return repeated
}