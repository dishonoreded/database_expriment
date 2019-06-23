package models

import (
	"fmt"
)

type Teacherinfo struct {
	T_num string `json:"t_num"`
	T_name	string `json:"t_name"`
	T_salary int `json:"t_salary"`
	T_title string `json:"t_title"`
	T_password string `json:"t_password"`
}
type Course struct {
	C_num string `json:"c_num"`
	C_name string `json:"c_name"`
	C_credit string `json:"c_credit"`
	Direction string `json:"direction"`
}

func Teacher_login(count string,password string) bool {
	var db_password string
	var permmit bool
	stmt:="select t_password from teacher where t_num=$1"
	row,err:=db.Query(stmt,count)
	checkError(err)
	if row.Next(){
		row.Scan(&db_password)
		if db_password==password{
			permmit=true
		}else{
			permmit=false
		}
	}
	return permmit
}



func Get_t_info() []Teacherinfo {
	var t_info_list []Teacherinfo
	var t_info Teacherinfo
	rows,err:=db.Query("select * from teacher order by cast(t_num as int);")
	defer rows.Close()
	checkError(err)
	for rows.Next(){
		rows.Scan(&t_info.T_num,&t_info.T_name,&t_info.T_title,&t_info.T_salary,&t_info.T_password)
		t_info_list=append(t_info_list, t_info)
	}
	fmt.Println("教师：",t_info_list)
	return t_info_list
}

func Get_all_course() []Course{
	var course_list []Course
	var course Course
	rows,err:=db.Query("select c_num,c_name from course where c_num not in(select c_num from tc);")
	defer rows.Close()
	checkError(err)
	for rows.Next(){
		rows.Scan(&course.C_num,&course.C_name)
		course_list=append(course_list,course)
	}
	fmt.Println("课程:",course_list)
	return course_list
 }

func Get_t_course(t_num string)[]Course{
	var course_list []Course
	var course Course
	rows,err:=db.Query("select tc.c_num,c_name from tc,course where tc.t_num=$1 and tc.c_num=course.c_num;",t_num)
	defer rows.Close()
	checkError(err)
	for rows.Next(){
		rows.Scan(&course.C_num,&course.C_name)
		course_list=append(course_list,course)
	}
	fmt.Println("t课程:",course_list)
	return course_list
}

func T_update(t_info Teacherinfo,c_list []Course){
	stmt,err:=db.Prepare("update teacher set t_name=$2,t_password=$3,t_salary=$4,t_title=$5 where t_num=$1")
	checkError(err)
	_,err=stmt.Exec(t_info.T_num,t_info.T_name,t_info.T_password,t_info.T_salary,t_info.T_title)
	checkError(err)
	stmt,err=db.Prepare("delete from tc where t_num=$1")
	checkError(err)
	stmt.Exec(t_info.T_num)

	stmt,err= db.Prepare("insert into tc(t_num, c_num) values ($1,$2);")
	for _,values :=range c_list{
		if values.Direction=="right"{
			stmt.Exec(t_info.T_num,values.C_num)
		}
	}
}

func T_create(t_info Teacherinfo,c_list []Course){
	stmt,err:=db.Prepare("insert into teacher values($1,$2,$3,$4,$5);")
	checkError(err)
	_,err=stmt.Exec(t_info.T_num,t_info.T_name,t_info.T_title,t_info.T_salary,t_info.T_password)
	checkError(err)
	stmt,err=db.Prepare("delete from tc where t_num=$1")
	checkError(err)
	stmt.Exec(t_info.T_num)

	stmt,err= db.Prepare("insert into tc(t_num, c_num) values ($1,$2);")
	for _,values :=range c_list{
		if values.Direction=="right"{
			stmt.Exec(t_info.T_num,values.C_num)
		}
	}
}

func T_delete(t_num string){
	fmt.Println("delete",t_num)
	stmt,err:=db.Prepare("delete from teacher where t_num=$1;")
	defer stmt.Close()
	checkError(err)
	stmt.Exec(t_num)
}
func T_retieve(keyword string) []Teacherinfo{
	fmt.Println("retieve")
	var t_info_list []Teacherinfo
	var t_info Teacherinfo
	rows,err:=db.Query("select * from teacher where t_name like $1;","%"+keyword+"%")
	checkError(err)
	defer rows.Close()
	for rows.Next(){
		rows.Scan(&t_info.T_num,&t_info.T_name,&t_info.T_title,&t_info.T_salary,&t_info.T_password)
		t_info_list=append(t_info_list,t_info)
	}
	fmt.Println("retieve list:",t_info_list)
	return t_info_list
}

