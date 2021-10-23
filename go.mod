module github.com/hmhuan/climate-news-api

go 1.16

require (
	api v1.0.0
	github.com/gin-gonic/gin v1.7.4
	golang.org/x/sys v0.0.0-20210423082822-04245dca01da // indirect
)

replace api v1.0.0 => ./src/api
