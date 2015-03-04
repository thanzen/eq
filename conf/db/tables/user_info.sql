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
  company text,
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

