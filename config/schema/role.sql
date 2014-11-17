CREATE TABLE role(
	id bigint NOT NULL default id_generator(),
	name text,
	active boolean,
	is_system_role boolean,
	system_name text,
	CONSTRAINT role_pkey PRIMARY KEY (id)
);
CREATE INDEX index_role_on_id_type ON role USING btree (id);