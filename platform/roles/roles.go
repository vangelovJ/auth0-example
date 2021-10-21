package roles

import (
"fmt"
"net/http"
"io/ioutil"
"os"
)


func checkRole(user string){
	url := "https://"+os.Getenv("AUTH0_DOMAIN")+"/api/v2/users/"+user+"/roles"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Bearer "+os.Getenv("MGMT_API_ACCESS_TOKEN")+"")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}