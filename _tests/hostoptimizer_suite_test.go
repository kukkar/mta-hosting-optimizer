package tests

import (
	"fmt"
	"testing"

	//	"github.com/bharatpe/merchant/src/merchant/unittests"
	config "github.com/kukkar/mta-hosting-optimizer/tests/conf"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var apiUnusedHosts string

func TestTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tests Suite")
}

var _ = BeforeSuite(func() {
	con, err := config.GetConfig()
	if err != nil {
		fmt.Println("Error in fetching configuration")
		panic(err)
	}
	//	unittests.CleanUPTestData(testUser)
	apiUnusedHosts = con.MTAHostingOptimizer.IPPort + "/mta-hosting-optimizer/v1/listunusedhost"
})
