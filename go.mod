module github.com/Sankara98/backing-catalog

replace github.com/Sankara98/backing-fulfillment/service => ./service

go 1.13

require (
	github.com/cloudfoundry-community/go-cfenv v1.18.0
	github.com/cloudnativego/cf-tools v0.0.0-20160521031255-f59655db37a9
	github.com/gorilla/mux v1.7.4
	github.com/unrolled/render v1.0.2
	github.com/urfave/negroni v1.0.0
)
