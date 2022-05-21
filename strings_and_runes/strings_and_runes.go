package main
import( 
	"fmt"
	"unicode/utf8"
)
func main() {
	const s = "สวัสดี"
	fmt.Println("Length of s: ", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	for index, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, index)
	}

	for i,w := 0,0; i < len(s); i+=w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r =='t' {
		fmt.Println("Found tee")
	} else if r == 'ส'{
		fmt.Println("Found so sua")
	}
}

