package auth

import (
	"context"
	"os"
	"time"

	"nerdoverapi/db"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/api/idtoken"
)

const userCollectionName = "user"
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func VerifyGoogleToken(ctx context.Context, idToken string) (string, error) {
	payload, err := idtoken.Validate(ctx, idToken, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		return "", err
	}

	email, ok := payload.Claims["email"].(string)
	if !ok {
		return "", err
	}

	user, err := FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func FindByEmail(ctx context.Context, email string) (User, error) {
	doc, err := db.Client.Collection(userCollectionName).Where("email", "==", email).Documents(ctx).Next()
	if err != nil {
		return User{}, err
	}

	var user User
	if err := doc.DataTo(&user); err != nil {
		return User{}, err
	}
	return user, nil
}
