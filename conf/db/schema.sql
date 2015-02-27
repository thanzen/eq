CREATE SEQUENCE global_id_sequence
  INCREMENT 1
  MINVALUE 1
  MAXVALUE 9223372036854775807
  START 1043
  CACHE 1;

  -- Function: id_generator()

  -- DROP FUNCTION id_generator();

  CREATE OR REPLACE FUNCTION id_generator(OUT id bigint)
    RETURNS bigint AS
  $BODY$
  DECLARE
      our_epoch bigint := 1314220021721;
      seq_id bigint;
      now_millis bigint;
      -- the id of this DB shard, must be set for each
      -- schema shard you have - you could pass this as a parameter too
      shard_id bigint := 1;
  BEGIN
      SELECT nextval('global_id_sequence') % 1024 INTO seq_id;

      SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000) INTO now_millis;
      id  := (now_millis - our_epoch) << 23;
      id  := id  | (shard_id << 10);
      id  := id  | (seq_id);
      --return id;
  END;
  $BODY$
    LANGUAGE plpgsql VOLATILE
    COST 100;
  ALTER FUNCTION id_generator()
    OWNER TO postgres;

-- Table: permission

-- DROP TABLE permission;

CREATE TABLE permission
(
  id bigint NOT NULL DEFAULT id_generator(),
  name text,
  category text,
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
-- Table: user_type

-- DROP TABLE user_type;

CREATE TABLE user_type
(
  id bigint NOT NULL DEFAULT id_generator(),
  name text,
  CONSTRAINT user_type_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
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

-- Table: user_info

-- DROP TABLE user_info;

CREATE TABLE user_info
(
  id bigint NOT NULL DEFAULT id_generator(),
  user_name text,
  first_name text,
  last_name text,
  email text,
  password text,
  password_salt text,
  admin_comment text,
  active boolean,
  deleted boolean,
  last_ip text,
  created timestamp without time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  last_login_date timestamp without time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  last_activity_date timestamp without time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  cell_phone text,
  office_phone text,
  fax text,
  country text,
  city text,
  post_code text,
  user_type_id bigint,
  lang text NOT NULL DEFAULT 'en-US'::text,
  gravatar_email text,
  updated timestamp without time zone NOT NULL DEFAULT timezone('utc'::text, now()),
  CONSTRAINT user_info_pkey PRIMARY KEY (id),
  CONSTRAINT user_info_user_type_fk FOREIGN KEY (user_type_id)
      REFERENCES user_type (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH (
  OIDS=FALSE
);
-- Index: fki_user_info_user_type_fk

-- DROP INDEX fki_user_info_user_type_fk;

CREATE INDEX fki_user_info_user_type_fk
  ON user_info
  USING btree
  (user_type_id);

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