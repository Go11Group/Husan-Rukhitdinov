create table transactions(
    id uuid primary key  default gen_random_uuid(),
    card_id uuid not null  references cards(id),
    amount bigint,
    terminal_id uuid not null  references terminals(id) default null,
    transaction_type TransactionType,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at bigint default 0


);