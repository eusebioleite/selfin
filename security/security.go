package security

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"crypto/rand"
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/eusebioleite/selfin/database"
	"github.com/eusebioleite/selfin/models"
	repo "github.com/eusebioleite/selfin/repository"
	"golang.org/x/crypto/argon2"
)

type Session = models.Session
type UserSession = models.UserSession

func Auth(login string, password string) (int64, error) {
	if login == "" || password == "" {
		return 0, fmt.Errorf("Error autheticating user -> login or password is empty!")
	}

	user, err := repo.GetUserByLogin(login)
	if err != nil {
		return 0, fmt.Errorf("User not found -> %w", err)
	}

	err = comparePassword(user.Password, password)
	if err != nil {
		return 0, fmt.Errorf("Error comparing password hash -> %w", err)
	}

	return user.ID, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		err := validateSession(c)
		if err != nil {
			c.Error(fmt.Errorf("Error validating session -> %w", err))
			c.Redirect(http.StatusFound, "/auth")
			c.Abort()
			return
		}

		c.Next()
	}
}

func AuthHandler(c *gin.Context) {
	var req struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	userID, err := Auth(req.Login, req.Password)
	if err != nil {
		c.String(http.StatusUnauthorized, fmt.Sprintf("invalid credentials -> %s", err))
		return
	}

	establishSession(c, int(userID))

	c.Header("HX-Redirect", "/dashboards")
	c.Status(http.StatusOK)
}

func establishSession(c *gin.Context, userID int) {
	userSession := UserSession{
		SessionID: uuid.NewString(),
		UserID:    int64(userID),
	}

	err := newSession(userSession)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error creating session on database -> %w", err))
	}

	c.SetCookie("session_id", userSession.SessionID, 86400, "/", "localhost", false, true)

}

func validateSession(c *gin.Context) error {

	id, err := c.Cookie("session_id")
	if err != nil {
		return fmt.Errorf("Error getting cookie -> %w", err)
	}

	err = sessionExists(id)
	if err != nil {
		return fmt.Errorf("Error checking if session exists -> %w", err)
	}

	return nil

}

func sessionExists(id string) error {
	exists := 0

	query := `
		SELECT 1
		FROM sessions
		WHERE id = ?
		AND expires_at > DATETIME('NOW', 'LOCALTIME')
		`

	err := database.DB.QueryRow(query, id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Session not found.")
		} else {
			return fmt.Errorf("Error checking session -> %w", err)
		}
	}

	return nil

}

func getSessionByLogin(login string) (UserSession, error) {
	var userSession UserSession

	query := `
		SELECT
			user.id,
			user.login,
			user.password,
			session.id,
			session.expires_at,
			session.created_at
		FROM users
		JOIN sessions ON sessions.user_id = user.id
		WHERE login = ?
		ORDER BY id DESC
		`
	err := database.DB.QueryRow(query, login).
		Scan(
			&userSession.UserID,
			&userSession.Login,
			&userSession.Password,
			&userSession.SessionID,
			&userSession.ExpiresAt,
			&userSession.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return userSession, errors.New("Session not found.")
		} else {
			return userSession, fmt.Errorf("Error checking session -> %w", err)
		}
	}

	return userSession, nil
}

func newSession(userSession UserSession) error {

	query := `
	INSERT INTO sessions (id, user_id)
	VALUES (?, ?)
	`
	_, err := database.DB.Exec(query, userSession.SessionID, userSession.UserID)
	if err != nil {
		return fmt.Errorf("Error creating session -> %w", err)
	}

	return nil
}

const (
	memory      = 64 * 1024
	iterations  = 3
	parallelism = 2
	saltLength  = 16
	keyLength   = 32
)

func HashPassword(password string) (string, error) {

	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return b64Salt + "." + b64Hash, nil
}

func comparePassword(storedHash, password string) error {

	parts := strings.Split(storedHash, ".")
	if len(parts) != 2 {
		return errors.New("invalid hash format.")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return fmt.Errorf("Error decoding salt -> %w", err)
	}
	hash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return fmt.Errorf("Error decoding hash -> %w", err)
	}

	newHash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	if subtle.ConstantTimeCompare(hash, newHash) == 1 {
		return nil
	}

	return errors.New("hashes don't match.")
}
