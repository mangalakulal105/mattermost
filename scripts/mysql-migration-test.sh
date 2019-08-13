TMPDIR=`mktemp -d 2>/dev/null || mktemp -d -t 'tmpConfigDir'`
DUMPDIR=`mktemp -d 2>/dev/null || mktemp -d -t 'dumpDir'`

CONTAINER="$COMPOSE_PROJECT_NAME_mysql"
NETWORK="--network $COMPOSE_PROJECT_NAME_mm_test"
if [ "$IS_CI" != 'true' ]; then
        CONTAINER="mattermost-mysql"
        NETWORK=""
fi

cp config/config.json $TMPDIR

echo "Creating databases"
docker exec $NETWORK $CONTAINER mysql -uroot -pmostest -e "CREATE DATABASE migrated; CREATE DATABASE latest; GRANT ALL PRIVILEGES ON migrated.* TO mmuser; GRANT ALL PRIVILEGES ON latest.* TO mmuser"

echo "Importing mysql dump from version 5.0"
docker exec $NETWORK -i $CONTAINER mysql -D migrated -uroot -pmostest < $(pwd)/scripts/mattermost-mysql-5.0.sql

echo "Setting up config for db migration"
make ARGS="config set SqlSettings.DataSource 'mmuser:mostest@tcp(localhost:3306)/migrated?charset=utf8mb4,utf8&readTimeout=30s&writeTimeout=30s' --config $TMPDIR/config.json" run-cli
make ARGS="config set SqlSettings.DriverName 'mysql' --config $TMPDIR/config.json" run-cli

echo "Running the migration"
make ARGS="version --config $TMPDIR/config.json" run-cli

echo "Setting up config for fresh db setup"
make ARGS="config set SqlSettings.DataSource 'mmuser:mostest@tcp(localhost:3306)/latest?charset=utf8mb4,utf8&readTimeout=30s&writeTimeout=30s' --config $TMPDIR/config.json" run-cli

echo "Setting up fresh db"
make ARGS="version --config $TMPDIR/config.json" run-cli

echo "Generating dump"
docker exec $NETWORK $CONTAINER mysqldump --skip-opt --no-data --compact -u root -pmostest migrated > $DUMPDIR/migrated.sql
docker exec $NETWORK $CONTAINER mysqldump --skip-opt --no-data --compact -u root -pmostest latest > $DUMPDIR/latest.sql

echo "Removing databases created for db comparison"
docker exec $NETWORK $CONTAINER mysql -uroot -pmostest -e "DROP DATABASE migrated; DROP DATABASE latest"

echo "Generating diff"
diff $DUMPDIR/migrated.sql $DUMPDIR/latest.sql > $DUMPDIR/diff.txt
diffErrorCode=$?

if [ $diffErrorCode -eq 0 ]; then echo "Both schemas are same";else cat $DUMPDIR/diff.txt; fi
rm -rf $TMPDIR $DUMPDIR

exit $diffErrorCode
