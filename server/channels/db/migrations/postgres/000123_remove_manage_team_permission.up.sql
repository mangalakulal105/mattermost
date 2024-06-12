DO $$
<<remove_manage_team_permission>>
DECLARE
  id_offset text := '';
  rows_updated integer;
BEGIN
  LOOP
    WITH table_holder AS (
      SELECT id FROM roles
          WHERE id > id_offset
          ORDER BY id ASC limit 100
    )
    UPDATE Roles r set Permissions = REGEXP_REPLACE(Permissions, 'manage_team(\?|\s)', '') 
        WHERE r.id in (SELECT id FROM table_holder)
            AND Permissions ~~ '%manage_team%'
            AND (Permissions ~~ '%sysconsole_write_user_management_chanels%'
            OR Permissions ~~ '%sysconsole_write_user_management_groups%')
            AND Permissions !~~ '%sysconsole_write_user_management_teams%';
    GET DIAGNOSTICS rows_updated = ROW_COUNT;

     -- We have to run the select query again
     -- becaue "select into" isn't allowed inside a CTE
     -- and without CTE, we have to use a temp table (because you can't select into a table)
     -- and with a temp table, you run into max_locks_inside_transaction limit.
     -- Probably there is a better way but keeping things simple for now.
    SELECT id INTO id_offset FROM (
        SELECT id from roles 
            WHERE id > id_offset
            ORDER BY id ASC limit 100
    ) as temp order by id desc limit 1;

    EXIT WHEN rows_updated = 0;
  END LOOP;
END remove_manage_team_permission $$;