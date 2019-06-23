package controllers

import (
	"../models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Student_controllers(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("收到student请求")

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
	var sc_list couse_list
	var stu_info models.Studentinfo
	stu_info.Num=rq.Stu_num
	if rq.Do=="get_course"{
		rp.Object=models.Stu_score(stu_info)
	} else if rq.Do=="get_info"{
		rp.Object=models.Stu_info(stu_info.Num)
	}else if rq.Do=="get_course_list"{
		sc_list.Ac_list=models.Get_s_course(rq.Stu_num)
		sc_list.Tc_list=models.Get_s_without_course(rq.Stu_num)
		rp.Object=sc_list
	}else if rq.Do=="update_course"{
		models.Stu_course_list_update(rq.Stu_num,rq.C_list)
	}else if rq.Do=="update_password"{
		models.Stu_update_password(rq.Stu_num,rq.Stu_password)
	}
	fmt.Println(rp)
	res, _ := json.Marshal(rp)
	w.Write(res)

}

