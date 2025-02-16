create table if not exists users (
    user_id bigserial primary key,
    username varchar(50) not null,
    hash_password varchar(255) not null,
    balance bigint not null
    );

create table if not exists transactions (
    transaction_id bigserial primary key,
    from_user_id bigint not null,
    to_user_id bigint not null,
    money bigint not null,
    created_at timestamp not null,
    foreign key (from_user_id) references users(user_id),
    foreign key (to_user_id) references users(user_id)
    );

create table if not exists merch (
    merch_id bigserial primary key,
    merch_name varchar(20) not null,
    price int not null
    );

create table if not exists user_merch (
    user_id bigint not null,
    merch_id bigint not null,
    quantity bigint not null,
    foreign key (user_id) references users(user_id),
    foreign key (merch_id) references merch(merch_id),
    primary key (user_id, merch_id)
    );

insert into merch (merch_name, price)
values
    ('t-shirt', 80),
    ('cup', 20),
    ('book', 50),
    ('pen', 10),
    ('powerbank', 200),
    ('hoody', 300),
    ('umbrella', 200),
    ('socks', 10),
    ('wallet', 50),
    ('pink-hoody', 500);