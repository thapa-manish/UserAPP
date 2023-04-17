package db_test

// import (
// 	"database/sql"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"testing"

// 	_ "github.com/mattn/go-sqlite3"
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"

// 	"use/db"
// )

// func TestDb(t *testing.T) {
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "Db Suite")
// }

// var _ = Describe("RunMigration", func() {
// 	var (
// 		testDb *sql.DB
// 	)

// 	BeforeEach(func() {
// 		var err error
// 		testDb, err = sql.Open("sqlite3", ":memory:")
// 		Expect(err).NotTo(HaveOccurred())
// 	})

// 	AfterEach(func() {
// 		testDb.Close()
// 	})

// 	It("applies the user.sql schema to the database", func() {
// 		// Load the expected schema from file.
// 		schemaBytes, err := ioutil.ReadFile("./db/user.sql")
// 		Expect(err).NotTo(HaveOccurred())
// 		expectedSchema := string(schemaBytes)

// 		// Apply the schema using RunMigration().
// 		db.RunMigration(testDb)

// 		// Verify that the expected schema is present in the database.
// 		var actualSchema string
// 		rows, err := testDb.Query("SELECT sql FROM sqlite_master WHERE type = 'table'")
// 		Expect(err).NotTo(HaveOccurred())
// 		defer rows.Close()
// 		for rows.Next() {
// 			var schema string
// 			err := rows.Scan(&schema)
// 			Expect(err).NotTo(HaveOccurred())
// 			actualSchema += schema
// 		}
// 		Expect(actualSchema).To(Equal(expectedSchema))
// 	})

// 	Context("when the schema file is missing", func() {
// 		It("logs an error and returns without applying the schema", func() {
// 			// Temporarily remove the schema file.
// 			err := os.Remove("./db/user.sql")
// 			Expect(err).NotTo(HaveOccurred())

// 			// Call RunMigration() and verify that it logs an error and returns.
// 			var loggedError string
// 			log.SetOutput(&fakeLogWriter{&loggedError})
// 			db.RunMigration(testDb)
// 			Expect(loggedError).To(ContainSubstring("cant read sql file"))

// 			// Restore the schema file.
// 			err = ioutil.WriteFile("./db/user.sql", []byte(""), 0644)
// 			Expect(err).NotTo(HaveOccurred())
// 		})
// 	})
// })

// type fakeLogWriter struct {
// 	logString *string
// }

// func (w *fakeLogWriter) Write(p []byte) (int, error) {
// 	*w.logString += string(p)
// 	return len(p), nil
// }
