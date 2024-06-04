create table products
(
    id       uuid primary key,
    name     varchar,
    category varchar,
    cost     int
);

-- single index
create index product_id_idx on product (id);

-- multi index
create index product_id_idx_cat on product (name,category);

drop index product_id_idx_cat;

create unique index product_id_idx on product (id, cost);

explain (analyse )
select * from product where id  = '20e930be-31a0-4208-8d9f-13ae79540f0a' and cost=85955 -- id = 'b2457480-75c7-481b-8b76-c40795ec8ff0';

