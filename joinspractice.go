CREATE TABLE `custom`.`customers` (
  `cid` INT NOT NULL,
  `cname` VARCHAR(45) NOT NULL,
  `cemail` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`cid`, `cname`, `cemail`));



  /*create database joins;
use joins;
create table table1(column1 int);
create table table2(column2 int);
show tables 

select*from table1 INNER JOIN table2 on table1.column1=table2.column2;/*use of inner join 
use ccustomers;
select* from customers;
select * from orders;
select cid,cname,cemail from customers INNER JOIN projects ON customers.cid = orders.cid;*/

