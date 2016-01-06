package problem

import(
	"fmt"
	"strings"
)

func Vampire(session string){
	file := "vampire_882335819ad8843fca1ee6f67976335f.php"
	param := map[string]string{
		"id" : "AdMin",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"VAMPIRE Clear!"){
		fmt.Println("VAMPIRE Clear!")
	}
}
