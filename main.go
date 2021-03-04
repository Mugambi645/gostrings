//simple word game
//words must start with 'be' 
//contain letter 'a' and
//end with 'ng'
package main
import (
	"fmt"
	"strings"
)
func main(){
	Loop:
	fmt.Print("Enter a word: ")
	var userWord string
	fmt.Scanln(&userWord)
	strings.ToLower(userWord)
	if (strings.Contains(userWord, "a") && strings.HasPrefix(userWord,"be") && strings.HasSuffix(userWord, "ng")){
		fmt.Println("Congrats!!...found a match!:", userWord)
	goto Loop
}else {
	fmt.Println("Failed..?!No match(Exiting programm):", userWord)
}
}