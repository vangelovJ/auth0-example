package user

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"os"
	"fmt"
)

// Handler for our logged-in user page.
func Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	ctx.HTML(http.StatusOK, "user.html", profile)
	ctx.Redirect(http.StatusTemporaryRedirect, "http://192.168.163.132:8080/test1/")
}


func CheckRole(user string){
	url := "https://"+os.Getenv("AUTH0_DOMAIN")+"/api/v2/users/"+user+"/roles"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Bearer "+os.Getenv("MGMT_API_ACCESS_TOKEN")+"")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}