package auth

import (
	"fotongo/app/modules/auth/dtos"
	"fotongo/infrastructure/services/prisma/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Auth", func() {
	var (
		dbTest *db.PrismaClient
	)

	BeforeEach(func() {
		dbTest = db.NewClient()
		dbTest.Prisma.Connect()
	})

	Describe("Register", func() {
		It("Positive Case - User data", func() {
			service := NewAuthService(dbTest)
			request := dtos.RegisterRequest{
				Username: "test",
				Email:    "test@gmail.com",
				Password: "lalayeye",
			}

			resp, err := service.Register(request)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

	})

	// Describe("Login", func() {
	// 	It("Negative Case - Returning 400 invalid email or password", func() {

	// 		service := authService.NewAuthService(db)

	// 		request := dtos.LoginRequest{
	// 			Email:    "test@gmail.com",
	// 			Password: "lalayeye",
	// 		}

	// 		_, err := service.Login(request)
	// 		Expect(err).Should(HaveOccurred())
	// 		//Expect(resp.Code).Should(Equal(500))
	// 	})

	// 	It("Positive Case - Return JWT Token", func() {
	// 		service := authService.NewAuthService(db)
	// 		request := dtos.LoginRequest{
	// 			Email:    "test@gmail.com",
	// 			Password: "lalayeye",
	// 		}

	// 		resp, err := service.Login(request)
	// 		Expect(err).Should(Succeed())
	// 		Expect(resp.Code).Should(Equal(200))
	// 	})

	// })
})
