create table customer
(
    id            int primary key auto_increment,
    name          text null,
    city          text null,
    zipcode       text null,
    date_of_birth timestamp null,
    status        numeric null
);
