migrate \
  -path "${PWD}" \
  -database "postgres://erikrios:erikrios@localhost:5432/ponorogo_regency_db?sslmode=disable" \
  up
