package middleware_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"use/internal/middleware"
	"use/internal/model"
)

var _ = Describe("UserValidator middleware", func() {
	var (
		e    *echo.Echo
		req  *http.Request
		rec  *httptest.ResponseRecorder
		ctx  echo.Context
		next echo.HandlerFunc
	)

	BeforeEach(func() {
		e = echo.New()
		req = httptest.NewRequest(http.MethodPost, "/", nil)
		rec = httptest.NewRecorder()
		ctx = e.NewContext(req, rec)
		next = func(c echo.Context) error {
			return c.NoContent(http.StatusOK)
		}
	})

	It("should allow valid requests", func() {
		user := model.User{
			Email:     "test@example.com",
			FirstName: "John",
			LastName:  "Doe",
		}
		bodyBytes, err := json.Marshal(user)
		Expect(err).NotTo(HaveOccurred())

		req.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mw := middleware.UserValidator(next)

		err = mw(ctx)
		Expect(err).NotTo(HaveOccurred())

		Expect(rec.Code).To(Equal(http.StatusOK))
	})

	It("should return an error for requests with invalid email", func() {
		user := model.User{
			Email:     "invalid-email",
			FirstName: "John",
			LastName:  "Doe",
		}
		bodyBytes, err := json.Marshal(user)
		Expect(err).NotTo(HaveOccurred())

		req.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mw := middleware.UserValidator(next)

		err = mw(ctx)
		Expect(err).To(HaveOccurred())

		Expect(rec.Code).To(Equal(http.StatusUnprocessableEntity))
		Expect(rec.Body.String()).To(Equal(`{"email":"email is invalid."}`))
	})

	It("should return an error for requests missing required fields", func() {
		user := model.User{
			FirstName: "John",
			LastName:  "Doe",
		}
		bodyBytes, err := json.Marshal(user)
		Expect(err).NotTo(HaveOccurred())

		req.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		mw := middleware.UserValidator(next)

		err = mw(ctx)
		Expect(err).To(HaveOccurred())

		Expect(rec.Code).To(Equal(http.StatusUnprocessableEntity))
		Expect(rec.Body.String()).To(Equal(`{"email":"email is a required field."}`))
	})
})
