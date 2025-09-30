package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/suraj-9849/Golang/internal/types"
	"github.com/suraj-9849/Golang/internal/utils"
	"github.com/suraj-9849/Golang/prisma/db"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	client *db.PrismaClient
}

func NewHandler(client *db.PrismaClient) *Handler {
	return &Handler{
		client: client,
	}
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *Handler) Home(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Golang Jindabad!")
}

func (h *Handler) Signup(ctx *gin.Context) {
	var userRequest types.UserRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		validateErrors := err.(validator.ValidationErrors)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": validateErrors.Error(),
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Hash Password!",
		})
		return
	}
	existingUser, err := h.client.User.FindFirst(
		db.User.Email.Equals(userRequest.Email),
	).Exec(context.Background())

	if err == nil && existingUser != nil {
		slog.Info("Email is Already Registered!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is already registered!",
		})
		return
	}
	user, err := h.client.User.CreateOne(
		db.User.Username.Set(userRequest.Username),
		db.User.Email.Set(userRequest.Email),
		db.User.Password.Set(string(hash)),
	).Exec(context.Background())

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user: " + err.Error(),
		})
		return
	}
	token, err := utils.GenerateJWT(user.ID, user.Username, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"token":   token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (h *Handler) Login(ctx *gin.Context) {
	var loginRequest types.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		validateErrors := err.(validator.ValidationErrors)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": validateErrors.Error(),
		})
		return
	}

	user, err := h.client.User.FindFirst(
		db.User.Email.Equals(loginRequest.Email),
	).Exec(context.Background())

	if err != nil || user == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong Password",
		})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Username, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func (h *Handler) GetProfile(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	username := ctx.GetString("username")
	email := ctx.GetString("email")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Protected route accessed successfully",
		"user": gin.H{
			"id":       userID,
			"username": username,
			"email":    email,
		},
	})
}

func (h *Handler) GetDashboard(ctx *gin.Context) {
	username := ctx.GetString("username")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welcome to your dashboard, " + username + "!",
		"data":    "This is protected data only available to authenticated users",
	})
}
