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
