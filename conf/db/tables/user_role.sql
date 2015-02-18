-- Table: user_role

-- DROP TABLE user_role;

CREATE TABLE user_role
(
  user_info_id bigint NOT NULL,
  role_id bigint NOT NULL,
  CONSTRAINT user_role_pkey PRIMARY KEY (user_info_id, role_id),
  CONSTRAINT user_role_role FOREIGN KEY (role_id)
      REFERENCES role (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT user_role_user_fk FOREIGN KEY (user_info_id)
     REFERENCES user_info (id) MATCH SIMPLE 
     ON DELETE CASCADE
)
WITH (
  OIDS=FALSE
);
-- Index: fki_user_role_role

-- DROP INDEX fki_user_role_role;

CREATE INDEX fki_user_role_role
  ON user_role
  USING btree
  (role_id);

