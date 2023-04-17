package repository_test

import (
	"database/sql"
	"fmt"
	"testing"
	"use/internal/model"
	"use/internal/repository"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

var _ = Describe("User Repository", func() {
	var (
		db             *sql.DB
		mock           sqlmock.Sqlmock
		userRepo       repository.IUserRepository
		testUsers  []model.User
	)

	BeforeEach(func() {
		var err error
		db, mock, err = sqlmock.New()
		if err != nil {
			Fail("Failed to create mock DB")
		}
		userRepo = repository.NewUserRepository(db)
		testUsers = make([]model.User, 0)
		for i := 0; i <= 100; i++ {
			testUser := model.User{
				ID:         int64(i),
				Email:      fmt.Sprintf("user_%d@example.com", i),
				UserName:   fmt.Sprintf("user_%d", i),
				FirstName:  fmt.Sprintf("user_%d", i),
				LastName:   "last_name",
				UserStatus: "",
				Department: "dept-1",
			}
			testUsers = append(testUsers, testUser)
		}
	})

	AfterEach(func() {
		err := mock.ExpectationsWereMet()
		Expect(err).NotTo(HaveOccurred())
		db.Close()
	})

	Describe("FindByID", func() {
		Context("when the user exists", func() {
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id", "user_name", "email", "first_name", "last_name", "user_status", "department"}).
					AddRow(testUsers[0].ID, testUsers[0].UserName, testUsers[0].Email, testUsers[0].FirstName, testUsers[0].LastName, testUsers[0].UserStatus, testUsers[0].Department)
				mock.ExpectQuery("^SELECT").WithArgs(testUsers[0].ID).WillReturnRows(rows)
			})

			It("should return the user", func() {
				user, err := userRepo.FindByID(testUsers[0].ID)
				Expect(err).ToNot(HaveOccurred())
				Expect(*user).To(Equal(testUsers[0]))
			})
		})

		Context("when the user does not exist", func() {
			BeforeEach(func() {
				mock.ExpectQuery("^SELECT").WithArgs(testUsers[0].ID).WillReturnError(sql.ErrNoRows)
			})

			It("should return an error", func() {
				_, err := userRepo.FindByID(testUsers[0].ID)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("user not found"))
			})
		})
	})

	Describe("FindAll", func() {
		Context("when the users exists in the database", func() {
			BeforeEach(func() {
				rows := sqlmock.NewRows([]string{"id", "user_name", "email", "first_name", "last_name", "user_status", "department"})
				for _, user := range testUsers[:100] {
					rows.AddRow(user.ID, user.UserName, user.Email, user.FirstName, user.LastName, user.UserStatus, user.Department)
				}
				mock.ExpectQuery("SELECT").WillReturnRows(rows)
			})
			It("should return all users", func() {
				users, err := userRepo.FindAll(0, 100)
				Expect(err).ToNot(HaveOccurred())
				Expect(users).To(Equal(testUsers[:100]))
			})
		})

		Context("when error in sql connection", func() {
			BeforeEach(func() {
				mock.ExpectQuery("SELECT").WillReturnError(sql.ErrConnDone)
			})
			It("should return first 10 users", func() {
				_, err := userRepo.FindAll(1, 10)
				Expect(err).To(HaveOccurred())
				expectedErr := fmt.Errorf("failed to query users: %v", sql.ErrConnDone)
				Expect(err.Error()).To(Equal(expectedErr.Error()))
			})
		})
	})
})