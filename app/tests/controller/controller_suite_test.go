package controller_test

import (
	"battle-of-monsters/app/db"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = BeforeSuite(func() {
	db.Connect()
	db.Up()
})

var _ = AfterSuite(func() {
	db.Down()
})
