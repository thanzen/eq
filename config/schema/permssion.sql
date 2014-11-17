create table permission(
	id bigint not null default id_generator(),
	name text,
	category boolean,
	constraint permission_pkey primary key (id)
);
create index index_permission_on_id_type on permission using btree (id);