

./gen --sqltype=sqlite3 --connstr "./example/sample.db" --database main --json --gorm --guregu --rest --out ./example --module example.com/rest/example --mod --server --makefile --json-fmt=snake --generate-dao --generate-proj --overwrite

./gen --sqltype=postgres --connstr "user=postgres password=postgres dbname=keycloak sslmode=disable" --database keycloak --json --gorm --guregu --rest --out ./keycloak --module example.com/rest/example --mod --server --makefile --json-fmt=snake --generate-dao --generate-proj --overwrite

./gen --sqltype=mysql --connstr "root:rootroot@tcp(localhost:3306)/diyhibbs-pro?charset=utf8mb4&parseTime=True&loc=Local" --database diyhibbs --json --gorm --guregu --rest --out ./diyhibbs --module example.com/rest/example --mod --server --makefile --json-fmt=snake --generate-dao --generate-proj --overwrite

./gen --sqltype=mysql --connstr "root:rootroot@tcp(localhost:3306)/test1?charset=utf8mb4&parseTime=True&loc=Local" --database test1 --json --gorm --rest --out ./test1 --module example.com/rest/example --mod --server --makefile --json-fmt=snake --generate-dao --generate-proj --overwrite
