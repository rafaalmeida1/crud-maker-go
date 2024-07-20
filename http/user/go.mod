module http

go 1.18

require (
	github.com/gorilla/mux v1.8.1
	github.com/lib/pq v1.10.9
	http/traits v0.0.0-00010101000000-000000000000
)

replace http/traits => ../traits
