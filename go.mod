module github.com/sakuraapp/shared

go 1.18

replace github.com/sakuraapp/pubsub => /Users/jackie/Documents/Projects/Sakura/Source/pubsub

require (
	github.com/go-chi/jwtauth/v5 v5.0.2
	github.com/go-chi/render v1.0.1
	github.com/go-redis/redis/v8 v8.11.4
	github.com/lestrrat-go/jwx v1.2.18
	github.com/vmihailenco/msgpack/v5 v5.3.5
	gopkg.in/guregu/null.v4 v4.0.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/sakuraapp/pubsub v0.0.0-20230313165424-054ed0eba8a5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
)
