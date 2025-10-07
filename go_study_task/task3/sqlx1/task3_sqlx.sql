create table if not exists employees (
  id int unsigned auto_increment not null comment 'ID' ,
  name varchar(32) not null comment '姓名',
  department varchar(32) not null comment '部门',
  salary double not null comment '薪水',
  primary key(id)
  )engine = InnoDb default charset=utf8 comment='员工表';

create table if not exists books (
   id int unsigned auto_increment not null comment 'ID' ,
   title varchar(32) not null comment '标题',
  author varchar(32) not null comment '作者',
  price double not null comment '价格',
  primary key(id)
  )engine = InnoDb default charset=utf8 comment='书籍表';
