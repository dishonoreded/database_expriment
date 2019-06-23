package models

import "fmt"

func Manager_login(count string,password string) bool {
	var db_password string
	var permmit bool
	stmt:="select password from manager where count=$1"
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

func Count_teacher(title string)(t_count int,s_count float64){
	fmt.Println(title)
	rows,err:=db.Query("select count(*),avg(t_salary) from teacher where t_title=$1;",title)
	checkError(err)
	for rows.Next(){
		rows.Scan(&t_count,&s_count)
		fmt.Println(t_count,s_count)
	}
	fmt.Println(t_count,s_count)
	return
}

func Count_course(c_name string)(max int,min int,avg float64){
	rows,err:=db.Query("select max(score),min(score),avg(score) from sc,course where c_name=$1 and sc.c_num=course.c_num;",c_name)
	checkError(err)
	for rows.Next(){
		rows.Scan(&max,&min,&avg)
	}
	return
}

