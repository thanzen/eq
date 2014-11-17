-- Table: role_permission

-- DROP TABLE role_permission;

CREATE TABLE role_permission
(
  role_id bigint not null REFERENCES role (id),
  permssion_id bigint  not null REFERENCES permssion (id),
  CONSTRAINT user_role_pkey PRIMARY KEY (role_id, permssion_id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE role_permission
  OWNER TO postgres;
