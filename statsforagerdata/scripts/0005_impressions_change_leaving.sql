DO $$
DECLARE
	next_version INTEGER := 5;
BEGIN
	
	IF EXISTS(SELECT * FROM db_version WHERE version >= next_version)
	THEN
		-- we're on or beyond this version already
		RAISE NOTICE 'Skipping version %', next_version;
	ELSE
		-- execute migration
		ALTER TABLE impressions DROP CONSTRAINT impressions_pkey;
		ALTER TABLE impressions ADD PRIMARY KEY (impression_id);
		ALTER TABLE impressions 
			ADD COLUMN created_utc timestamp default (timezone('utc', now())) NOT NULL;
		ALTER TABLE impressions 
			ADD COLUMN started_utc timestamp default (timezone('utc', now())) NOT NULL;
		ALTER TABLE impressions 
			ADD COLUMN completed_utc timestamp default (timezone('utc', now())) NOT NULL;
		ALTER TABLE impressions DROP COLUMN date_time;
		ALTER TABLE impressions DROP COLUMN is_leaving;
		ALTER TABLE impressions DROP COLUMN event_date_time_utc;
		-- we're done. set current version
		UPDATE db_version SET version = next_version;
		RAISE NOTICE 'Update to version % COMPLETE', next_version;
	END IF;
END $$;



