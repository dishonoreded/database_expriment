package controllers

import (
	"../models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type couse_list struct {
	Tc_list []models.Course `json:"tc_list"`
	Ac_list []models.Course `json:"ac_list"`
}
type count struct {
	T_count int `json:"t_count"`
	S_count float64 `json:"s_count"`
	Max int `json:"max"`
	Min int `json:"min"`
	Avg float64 `json:"avg"`
}

func Manager_controllers(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("收到manager请求")

	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	defer r.Body.Close()
	con, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("请求error:", err)
	}
	var rq req
	err = json.Unmarshal([]byte(string(con)), &rq)
	if err != nil {
		fmt.Println("json解析error:", err)
	}
	fmt.Println("收到的请求",rq)
	var rp res
	var stu_info models.Studentinfo
	var t_info	models.Teacherinfo
	var c count
	stu_info.Num=rq.Stu_num
 	if rq.Do=="init"{
		rp.Object=models.Get_stu_info()
	}else if rq.Do=="create"{
		stu_info.Password=rq.Stu_password
		stu_info.Name=rq.Stu_name
		stu_info.Age=rq.Stu_age
		stu_info.Sex=rq.Stu_sex
		if(models.Stu_create(stu_info)){
			rp.Object=true
		}else{
			rp.Object=false
		}
	}else if rq.Do=="update"{
		stu_info.Password=rq.Stu_password
		stu_info.Name=rq.Stu_name
		stu_info.Age=rq.Stu_age
		stu_info.Sex=rq.Stu_sex
		models.Stu_update(stu_info)
		rp.Object=true
	}else if rq.Do=="retieve"{
		rp.Object=models.Stu_retieve(rq.Keyword)
	}else if rq.Do=="delete"{
		models.Stu_delete(stu_info)
	}else if rq.Do=="get_score"{
		rp.Object=models.Stu_score(stu_info)
	}else if rq.Do=="score_update"{
		fmt.Println("stu:",rq.Stu_num,"cnum:",rq.Cou_num,"csc:",rq.Cou_score)
		models.Stu_score_update(rq.Stu_num,rq.Cou_num,rq.Cou_score)
	}else if rq.Do=="get_teacher"{
		rp.Object=models.Get_t_info()
	}else if rq.Do=="get_course"{
		var t_a couse_list
		t_a.Ac_list=models.Get_all_course()
		t_a.Tc_list=models.Get_t_course(rq.T_num)
		rp.Object=t_a
	}else if rq.Do=="t_update"{
		t_info.T_num=rq.T_num
		t_info.T_name=rq.T_name
		t_info.T_password=rq.T_password
		t_info.T_salary=rq.T_salary
		t_info.T_title=rq.T_title
		models.T_update(t_info,rq.C_list)
	}else if rq.Do=="t_create"{
		t_info.T_num=rq.T_num
		t_info.T_name=rq.T_name
		t_info.T_password=rq.T_password
		t_info.T_salary=rq.T_salary
		t_info.T_title=rq.T_title
		models.T_create(t_info,rq.C_list)
	}else if rq.Do=="t_delete"{
		models.T_delete(rq.T_num)
	}else if rq.Do=="t_retieve"{
		rp.Object= models.T_retieve(rq.Keyword)
	}else if rq.Do=="count"{
		c.T_count,c.S_count=models.Count_teacher(rq.T_title)
		c.Max,c.Min,c.Avg=models.Count_course(rq.Cou_name)
		fmt.Println("abg:",c.Avg)
		rp.Object=c
	}

	fmt.Println(rp)
	res, _ := json.Marshal(rp)
	w.Write(res)


}
