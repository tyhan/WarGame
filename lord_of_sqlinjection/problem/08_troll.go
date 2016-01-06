package problem

import(
	"fmt"
	"strings"
)

func Troll(session string){
	file := "troll_4db5bdeebb8a99cbe4534cf1fcfc37eb.php"
	param := map[string]string{
		"id" : "AdMin",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"TROLL Clear!"){
		fmt.Println("TROLL Clear!")
	}
}
