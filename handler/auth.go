package handler

import (
	"fmt"
	model "fotongo/model"
	"fotongo/provider"
	utils "fotongo/utils"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthInput struct {
	Name            string `json:"name" form:"name"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginIdentity struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// LoginHandler is for handle user login
func LoginHandler(c *fiber.Ctx) error {

	loginInput := new(AuthInput)
	if err := c.BodyParser(loginInput); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	validate, isValid := isLoginInputValid(AuthInput{
		Email:    loginInput.Email,
		Password: loginInput.Password,
	})
	if !isValid {
		return utils.ResponseValidationError(c, validate, "Validation Error")
	}

	findUserByEmail, err := model.FindUserByEmail(loginInput.Email)
	if err != nil {
		log.Println(err)
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(findUserByEmail.Password), []byte(loginInput.Password))
	if err != nil {
		log.Println(err)
		return utils.ResponseError(c, nil, "Password anda salah")
	}

	token, err := generateToken(LoginIdentity{
		ID:    findUserByEmail.ID,
		Name:  findUserByEmail.Name,
		Email: findUserByEmail.Email,
	}, "user")

	return utils.ResponseSuccess(c, token, "Login sukses!")
}

func RegisterHandler(c *fiber.Ctx) error {
	registerInput := new(AuthInput)
	if err := c.BodyParser(registerInput); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	validate, isValid := isRegisterInputValid(AuthInput{
		Name:            registerInput.Name,
		Email:           registerInput.Email,
		Password:        registerInput.Password,
		ConfirmPassword: registerInput.ConfirmPassword,
	})
	if !isValid {
		return utils.ResponseValidationError(c, validate, "Validation Error")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), 14)
	if err != nil {
		return err
	}

	id, err := model.InsertUser(model.User{
		Name:     registerInput.Name,
		Email:    registerInput.Email,
		Password: string(hashPassword),
		Level:    1,
	})
	if err != nil {
		return err
	}

	type response struct {
		ID *string `json:"id"`
	}

	return utils.ResponseSuccess(c, response{ID: id}, "Registrasi sukses!")
}

func CheckTokenHandler(c *fiber.Ctx) error {
	user, err := provider.GetUser(c)
	if err != nil {
		log.Println(err)
		return utils.ResponseError(c, "Token invalid", "Something went wrong")
	}

	return utils.ResponseSuccess(c, &LoginIdentity{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, "Token valid!")
}

func RequestTokenHandler(c *fiber.Ctx) error {
	jwtKey := os.Getenv("JWT_KEY")
	tokenString := c.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		log.Println(err)
		return utils.ResponseError(c, err.Error, "Something went wrong")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return utils.ResponseError(c, err.Error, "Something went wrong")
	}

	userID := claims["identity"].(map[string]interface{})["id"]
	email := claims["identity"].(map[string]interface{})["email"]
	name := claims["identity"].(map[string]interface{})["name"]

	genToken, err := generateToken(LoginIdentity{
		ID:    userID.(string),
		Name:  name.(string),
		Email: email.(string),
	}, "user")
	if err != nil {
		log.Println(err)
		return utils.ResponseError(c, err.Error, "Something went wrong")
	}

	return utils.ResponseSuccess(c, genToken, "Request access token success!")
}

func generateToken(identity LoginIdentity, role string) (*Token, error) {
	secretKeyAt := os.Getenv("JWT_KEY")
	secretKeyRt := os.Getenv("JWT_KEY")

	// at is access token
	at := jwt.New(jwt.SigningMethodHS256)
	atClaims := at.Claims.(jwt.MapClaims)
	atClaims["identity"] = LoginIdentity{
		ID:    identity.ID,
		Email: identity.Email,
		Name:  identity.Name,
	}

	if role == "admin" {
		atClaims["admin"] = true
		atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	} else {
		atClaims["admin"] = false
		atClaims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	}

	atHashed, err := at.SignedString([]byte(secretKeyAt))
	if err != nil {
		return nil, err
	}

	// rt is refresh token
	rt := jwt.New(jwt.SigningMethodHS256)
	rtClaims := rt.Claims.(jwt.MapClaims)
	rtClaims["identity"] = LoginIdentity{
		ID:    identity.ID,
		Email: identity.Email,
		Name:  identity.Name,
	}

	if role == "admin" {
		rtClaims["admin"] = true
	} else {
		rtClaims["admin"] = false
	}

	rtClaims["exp"] = time.Now().Add(time.Hour * 4320).Unix()

	rtHashed, err := rt.SignedString([]byte(secretKeyRt + "gandos"))
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  atHashed,
		RefreshToken: rtHashed,
	}, nil
}

func isLoginInputValid(input AuthInput) (*AuthInput, bool) {
	var validate AuthInput
	isValid := true

	if len(input.Email) < 1 {
		validate.Email = "Email tidak boleh kosong"
		isValid = false
	}

	if len(input.Name) > 255 {
		validate.Name = "Email terlalu panjang maksimal 255"
	}

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(input.Email) {
		validate.Email = "Format email tidak valid"
		isValid = false
	}

	if len(input.Password) < 1 {
		validate.Password = "Password tidak boleh kosong"
		isValid = false
	}

	return &validate, isValid
}

func isRegisterInputValid(input AuthInput) (*AuthInput, bool) {
	var validate AuthInput
	isValid := true

	if len(input.Name) < 2 {
		validate.Name = "Nama tidak boleh kurang dari 2"
		isValid = false
	}

	if len(input.Name) > 255 {
		validate.Name = "Nama terlalu panjang maksimal 255"
	}

	if len(input.Email) < 1 {
		validate.Email = "Email tidak boleh kosong"
		isValid = false
	}

	if len(input.Name) > 255 {
		validate.Name = "Email terlalu panjang maksimal 255"
	}

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(input.Email) {
		validate.Email = "Format email tidak valid"
		isValid = false
	}

	findUserByEmail, err := model.FindUserByEmail(input.Email)
	if err != nil {
		log.Println(err)
		return &validate, isValid
	}

	if len(findUserByEmail.ID) > 1 {
		validate.Email = "Email sudah digunakan"
		isValid = false
	}

	if len(input.Password) < 6 {
		validate.Password = "Password tidak boleh kurang dari 6"
		isValid = false
	}

	if input.Password != input.ConfirmPassword {
		validate.Password = "Password tidak sama"
		isValid = false
	}

	return &validate, isValid
}
