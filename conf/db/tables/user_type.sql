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
