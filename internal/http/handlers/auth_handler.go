package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/levor/to-do/internal/config"
	"github.com/levor/to-do/internal/models"
	"github.com/levor/to-do/internal/repositories"
	"gopkg.in/errgo.v2/errors"
	"net/http"
	"time"
)

type AuthHandler struct {
	cfg *config.Config
	ur  *repositories.UserRepository
}

func NewAuthHandler(cfg *config.Config, ur *repositories.UserRepository) *AuthHandler {
	return &AuthHandler{cfg: cfg, ur: ur}
}

var sessions = map[string]session{}

type session struct {
	login  string    `json:"login"`
	expiry time.Time `json:"expiry"`
}

func (s session) IsExpired() bool {
	return s.expiry.Before(time.Now())
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var f struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	err := c.BindJSON(&f)
	if err != nil {
		HandleError(err, c)
	}
	user, err := ah.ur.FindByLogin(f.Login)
	if err != nil || user.Password != f.Password {
		err := errors.New("Wrong login or password")
		c.JSON(http.StatusUnauthorized, "")
		HandleError(err, c)
	} else {
		sessionToken := uuid.NewString()
		expiresAt := time.Now().Add(5 * time.Minute)
		sessions[sessionToken] = session{
			login:  user.Login,
			expiry: expiresAt,
		}
		http.SetCookie(c.Writer, &http.Cookie{
			Name:    "session_token",
			Value:   sessionToken,
			Expires: expiresAt,
		})
		c.JSON(http.StatusOK, gin.H{
			"status": fmt.Sprintf("Welcome %s!", user.Login),
		})
	}
}

func (ah *AuthHandler) Refresh(c *gin.Context) {
	t, err := c.Cookie("session_token")
	if err != nil {
		HandleError(err, c)
		return
	}

	userSession, exists := sessions[t]
	if !exists {
		err := errors.New("Unauthorized")
		HandleError(err, c)
		return
	}
	if userSession.IsExpired() {
		delete(sessions, t)
		err := errors.New("Session time expired")
		HandleError(err, c)
		return
	}

	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(5 * time.Minute)
	sessions[newSessionToken] = session{userSession.login, expiresAt}
	delete(sessions, t)

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: expiresAt,
	})

	c.JSON(http.StatusOK, gin.H{
		"status": "Token refreshed successful",
	})
}

func (ah *AuthHandler) Logout(c *gin.Context) {
	t, err := c.Cookie("session_token")
	if err != nil {
		HandleError(err, c)
		return
	}

	userSession, exists := sessions[t]
	if !exists {
		err := errors.New("Unauthorized")
		HandleError(err, c)
		return
	}
	if userSession.IsExpired() {
		delete(sessions, t)
		err := errors.New("Session time expired")
		HandleError(err, c)
		return
	}
	delete(sessions, t)
	http.SetCookie(c.Writer, &http.Cookie{Name: "session_token", Value: "", Expires: time.Now()})
	c.JSON(http.StatusOK, gin.H{
		"status": "Logout successful",
	})
}

func (ah *AuthHandler) Signup(c *gin.Context) {
	f := new(models.User)
	err := c.BindJSON(&f)
	if err != nil {
		HandleError(err, c)
	}
	u, _ := ah.ur.FindByLogin(f.Login)
	if u.Login != f.Login {
		err = ah.ur.Create(f)
		if err != nil {
			HandleError(err, c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "User created successful",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "Login already exist",
	})
}
