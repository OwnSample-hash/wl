#!/bin/bash

getValues() {
  python -c "import yaml,sys;eval(f'print(yaml.safe_load(open(sys.argv[1],\"r\")){\"\".join([f\"[\\\"{x}\\\"]\"for x in sys.argv[2].split(\".\")])})')" $1 $2
}

CONFIG="config.yaml"
DB_HOST=$(getValues $CONFIG "db.host")
DB_PORT=$(getValues $CONFIG "db.port")
DB_USER=$(getValues $CONFIG "db.user")
DB_PASS=$(getValues $CONFIG "db.pass")
DB_DBNAME=$(getValues $CONFIG "db.database")

echo "DB_HOST: $DB_HOST"
echo "DB_PORT: $DB_PORT"
echo "DB_USER: $DB_USER"
echo "DB_PASS: $DB_PASS"
echo "DB_DBNAME: $DB_DBNAME"

# Check if the database is up
while ! mariadb -h $DB_HOST -P $DB_PORT -u $DB_USER -p$DB_PASS -e "USE $DB_DBNAME;" 2>/dev/null; do
  echo "Waiting for database to be up..."
  sleep 1
done

# Run the database migration

echo "Running database migration..."

./main -migrate

# Check if the migration was successful
if [ $? -ne 0 ]; then
  echo "Database migration failed."
  exit 1
fi

for seed in $(ls Seeds);do
  echo "Running seed: $seed"
  mariadb -h $DB_HOST -P $DB_PORT -u $DB_USER -p$DB_PASS $DB_DBNAME < Seeds/$seed
  if [ $? -ne 0 ]; then
    echo "Seeding failed for $seed."
    exit 1
  fi
done

./main
