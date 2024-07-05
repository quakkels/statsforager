-- upgrade to version 3

DO $$
DECLARE
	next_version INTEGER := 3;
BEGIN
	
	IF EXISTS(SELECT * FROM db_version WHERE version >= next_version)
	THEN
		-- we're on or beyond this version already
		RAISE NOTICE 'Skipping version %', next_version;
	ELSE
		-- execute migration
		ALTER TABLE impressions DROP CONSTRAINT impressions_pkey;
		ALTER TABLE impressions ADD PRIMARY KEY (impression_id, is_leaving);
		ALTER TABLE impressions ADD COLUMN event_date_time_utc timestamp default (timezone('utc', now())) NOT NULL;

		-- we're done. set current version
		UPDATE db_version SET version = next_version;
	END IF;
END $$;
