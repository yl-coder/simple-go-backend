create table account
(
    id            int primary key auto_increment,
    customer_id int,
    type          text,
    balance          decimal ,
    created       timestamp,
    status numeric
);
alter table account add constraint account_customer_id_fk foreign key (customer_id) references customer (id);