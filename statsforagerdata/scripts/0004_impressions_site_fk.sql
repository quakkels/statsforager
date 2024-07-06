DO $$
DECLARE
	-- start upgrade to this version
	next_version INTEGER := 4;
BEGIN
	
	IF EXISTS(SELECT * FROM db_version WHERE version >= next_version)
	THEN
		-- we're on or beyond this version already
		RAISE NOTICE 'Skipping version %', next_version;
	ELSE
		-- execute migration
		ALTER TABLE impressions ADD COLUMN site_key uuid NOT NULL;
		ALTER TABLE impressions
			ADD CONSTRAINT fk_impressions_sites
			FOREIGN KEY (site_key) REFERENCES sites (site_key);
		-- we're done. set current version
		UPDATE db_version SET version = next_version;
	END IF;
END $$;

