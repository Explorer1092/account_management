drop table if exists `person_info`;
create table `person_info`(
  `id` int(11) not null ,
  `name` varchar(20) ,
  `img_src` varchar(100),
  `check_in` int default 0,
  `check_time` varchar(20) default "2006-01-02 15:04:05",
   primary key (`id`)
)engine=innodb  charset=utf8mb4;



drop table if exists `verification_info`;
create table `verification_info`(
  `id` int(11) not null auto_increment,
  `phone_num` varchar(20) not null,
  `password` varchar(20) not null ,
  primary key (`id`)
)engine=innodb  charset=utf8mb4;


drop table if exists `bill_info`;
create table `bill_info`(
    `id` int(11)not null ,
    `time` varchar (20),
    `type` varchar (20),
    `otherside` varchar(100),
    `goodsname` varchar(100),
    `way` varchar(20),
    `price` varchar(20),
    `date` varchar(20),
    `timestamp` int(11)  ,
    `classification` varchar(20) default '其他'
)engine=innodb  charset=utf8mb4;

