module main

go 1.17

require (
	example.com/app v0.0.0
	example.com/controllers v0.0.0
	example.com/routes v0.0.0
	example.com/middlewares v0.0.0
	example.com/helpers v0.0.0
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
)

replace example.com/app => ./app
replace example.com/controllers => ./app/controllers
replace example.com/middlewares => ./app/middlewares
replace example.com/helpers => ./app/helpers
replace example.com/routes => ./app/routes
