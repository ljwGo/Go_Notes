module go/hello

go 1.24.3

replace go/greetings => ./greetings

require go/greetings v0.0.0-00010101000000-000000000000

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.2 // indirect
)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	golang.org/x/tools v0.33.0
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.3.0 // indirect
)
