module app

go 1.18

require (
	github.com/gorilla/mux v1.8.1
	http/userhttp v0.0.0-00010101000000-000000000000
)

require (
	github.com/lib/pq v1.10.9 // indirect
	http/traits v0.0.0-00010101000000-000000000000 // indirect
)

replace userhttp => ./http/user

replace http/userhttp => ./http/user

replace http/traits => ./http/traits
