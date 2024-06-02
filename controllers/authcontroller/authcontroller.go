package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kyyyyyyyyyyyyyy/golang-jwt-mux1/config"
	"github.com/kyyyyyyyyyyyyyy/golang-jwt-mux1/helpers"
	"github.com/kyyyyyyyyyyyyyy/golang-jwt-mux1/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {

	//mengambil inoutan json
	var userinput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userinput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	//ambil data user berdasarkan username
	var user models.User

	if err := models.DB.Where("username = ?", userinput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "username atau password salah"}
			helpers.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helpers.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	//cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userinput.Password)); err != nil {
		response := map[string]string{"message": "username atau password salah"}
		helpers.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	//pembuatan jwt
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux1",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	//mendeklarasikan algoritma yang akan digunakan untuk sign in
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//sign token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	//set token ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "login success"}
	helpers.ResponseJSON(w, http.StatusOK, response)

}
func Register(w http.ResponseWriter, r *http.Request) {

	//mengambil inoutan json
	var userinput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userinput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	//hash password with backryptc
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userinput.Password), bcrypt.DefaultCost)
	userinput.Password = string(hashPassword)

	//insert database
	if err := models.DB.Create(&userinput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helpers.ResponseJSON(w, http.StatusOK, response)

}
func Logout(w http.ResponseWriter, r *http.Request) {

	//hapus token di cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "logout success"}
	helpers.ResponseJSON(w, http.StatusOK, response)

}
