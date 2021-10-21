package user

import (
	"github.com/gin-contrib/sessions"
	"io/ioutil"
	"net/http"

	"fmt"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"os"
)

type User struct{
	Name string `json:"name"`
	Sub string `json:"sub"`
	Allowed bool
}


type Roles struct{
	AssignedRoles []Role
}

type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
// Handler for our logged-in user page.
func Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	//ctx.HTML(http.StatusOK, "user.html", profile)
	user := checkUser(profile)
	checkRole(&user,"Authorized")
	if user.Allowed{
		ctx.Redirect(http.StatusTemporaryRedirect, "http://192.168.163.132:8080/test1/")
	}else{
		ctx.HTML(http.StatusOK, "user.html", profile)
	}

}

func checkUser(profile interface{}) User{
	profileData,err := json.Marshal(profile)
	if err != nil{
		fmt.Errorf(err.Error())
	}
	var u User
	err = json.Unmarshal(profileData,&u)
	if err != nil{
		fmt.Errorf(err.Error())
	}
	return u
}

func checkRole(u *User, checkRole string){

	url := "https://"+os.Getenv("AUTH0_DOMAIN")+"/api/v2/users/"+u.Sub+"/roles"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Bearer "+os.Getenv("MGMT_API_ACCESS_TOKEN")+"")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var roles Roles
	err := json.Unmarshal(body,&roles)
	if err != nil{
		fmt.Errorf(err.Error())
	}
	for _,role := range(roles.AssignedRoles) {
		if role.Name == checkRole {
			u.Allowed = true
			break
		}else{
			u.Allowed = false
		}
	}

	//fmt.Println(res)
	//fmt.Println(u.Sub)
	//fmt.Println(string(body))
}