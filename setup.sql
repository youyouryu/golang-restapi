create database if not exists restapi;
use restapi;

drop table if exists users;

create table users (
  user_id varchar(255) not null unique,
  password varchar(255),
  nickname varchar(255),
  comment varchar(255)
);

insert into users values ("TaroYamada", "af205051b1fd61b9a7bdda37e336eebbf8730df3b5dc0e3a16dc830c4706f37a1350621d815a264ea9456f67914d6c0c2c713dbc1d0de9a920e402c6abaa8d3d", "たろー", "僕は元気です");
select * from users;
