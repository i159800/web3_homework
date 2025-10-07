create table if not exists students (
  id int unsigned auto_increment not null comment 'ID' ,
  name varchar(32) not null comment '姓名',
  age int unsigned not null comment '年龄',
  grade varchar(32) not null comment '年级',
  primary key(id)
)engine = InnoDb default charset=utf8 comment='学生表';

-- 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
insert students (name, age, grade) values ("张三", 20, "三年级");

-- 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
select * from students where age > 18;

-- 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。

update students set grade = '四年级' where name = '张三';

-- 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

delete from students where age < 15;

create table if not exists accounts (
  id int unsigned auto_increment not null comment 'ID',
  balance double not null comment '余额',
  primary key(id)
)engine = InnoDb default charset=utf8 comment='账户表';

create table if not exists transactions (
  id int unsigned auto_increment not null comment 'ID',
  from_account_id int unsigned not null comment '转出账户ID',
  to_account_id int unsigned not null comment '转入账户ID',
  amount double not null comment '转账金额',
  primary key(id)
)engine = InnoDb default charset=utf8 comment='交易表';

insert into accounts (id, balance) values (1001, 10000), (1002, 20000);

-- 定义存储过程

delimiter $$
drop procedure if exists multiTrans;
create procedure multiTrans(in fromId int, in toId int, in bal int)
begin

set @curBal = 0;
set @@autocommit=0;
-- select @@autocommit;
start transaction;
select balance into @curBal from accounts where id = fromId;
update accounts set balance = balance - bal where id = fromId;
update accounts set balance = balance + bal where id = toId;
insert into transactions (from_account_id, to_account_id, amount) values (fromId, toId, bal);
if @curBal >= bal then
  select 'commit' as message;
  commit;
else
  select 'rollback' as message;
  rollback;
end if;

end $$

-- 调用存储过程

call multiTrans(1001, 1002, 10000);

