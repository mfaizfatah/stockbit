-- create table user
CREATE TABLE IF NOT EXISTS user(
  id int(11) NOT NULL,
  username varchar(20) NOT NULL,
  parent int(11) DEFAULT 0,
  PRIMARY KEY (id)
);

-- insert value to table user
INSERT INTO user (id,username,parent) VALUES
  (1, 'Ali',2),
  (2, 'Budi',0),
  (3, 'Cecep',1);
  
-- sql query for displaying ParentUserName
SELECT a.id as ID, a.username as UserName, b.username as ParentUserName FROM `user` a LEFT JOIN `user` b ON a.parent = b.id;