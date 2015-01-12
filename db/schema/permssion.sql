-- Table: permission

-- DROP TABLE permission;

CREATE TABLE permission
(
  id bigint NOT NULL DEFAULT id_generator(),
  name text,
  category boolean,
  CONSTRAINT permission_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

-- Index: index_permission_on_id_type

-- DROP INDEX index_permission_on_id_type;

CREATE INDEX index_permission_on_id_type
  ON permission
  USING btree
  (id);

