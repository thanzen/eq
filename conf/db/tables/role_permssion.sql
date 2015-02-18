-- Table: role_permission

-- DROP TABLE role_permission;

CREATE TABLE role_permission
(
  role_id bigint NOT NULL,
  permission_id bigint NOT NULL,
  CONSTRAINT role_permission_pkey PRIMARY KEY (role_id, permission_id),
  CONSTRAINT role_permission_permission_fk FOREIGN KEY (permission_id)
      REFERENCES permission (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT role_permission_role_fk FOREIGN KEY (role_id)
      REFERENCES role (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
  OIDS=FALSE
);
-- Index: fki_role_permission_permission

-- DROP INDEX fki_role_permission_permission;

CREATE INDEX fki_role_permission_permission
  ON role_permission
  USING btree
  (permission_id);

