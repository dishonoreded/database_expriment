/*==============================================================*/
/* DBMS name:      PostgreSQL 9.x                               */
/* Created on:     2019/6/11 23:25:20                           */
/*==============================================================*/


drop view v_sc;

drop table course;

drop table manager;

drop table sc;

drop table student;

drop table tc;

drop table teacher;

/*==============================================================*/
/* Table: course                                                */
/*==============================================================*/
create table course (
   c_num                VARCHAR(10)          not null,
   c_name               VARCHAR(10)          null,
   c_credit             INT4                 null,
   constraint PK_COURSE primary key (c_num)
);

comment on column course.c_num is
'课程号';

comment on column course.c_name is
'课程名';

comment on column course.c_credit is
'学分';

/*==============================================================*/
/* Table: manager                                               */
/*==============================================================*/
create table manager (
   count                VARCHAR(10)          null,
   password             VARCHAR(50)          null
);

comment on column manager.count is
'账号';

comment on column manager.password is
'密码';

/*==============================================================*/
/* Table: sc                                                    */
/*==============================================================*/
create table sc (
   s_num                VARCHAR(10)          not null,
   c_num                VARCHAR(10)          not null,
   score                INT4                 null,
   constraint PK_SC primary key (s_num, c_num)
);

comment on column sc.s_num is
'学号';

comment on column sc.c_num is
'课程号';

comment on column sc.score is
'成绩';

/*==============================================================*/
/* Table: student                                               */
/*==============================================================*/
create table student (
   s_num                VARCHAR(10)          not null,
   s_sex                VARCHAR(20)          null,
   s_age                INT4                 null,
   s_name               VARCHAR(50)          null,
   s_password           VARCHAR(50)          null,
   constraint PK_STUDENT primary key (s_num)
);

comment on column student.s_num is
'学号';

comment on column student.s_sex is
'性别';

comment on column student.s_age is
'年龄';

comment on column student.s_name is
'姓名';

comment on column student.s_password is
'密码';

/*==============================================================*/
/* Table: tc                                                    */
/*==============================================================*/
create table tc (
   t_num                VARCHAR(10)          not null,
   c_num                VARCHAR(10)          not null,
   constraint PK_TC primary key (t_num, c_num)
);

comment on column tc.t_num is
'工号';

comment on column tc.c_num is
'课程号';

/*==============================================================*/
/* Table: teacher                                               */
/*==============================================================*/
create table teacher (
   t_num                VARCHAR(10)          not null,
   t_name               VARCHAR(50)          null,
   t_title              VARCHAR(50)          null,
   t_salary             INT8                 null,
   t_password           VARCHAR(50)          null,
   constraint PK_TEACHER primary key (t_num)
);

comment on column teacher.t_num is
'工号';

comment on column teacher.t_name is
'姓名';

comment on column teacher.t_title is
'职称';

comment on column teacher.t_salary is
'工资';

comment on column teacher.t_password is
'密码';

/*==============================================================*/
/* View: v_sc                                                   */
/*==============================================================*/
create or replace view v_sc as
select s_num,sc.c_num,score,c_name,c_credit from sc,course where sc.c_num=course.c_num;

alter table sc
   add constraint FK_SC_REFERENCE_COURSE foreign key (c_num)
      references course (c_num)
      on delete cascade on update cascade;

alter table sc
   add constraint FK_SC_REFERENCE_STUDENT foreign key (s_num)
      references student (s_num)
      on delete cascade on update cascade;

alter table tc
   add constraint FK_TC_REFERENCE_COURSE foreign key (c_num)
      references course (c_num)
      on delete cascade on update cascade;

alter table tc
   add constraint FK_TC_REFERENCE_TEACHER foreign key (t_num)
      references teacher (t_num)
      on delete cascade on update cascade;

