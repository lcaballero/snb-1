SELECT *
FROM information_schema.tables
WHERE
   table_schema='public'
   and table_catalog=$1
   and table_name=$2;