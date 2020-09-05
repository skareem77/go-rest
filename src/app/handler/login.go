package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-api/src/app/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//Login user
func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(user)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	defer r.Body.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	user.Password = string(hash)
	if err = db.Save(user).Error; err != nil {
		RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	tk := &model.Token{UserID: user.Email}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte("SECRET"))
	fmt.Println(tokenString)
	ResponseJSON(w, http.StatusOK, map[string]string{"token": tokenString})
}
