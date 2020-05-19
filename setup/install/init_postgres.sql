/*表结构*/
CREATE TABLE public.user
(
  id serial primary key,
  name varchar(20)
);

ALTER TABLE public.user ADD COLUMN created timestamp default now();
ALTER TABLE public.user ADD COLUMN class_id integer default 1;


insert into public.user(name) values('ft');
insert into public.user(name) values('ft2');
insert into public.user(name,class_id) values('ft8',2);


CREATE TABLE public.class
(
  id serial primary key,
  name varchar(20)
);
insert into public.class(name) values('高中一班');
insert into public.class(name) values('高中二班');

/*常用到的操作*/
select * from public.class;
select * from public.user;
select u.id,u.name,c.name from public.user as u left join public.class as c on u.class_id=c.id;
