module main

go 1.14

require (
	github.com/SherClockHolmes/webpush-go v1.1.2 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/jinzhu/gorm v1.9.15 // indirect
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	local.packages/domains v0.0.0-00010101000000-000000000000
	local.packages/lib v0.0.0-00010101000000-000000000000 // indirect
	local.packages/models v0.0.0-00010101000000-000000000000 // indirect
)

replace local.packages/domains => ./domains

replace local.packages/models => ./models

replace local.packages/lib => ./lib
