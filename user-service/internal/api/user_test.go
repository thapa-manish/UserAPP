package api_test

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"

// 	"use/internal/api"
// 	"use/internal/model"
// 	"use/internal/service"

// 	"github.com/labstack/echo/v4"
// 	"github.com/onsi/ginkgo"
// 	"github.com/onsi/gomega"
// )

// var _ = ginkgo.Describe("UserAPI", func() {
// 	var (
// 		userService *service.UserService
// 		userAPI     *api.UserAPI
// 		e           *echo.Echo
// 	)

// 	ginkgo.BeforeEach(func() {
// 		userService = service.NewUserService(nil)
// 		userAPI = api.NewUserAPI(userService)
// 		e = echo.New()
// 	})

// 	ginkgo.Describe("ListUsers", func() {
// 		ginkgo.Context("when the service returns users", func() {
// 			ginkgo.BeforeEach(func() {
// 				userService.ListUsersFunc = func(page, perPage uint64) ([]*model.User, error) {
// 					return []*model.User{
// 						{ID: 1, Name: "Alice"},
// 						{ID: 2, Name: "Bob"},
// 					}, nil
// 				}
// 			})

// 			ginkgo.It("should return a list of users", func() {
// 				req := httptest.NewRequest(http.MethodGet, "/users?page=1&per_page=2", nil)
// 				rec := httptest.NewRecorder()
// 				c := e.NewContext(req, rec)

// 				err := userAPI.ListUsers(c)
// 				gomega.Expect(err).NotTo(gomega.HaveOccurred())
// 				gomega.Expect(rec.Code).To(gomega.Equal(http.StatusOK))

// 				var users []*model.User
// 				err = json.Unmarshal(rec.Body.Bytes(), &users)
// 				gomega.Expect(err).NotTo(gomega.HaveOccurred())
// 				gomega.Expect(users).To(gomega.HaveLen(2))
// 				gomega.Expect(users[0].ID).To(gomega.Equal(int64(1)))
// 				gomega.Expect(users[0].Name).To(gomega.Equal("Alice"))
// 				gomega.Expect(users[1].ID).To(gomega.Equal(int64(2)))
// 				gomega.Expect(users[1].Name).To(gomega.Equal("Bob"))
// 			})
// 		})

// 		ginkgo.Context("when the service returns an error", func() {
// 			ginkgo.BeforeEach(func() {
// 				userService.ListUsersFunc = func(page, perPage uint64) ([]*model.User, error) {
// 					return nil, service.ErrInternal
// 				}
// 			})

// 			ginkgo.It("should return an error response", func() {
// 				req := httptest.NewRequest(http.MethodGet, "/users?page=1&per_page=2", nil)
// 				rec := httptest.NewRecorder()
// 				c := e.NewContext(req, rec)

// 				err := userAPI.ListUsers(c)
// 				gomega.Expect(err).NotTo(gomega.HaveOccurred())
// 				gomega.Expect(rec.Code).To(gomega.Equal(http.StatusInternalServerError))

// 				var body map[string]string
// 				err = json.Unmarshal(rec.Body.Bytes(), &body)
// 				gomega.Expect(err).NotTo(gomega.HaveOccurred())
// 				gomega.Expect(body["error"]).To(gomega.Equal(service.ErrInternal.Error()))
// 			})
// 		})
// 	})
// })
