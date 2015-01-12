-- Table: acl

-- DROP TABLE acl;

CREATE TABLE acl
(
  id bigint NOT NULL DEFAULT id_generator(),
  entity_id bigint,
  entity_name text,
  role_id bigint,
  CONSTRAINT acl_pkey PRIMARY KEY (id),
  CONSTRAINT acl_role_fk FOREIGN KEY (role_id)
      REFERENCES role (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
  OIDS=FALSE
);

-- Index: fki_acl_role_fk

-- DROP INDEX fki_acl_role_fk;

CREATE INDEX fki_acl_role_fk
  ON acl
  USING btree
  (role_id);

