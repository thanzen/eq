-- Table: user_role

-- DROP TABLE user_role;

CREATE TABLE user_role
(
  user_id bigint  not null REFERENCES user_meta (id),
  role_id bigint  not null REFERENCES role (id),
  CONSTRAINT user_role_pkey PRIMARY KEY (user_id, role_id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE user_role
  OWNER TO postgres;
