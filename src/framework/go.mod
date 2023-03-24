module example

go 1.13

require geeorm v0.0.0

require (
	gee v0.0.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
)

replace geeorm => ./orm

replace gee => ./gee
