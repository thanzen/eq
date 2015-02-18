-- Table: role

-- DROP TABLE role;

CREATE TABLE role
(
  id bigint NOT NULL DEFAULT id_generator(),
  name text,
  active boolean,
  is_system_role boolean,
  description text,
  CONSTRAINT role_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
-- Index: index_role_on_id_type

-- DROP INDEX index_role_on_id_type;

CREATE INDEX index_role_on_id_type
  ON role
  USING btree
  (id);

