create table transaction
(
    id            int primary key auto_increment,
    account_id    int,
    type          text,
    amount          decimal ,
    created       timestamp
);
alter table transaction add constraint transaction_account_id_fk foreign key (account_id) references account (id);

