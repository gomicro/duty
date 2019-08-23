package config

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestRoute(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Routes", func() {
		g.Describe("Static", func() {
			g.It("should return static content for a static route", func() {
				f, _ := ParseFromFile()

				server := httptest.NewServer(f)
				defer server.Close()

				u := fmt.Sprintf("%v%v", server.URL, "/v1/static")

				res, err := http.Get(u)
				Expect(err).To(BeNil())
				defer res.Body.Close()

				firstb, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())
				Expect(string(firstb)).To(ContainSubstring("here lies a foo"))
				Expect(string(firstb)).To(ContainSubstring("git to the bar"))
				Expect(string(firstb)).To(ContainSubstring("let's get the baz back together"))

				res, err = http.Get(u)
				Expect(err).To(BeNil())
				defer res.Body.Close()

				secondb, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())
				Expect(string(secondb)).To(ContainSubstring("here lies a foo"))
				Expect(string(secondb)).To(ContainSubstring("git to the bar"))
				Expect(string(secondb)).To(ContainSubstring("let's get the baz back together"))

				Expect(firstb).To(Equal(secondb))
			})

			g.It("should treat a route as a static if no type is specified", func() {
				f, _ := ParseFromFile()

				server := httptest.NewServer(f)
				defer server.Close()

				u := fmt.Sprintf("%v%v", server.URL, "/v1/foo")

				res, err := http.Get(u)
				Expect(err).To(BeNil())
				defer res.Body.Close()

				firstb, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())
				Expect(string(firstb)).To(ContainSubstring("here lies a foo"))
				Expect(string(firstb)).To(ContainSubstring("git to the bar"))
				Expect(string(firstb)).To(ContainSubstring("let's get the baz back together"))

				res, err = http.Get(u)
				Expect(err).To(BeNil())
				defer res.Body.Close()

				secondb, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())
				Expect(string(secondb)).To(ContainSubstring("here lies a foo"))
				Expect(string(secondb)).To(ContainSubstring("git to the bar"))
				Expect(string(secondb)).To(ContainSubstring("let's get the baz back together"))

				Expect(firstb).To(Equal(secondb))
			})

			g.It("should return a response code for an endpoint with no response defined", func() {
				f, _ := ParseFromFile()

				server := httptest.NewServer(f)
				defer server.Close()

				u := fmt.Sprintf("%v%v", server.URL, "/v1/baz")

				res, err := http.Get(u)
				Expect(err).To(BeNil())
				defer res.Body.Close()

				Expect(res.StatusCode).To(Equal(401))

				u = fmt.Sprintf("%v%v", server.URL, "/v1/biz")

				res, err = http.Get(u)
				Expect(err).To(BeNil())
				defer res.Body.Close()

				Expect(res.StatusCode).To(Equal(200))
			})
		})

		g.Describe("Oridinal", func() {
			g.It("should return ordinal content for an ordinal route", func() {
				f, _ := ParseFromFile()

				server := httptest.NewServer(f)
				defer server.Close()

				u := fmt.Sprintf("%v%v", server.URL, "/v1/ordinal")

				res, err := http.Get(u)
				Expect(err).To(BeNil())
				defer res.Body.Close()

				firstb, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())
				Expect(string(firstb)).To(ContainSubstring("here lies a foo"))
				Expect(string(firstb)).To(ContainSubstring("git to the bar"))
				Expect(string(firstb)).To(ContainSubstring("let's get the baz back together"))

				res, err = http.Get(u)
				Expect(err).To(BeNil())
				defer res.Body.Close()

				secondb, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())
				Expect(string(secondb)).To(ContainSubstring("unauthorized"))

				Expect(string(firstb)).NotTo(Equal(string(secondb)))

				res, err = http.Get(u)
				Expect(err).To(BeNil())
				defer res.Body.Close()

				thirdb, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())
				Expect(string(secondb)).To(ContainSubstring("unauthorized"))

				Expect(string(firstb)).NotTo(Equal(string(thirdb)))
				Expect(string(secondb)).To(Equal(string(thirdb)))
			})

			g.It("should reset an ordinal route", func() {
				f, _ := ParseFromFile()

				server := httptest.NewServer(f)
				defer server.Close()

				u := fmt.Sprintf("%v%v", server.URL, "/v1/ordinal")

				res, err := http.Get(u)
				Expect(err).To(BeNil())
				defer res.Body.Close()

				firstb, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())
				Expect(string(firstb)).To(ContainSubstring("here lies a foo"))
				Expect(string(firstb)).To(ContainSubstring("git to the bar"))
				Expect(string(firstb)).To(ContainSubstring("let's get the baz back together"))

				res, err = http.Get(u)
				Expect(err).To(BeNil())
				defer res.Body.Close()

				secondb, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())
				Expect(string(secondb)).To(ContainSubstring("unauthorized"))

				Expect(string(firstb)).NotTo(Equal(string(secondb)))

				r := fmt.Sprintf("%v%v", server.URL, "/duty/reset")
				http.Get(r)

				res, err = http.Get(u)
				Expect(err).To(BeNil())
				defer res.Body.Close()

				thirdb, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())
				Expect(string(thirdb)).To(ContainSubstring("here lies a foo"))
				Expect(string(thirdb)).To(ContainSubstring("git to the bar"))
				Expect(string(thirdb)).To(ContainSubstring("let's get the baz back together"))

				Expect(string(firstb)).To(Equal(string(thirdb)))
				Expect(string(secondb)).NotTo(Equal(string(thirdb)))

				res, err = http.Get(u)
				Expect(err).To(BeNil())
				defer res.Body.Close()

				fourthb, err := ioutil.ReadAll(res.Body)
				Expect(err).To(BeNil())
				Expect(string(fourthb)).To(ContainSubstring("unauthorized"))

				Expect(string(firstb)).NotTo(Equal(string(fourthb)))
				Expect(string(secondb)).To(Equal(string(fourthb)))
				Expect(string(thirdb)).NotTo(Equal(string(fourthb)))
			})
		})
	})
}
