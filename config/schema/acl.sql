CREATE TABLE acl
(
	id bigint not null default id_generator(),
	entity_id bigint,
	entity_name text,
	role_id bigint not null REFERENCES role (id),
	constraint acl_pkey primary key (id)
);
