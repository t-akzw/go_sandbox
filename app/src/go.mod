module main

go 1.14

require (
	github.com/SherClockHolmes/webpush-go v1.1.2 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/jinzhu/gorm v1.9.15 // indirect
	local.packages/domains v0.0.0-00010101000000-000000000000
	local.packages/models v0.0.0-00010101000000-000000000000 // indirect
)

replace local.packages/domains => ./domains

replace local.packages/models => ./models
