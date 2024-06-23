-- first sql migration file
-- Command to run files:
--	psql -h {host} -U {username} -d stats -a -f deploy.sql

BEGIN;
\i scripts/0000.sql
\i scripts/0001.sql
\i scripts/0002.sql
COMMIT;

