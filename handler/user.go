package handler

import (
	"fundraising/helper"
	"fundraising/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	/* tangkap input dari user
	   map input dari user ke struct RegisterUserInput
	   struct diatas kita passing sebagai parameter service */
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("registrasi akun gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("registrasi akun gagal", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token, err := h.jwtService.GenerateToken()

	formatter := user.FormatUser(newUser, "tokentokentokentokentokentokentokentokentokentokentoken")

	response := helper.APIResponse("Akun berhasil didaftarkan", http.StatusOK, "succes", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	//  user memasukan input (email dan password)
	//  input ditangkap handler
	//  mapping dari input user ke input struct
	//  input struct passing service
	//  di service mencari dengan bantuan repository user dengan email x
	//  mencocokan password

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("login gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("login gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokenakehbdajbwdbebdfadjbdcandadcnkwjnjfbhr")

	response := helper.APIResponse("Successfully logged in", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
