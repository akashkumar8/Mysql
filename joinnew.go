CREATE TABLE `custom`.`customers` (
  `cid` INT NOT NULL,
  `cname` VARCHAR(45) NOT NULL,
  `cemail` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`cid`, `cname`, `cemail`));

CREATE TABLE 'custom'.'orders' (
  'oid' INT NOT NULL,
  'orderdate' DATE NOT NUll,
  'oamount' INT NOT NULL,
  'cid' INT NOT NULL,
  PRIMARY KEY('oid', 'orderdate', 'oamount', 'oid'));
  
SELECT column_name FROM table1 LEFT JOIN table2
   ON table1.column_name = table2.column_name;

SELECT customers.cid, cname, oamount
FROM customers
LEFT JOIN orders
ON customers.cid = orders.cid;