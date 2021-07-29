module github.com/jjonline/study_golang/validator

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/go-playground/validator/v10 v10.4.1
	github.com/jjonline/go-lib-backend/validation4gin v0.0.1
)

replace github.com/jjonline/go-lib-backend/validation4gin => ../../go-lib-backend/validation4gin
