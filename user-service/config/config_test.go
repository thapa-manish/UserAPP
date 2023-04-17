package config_test

// import (
// 	"errors"
// 	"testing"

// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// 	"github.com/spf13/viper"

// 	"use/config"
// )

// func TestConfig(t *testing.T) {
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "Config Suite")
// }

// var _ = Describe("NewConfig", func() {
// 	var (
// 		v          *viper.Viper
// 		configFile string
// 	)

// 	BeforeEach(func() {
// 		v = viper.New()
// 		configFile = "../env/.env"
// 	})

// 	AfterEach(func() {
// 		v = nil
// 		configFile = ""
// 	})

// 	Context("when the configuration file is present", func() {
// 		BeforeEach(func() {
// 			v.SetConfigFile(configFile)
// 			err := v.ReadInConfig()
// 			Expect(err).NotTo(HaveOccurred())
// 		})

// 		It("returns a valid Config object with the expected values", func() {
// 			expectedConfig := &config.Config{
// 				DbUser:     "postgresUser",
// 				DbPassword: "postgresPW",
// 				DbName:     "postgresDB",
// 				DbHost:     "localhost",
// 				DbPort:     "5432",
// 				HttpPort:   "8080",
// 			}

// 			c, err := config.NewConfig()

// 			Expect(err).NotTo(HaveOccurred())
// 			Expect(c).To(Equal(expectedConfig))
// 		})

// 	})

// 	Context("when the configuration file is missing", func() {

// 		BeforeEach(func() {
// 			v.SetConfigFile("./env/missing.env")
// 		})

// 		It("returns an error", func() {
// 			_, err := config.NewConfig()

// 			Expect(err).To(MatchError(errors.New("config file not found")))
// 		})

// 	})

// })
