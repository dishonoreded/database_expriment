package models

import "fmt"

type Studentinfo struct {
	Num string `json:"num"`
	Sex string `json:"sex"`
	Age int 	`json:"age"`
	Name string `json:"name"`
	Password string `json:"password"`
}
type scores struct {
	C_num string `json:"c_num"`
	C_name string `json:"c_name"`
	C_credit int `json:"c_credit"`

	Score int `json:"score"`
}
func Student_login(count string,password string) bool {
	var db_password string
	var permmit bool
	stmt:="select s_password from student where s_num=$1"
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


func Get_stu_info() []Studentinfo {
	var stu_info_list []Studentinfo
	var stu_info Studentinfo
	rows,err:=db.Query("select * from student order by cast(s_num as int);")
	defer rows.Close()
	checkError(err)
	for rows.Next(){
		rows.Scan(&stu_info.Num,&stu_info.Sex,&stu_info.Age,&stu_info.Name,&stu_info.Password)
		stu_info_list=append(stu_info_list, stu_info)
	}
	return stu_info_list
}

func Stu_create(studentinfo Studentinfo) bool {
	fmt.Println("create")
	rows,err:=db.Query("select * from student where s_num=$1;",studentinfo.Num)
	defer rows.Close()
	checkError(err)
	for rows.Next(){
		return false
	}
	stmt,err:=db.Prepare("insert into student values ($1,$2,$3,$4,$5);")
	checkError(err)
	defer stmt.Close()
	a,err:=stmt.Exec(studentinfo.Num,studentinfo.Sex,studentinfo.Age,studentinfo.Name,studentinfo.Password)
	checkError(err)
	fmt.Printf("type:%T",a)
	return true
}

func Stu_update(studentinfo Studentinfo){
	fmt.Println("update")
	fmt.Println(studentinfo)
	stmt,err:=db.Prepare("update student set s_sex=$1,s_age=$2,s_name=$3,s_password=$4 where s_num=$5;")
	defer stmt.Close()
	checkError(err)
	stmt.Exec(studentinfo.Sex,studentinfo.Age,studentinfo.Name,studentinfo.Password,studentinfo.Num)
}

func Stu_retieve(keyword string) []Studentinfo{
	fmt.Println("retieve")
	var stu_info_list []Studentinfo
	var stu_info Studentinfo
	rows,err:=db.Query("select * from student where s_name like $1;","%"+keyword+"%")
	checkError(err)
	defer rows.Close()
	for rows.Next(){
		rows.Scan(&stu_info.Num,&stu_info.Sex,&stu_info.Age,&stu_info.Name,&stu_info.Password)
		stu_info_list=append(stu_info_list, stu_info)
	}
	fmt.Println("retieve list:",stu_info_list)
	return stu_info_list
}

func Stu_delete(studentinfo Studentinfo){
	fmt.Println("delete",studentinfo.Num)
	stmt,err:=db.Prepare("delete from student where s_num=$1;")
	defer stmt.Close()
	checkError(err)
	stmt.Exec(studentinfo.Num)
}

func Stu_score(studentinfo Studentinfo)[]scores{
	var score scores
	var score_list []scores
	fmt.Println(studentinfo.Num)
	rows,err:=db.Query("select c_num,score,c_name,c_credit from v_sc where s_num=$1",studentinfo.Num)
	checkError(err)
	defer rows.Close()
	for rows.Next(){
		rows.Scan(&score.C_num,&score.Score,&score.C_name,&score.C_credit)
		score_list=append(score_list,score)
	}
	fmt.Println(score_list)
	return score_list
}

func Stu_score_update(s_num string,c_num string,score int){
	stmt,err:=db.Prepare("update sc set score=$1 where s_num=$2 and c_num=$3;")
	defer stmt.Close()
	checkError(err)
	result,err:=stmt.Exec(score,s_num,c_num)
	fmt.Println("result:",result)
	checkError(err)
}

func Stu_info(s_num string) (stu_info Studentinfo){
	fmt.Println("retieve")
	rows,err:=db.Query("select * from student where s_num=$1;",s_num)
	checkError(err)
	defer rows.Close()
	for rows.Next(){
		rows.Scan(&stu_info.Num,&stu_info.Sex,&stu_info.Age,&stu_info.Name,&stu_info.Password)
	}
	fmt.Println("stu_info::",stu_info)
	return
}
func Get_s_without_course(s_num string) []Course{
	var course_list []Course
	var course Course
	rows,err:=db.Query("select c_num, c_name, c_credit from course where c_num not in(select c_num from sc where s_num=$1);",s_num)
	defer rows.Close()
	checkError(err)
	for rows.Next(){
		rows.Scan(&course.C_num,&course.C_name,&course.C_credit)
		course_list=append(course_list,course)
	}
	return course_list
}

func Get_s_course(s_num string)[]Course{
	var course_list []Course
	var course Course
	fmt.Println("s_num",s_num)
	rows,err:=db.Query("select sc.c_num,c_name,c_credit from sc,course where sc.s_num=$1 and sc.c_num=course.c_num;",s_num)
	defer rows.Close()
	checkError(err)
	for rows.Next(){
		rows.Scan(&course.C_num,&course.C_name,&course.C_credit)
		course_list=append(course_list,course)
	}
	fmt.Println(course_list)
	return course_list
}


func Stu_course_list_update(s_num string,course_list []Course){
	fmt.Println(course_list)
	stmt,err:=db.Prepare("delete from sc where s_num=$1")
	checkError(err)
	stmt.Exec(s_num)
	stmt,err= db.Prepare("insert into sc(s_num, c_num) values ($1,$2);")
	for _,values :=range course_list{
		if values.Direction=="right"{
			stmt.Exec(s_num,values.C_num)
		}
	}
}

func Stu_update_password(s_num string,s_password string){
	stmt,err:=db.Prepare("update student set s_password=$2 where s_num=$1;")
	defer stmt.Close()
	checkError(err)
	result,err:=stmt.Exec(s_num,s_password)
	fmt.Println("result:",result)
	checkError(err)
}