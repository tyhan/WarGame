package problem

import (
	"net/http"
	"io/ioutil"
)
		
func connect(file string, param map[string]string, session string) string {
	url := "http://los.sandbox.cash/chall/" + file
	cookie := http.Cookie{Name:"PHPSESSID",Value: session}
	
	if param != nil{
		url = url + "?"
		first := true
		for key, value := range param {
			if first != true {
				url += "&"
			}
			url = url + key +"=" + value
			first = false
		}	
	}
	
	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(&cookie)

	client := &http.Client{}	
	resp, _ := client.Do(req)
	
	defer resp.Body.Close()	
	contents, _ := ioutil.ReadAll(resp.Body)
	return string(contents)
}
