-- upgrade to version 2

DO $$
DECLARE
	next_version INTEGER := 2;
BEGIN
	
	IF EXISTS(SELECT * FROM db_version WHERE version >= next_version)
	THEN
		-- we're on or beyond this version already
		RAISE NOTICE 'Skipping version %', next_version;
	ELSE
		-- execute migration
		ALTER TABLE impressions ADD COLUMN is_leaving BOOLEAN;
		UPDATE impressions SET is_leaving = 'f';
		ALTER TABLE impressions ALTER COLUMN is_leaving SET DEFAULT FALSE;
		ALTER TABLE impressions ALTER COLUMN is_leaving SET NOT NULL;

		-- we're done. set current version
		UPDATE db_version SET version = next_version;
	END IF;
END $$;
