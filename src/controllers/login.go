package controllers

import (
	"../models"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"time"
)

const SecretKey = "database_key"

//payload载荷
type jwtCustomClaims struct {
	jwt.StandardClaims
	// 追加自己需要的信息
	User_name string `json:"user_name"`
	User_type string `json:"User_type"`
	Admin     bool   `json:"admin"`
}

type req struct {
	Count        string          `json:"count"`
	Password     string          `json:"password"`
	User_type    string          `json:"user_type"`
	Keyword      string          `json:"keyword"`
	Stu_num      string          `json:"stu_num"`
	Stu_sex      string          `json:"stu_sex"`
	Stu_age      int             `json:"stu_age"`
	Stu_name     string          `json:"stu_name"`
	Stu_password string          `json:"stu_password"`
	Cou_num      string          `json:"cou_num"`
	Cou_score    int             `json:"cou_score"`
	Cou_name	string		`json:"cou_name"`
	T_num        string          `json:"t_num"`
	T_name       string          `json:"t_name"`
	T_salary     int             `json:"t_salary"`
	T_title      string          `json:"t_title"`
	T_password   string          `json:"t_password"`
	C_list       []models.Course `json:"c_list"`
	Do           string          `json:"do"`
}

type res struct {
	Permission bool   `json:"permission"`
	Count      string `json:"count"`
	Object     interface{}
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json") //返回数据格式是json
	defer r.Body.Close()
	con, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("请求error:", err)
	}
	var request req
	var response res
	err = json.Unmarshal([]byte(string(con)), &request)
	if err != nil {
		fmt.Println("json解析error:", err)
	}
	fmt.Println("收到请求：", request)
	if request.User_type == "manager" {
		response.Permission = models.Manager_login(request.Count, request.Password)
	} else if request.User_type == "teacher" {
		response.Permission = models.Teacher_login(request.Count, request.Password)
	} else if request.User_type == "student" {
		response.Permission = models.Student_login(request.Count, request.Password)
	} else {
		fmt.Println("User_type miss")
	}

	if response.Permission {
		w.Header().Set("Set-Cookie", NewCookie([]byte(SecretKey), request.Count, request.User_type, true))
	} else {
		fmt.Println("登录失败")
	}
	fmt.Println("返回的数据", response)
	byte_res, err := json.Marshal(response)
	w.Write(byte_res)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//生成无效的token返回
	w.Header().Set("Access-Control-Allow-Origin", "*")                           //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Authorization") //header的类型
	w.Header().Set("content-type", "application/json")                           //返回数据格式是json
	w.Header().Set("Set-Cookie", NewCookie([]byte(SecretKey),"","",false))
	w.Write([]byte(""))
}

func Check(w http.ResponseWriter, r *http.Request) {
	var response res
	r.ParseForm()
	for _, cookie := range r.Cookies() {
		token := cookie.Value
		if token != "" {
			response.Permission, response.Count = Check_permission(token, r.Form.Get("user_type"))
		} else {
			response.Permission = false
		}
	}
	res, _ := json.Marshal(response)
	w.Write(res)
}

func CreateToken(SecretKey []byte, issuer string, User_type string, isAdmin bool) (tokenString string, err error) { //创建token
	User_name := issuer
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    issuer,
		},
		User_name,
		User_type,
		isAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)

	return
}
func ParseToken(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) { //解析token
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims
	return
}

func Check_permission(req_token string, user_type string) (bool, string) {
	claims, err := ParseToken(req_token, []byte(SecretKey))
	if err != nil {
		fmt.Println("check_err:", err)
		return false, ""
	}
	if claims.(jwt.MapClaims)["User_type"] != user_type || claims.(jwt.MapClaims)["admin"] != true {
		return false, ""
	}
	fmt.Printf("claims uid:%v,claims uid type:%T\n", claims.(jwt.MapClaims)["user_name"], claims.(jwt.MapClaims)["user_name"])
	str_name := claims.(jwt.MapClaims)["user_name"].(string)
	return true, str_name
}

func NewCookie(SecretKey []byte, issuer string, User_type string, isAdmin bool) (newCookie string) {
	var cookie http.Cookie
	if isAdmin{
		newToken, _ := CreateToken([]byte(SecretKey), issuer, User_type, true)
		cookie= http.Cookie{
			Name:     "token",
			Value:    newToken,
			HttpOnly: true,
		}
	}else{
		cookie= http.Cookie{
			Name:     "token",
			Value:    "",
			HttpOnly: true,
		}
	}


	return cookie.String()
}
