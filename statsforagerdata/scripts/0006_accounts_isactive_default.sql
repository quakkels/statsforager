DO $$
DECLARE
	next_version INTEGER := 6;
BEGIN
	
	IF EXISTS(SELECT * FROM db_version WHERE version >= next_version)
	THEN
		-- we're on or beyond this version already
		RAISE NOTICE 'Skipping version %', next_version;
	ELSE
		-- execute migration
		ALTER TABLE accounts ALTER COLUMN is_active SET default false;

		-- we're done. set current version
		UPDATE db_version SET version = next_version;
		RAISE NOTICE 'Update to version % COMPLETE', next_version;

	END IF;
END $$;




