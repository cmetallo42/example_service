-- wb_db

-- create table orders (
-- 	order_uid varchar,
-- 	track_number varchar,
-- 	entry varchar,
-- 	locale varchar,
-- 	internal_signature varchar,
-- 	customer_id varchar,
-- 	delivery_service varchar,
-- 	shardkey varchar,
-- 	sm_id integer,
-- 	date_created timestamp with time zone,
-- 	oof_shard varchar,

-- 	primary key (order_uid)
-- )

-- create table deliveries (
-- 	order_uid varchar,
-- 	name varchar,
-- 	phone varchar,
-- 	zip varchar,
-- 	city varchar,
-- 	address varchar,
-- 	region varchar,
-- 	email varchar,

-- 	foreign key (order_uid) references orders (order_uid)
-- )

-- create table payments (
-- 	order_uid varchar,
-- 	transaction varchar,
-- 	request_id varchar,
-- 	currency varchar,
-- 	provider varchar,
-- 	amount integer,
-- 	payment_dt integer,
-- 	bank varchar,
-- 	delivery_cost integer,
-- 	goods_total integer,
-- 	custom_fee integer,

-- 	foreign key (order_uid) references orders (order_uid)
-- )

-- create table items (
-- 	order_uid varchar,
-- 	chrt_id integer,
-- 	track_number varchar,
-- 	price integer,
-- 	rid varchar,
-- 	name varchar,
-- 	sale integer,
-- 	size varchar,
-- 	total_price integer,
-- 	nm_id integer,
-- 	brand varchar,
-- 	status integer,

-- 	foreign key (order_uid) references orders (order_uid)
-- )

create table test (
	id uuid not null default gen_random_uuid(),
	data jsonb not null default '{}'
);