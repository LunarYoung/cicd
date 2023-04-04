package e2e

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("b E2E test", func() {
	Context("when ", Ordered, func() {
		It("测试,返回200", func() {
			url := fmt.Sprintf("%sping", "http://localhost/")
			resp, err := resty.New().R().
				Get(url)
			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		})
	})
})
