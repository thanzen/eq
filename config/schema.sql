CREATE TABLE user_meta
(
  id serial NOT NULL,
  first_name text,
  last_name text,
  user_name text,
  email text,
  age integer,
  country text,
  city text,
  post_code text,
  cell_phone text,
  home_phone text,
  office_phone text,
  fax text,
  CONSTRAINT user_meta_pkey PRIMARY KEY (id)
);
 CREATE INDEX index_user_meta_on_id_type ON user_meta USING btree (id);
 