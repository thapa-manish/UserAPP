package repository_test

// import (
// 	"database/sql"
// 	"errors"
// 	"fmt"

// 	"use/internal/model"
// 	"use/internal/repository"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/onsi/ginkgo"
// 	"github.com/onsi/gomega"
// )

// var _ = ginkgo.Describe("User Repository", func() {
// 	var (
// 		db             *sql.DB
// 		mock           sqlmock.Sqlmock
// 		userRepo       *repository.UserRepository
// 		expectedUsers  []model.User
// 		expectedUser   *model.User
// 		errExpected    error
// 		errNotExpected error
// 	)

// 	ginkgo.BeforeEach(func() {
// 		var err error
// 		db, mock, err = sqlmock.New()
// 		gomega.Expect(err).NotTo(gomega.HaveOccurred())

// 		userRepo = repository.NewUserRepository(db)

// 		expectedUsers = []model.User{
// 			model.User{
// 				ID:        1,
// 				Email:     "user1@example.com",
// 				FirstName: "John",
// 				LastName:  "Doe",
// 			},
// 			model.User{
// 				ID:        2,
// 				Email:     "user2@example.com",
// 				FirstName: "Jane",
// 				LastName:  "Doe",
// 			},
// 		}

// 		expectedUser = &model.User{
// 			ID:        1,
// 			Email:     "user1@example.com",
// 			FirstName: "John",
// 			LastName:  "Doe",
// 		}

// 		errExpected = errors.New("expected error")
// 		errNotExpected = errors.New("not expected error")
// 	})

// 	ginkgo.AfterEach(func() {
// 		err := mock.ExpectationsWereMet()
// 		gomega.Expect(err).NotTo(gomega.HaveOccurred())
// 		db.Close()
// 	})

// 	ginkgo.Describe("FindAll", func() {
// 		ginkgo.Context("when users exist in the database", func() {
// 			ginkgo.It("should return a list of users", func() {
// 				rows := sqlmock.NewRows([]string{"id", "email", "first_name", "last_name"}).
// 					AddRow(expectedUsers[0].ID, expectedUsers[0].Email, expectedUsers[0].FirstName, expectedUsers[0].LastName).
// 					AddRow(expectedUsers[1].ID, expectedUsers[1].Email, expectedUsers[1].FirstName, expectedUsers[1].LastName)
// 				mock.ExpectQuery("SELECT").WillReturnRows(rows)

// 				users, err := userRepo.FindAll(1, 10)
// 				gomega.Expect(err).NotTo(gomega.HaveOccurred())
// 				gomega.Expect(users).To(gomega.Equal(expectedUsers))
// 			})
// 		})

// 		ginkgo.Context("when an error occurs querying the database", func() {
// 			ginkgo.It("should return an error", func() {
// 				mock.ExpectQuery("SELECT").WillReturnError(errExpected)

// 				_, err := userRepo.FindAll(1, 10)
// 				gomega.Expect(err).To(gomega.Equal(fmt.Errorf("failed to query users: %v", errExpected)))
// 			})
// 		})
// 	})
// })
