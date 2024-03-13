package models_test

import (
	"battle-of-monsters/app/db"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}

var _ = BeforeSuite(func() {
	db.Connect()
	db.Up()
})

var _ = AfterSuite(func() {
	db.Down()
})
