package tests

import (
	"github.com/kukkar/mta-hosting-optimizer/tests/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var getUnusedHosts = func() {

	Context("Try to get Unused Hosts with success", func() {
		It("should return valid unused hosts data successfully", func() {
			var resHolder ResUnusedHosts
			custHeader := make(map[string]string, 0)
			endPoint := apiUnusedHosts
			err := common.HttpRequest(common.HTTP_REQUEST_GET, endPoint+"?threshhold=1", custHeader, nil, &resHolder)
			if err != nil {
				Expect(err).NotTo(HaveOccurred())
			}
			Expect(resHolder.Status).To(Equal(true))
		})
	})
}
