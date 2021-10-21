package user

import (
	"github.com/gin-contrib/sessions"
	"net/http"

	"fmt"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

type User struct{
	Name string `json:"name"`
}
// Handler for our logged-in user page.
func Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	//ctx.HTML(http.StatusOK, "user.html", profile)
	v := CheckRole(profile)
	fmt.Println(v)
	ctx.Redirect(http.StatusTemporaryRedirect, "http://192.168.163.132:8080/test1/")
}


func CheckRole(profile interface{}) string{
	//fmt.Println(profile)

	profileData,err := json.Marshal(profile)
	if err != nil{
		fmt.Errorf(err.Error())
	}
	var u User
	err = json.Unmarshal(profileData,&u)
	if err != nil{
		fmt.Errorf(err.Error())
	}
	fmt.Println(u.Name)
	return u.Name
	//url := "https://"+os.Getenv("AUTH0_DOMAIN")+"/api/v2/users/"+profile["name"]+"/roles"

	//req, _ := http.NewRequest("GET", url, nil)

	//req.Header.Add("authorization", "Bearer "+os.Getenv("MGMT_API_ACCESS_TOKEN")+"")

	//res, _ := http.DefaultClient.Do(req)

	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))
}