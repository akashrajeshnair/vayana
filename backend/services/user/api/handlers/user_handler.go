package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/akashrajeshnair/vayana/pkg/auth"
	"github.com/akashrajeshnair/vayana/pkg/logger"
	"github.com/akashrajeshnair/vayana/services/user/api/dto"
	"github.com/akashrajeshnair/vayana/services/user/config"
	"github.com/akashrajeshnair/vayana/services/user/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userRepo   *repository.UserRepository
	jwtManager *auth.JWTManager
	logger     *logger.Logger
	config     *config.UserServiceConfig
}

func NewUserHandler(userRepo *repository.UserRepository, jwtManager *auth.JWTManager, logger *logger.Logger, config *config.UserServiceConfig) *UserHandler {
	return &UserHandler{
		userRepo:   userRepo,
		jwtManager: jwtManager,
		logger:     logger,
		config:     config,
	}
}

// Register handles user registration
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid register request", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request format",
		})
		return
	}

	// Check if email already exists
	exists, err := h.userRepo.EmailExists(req.Email)
	if err != nil {
		h.logger.Error("Failed to check email existence", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
		})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, dto.ErrorResponse{
			Code:    http.StatusConflict,
			Message: "Email already registered",
		})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		h.logger.Error("Failed to hash password", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to process registration",
		})
		return
	}

	// Create user
	user, err := h.userRepo.CreateUser(req.Email, string(hashedPassword), req.Name)
	if err != nil {
		h.logger.Error("Failed to create user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

// Login handles user authentication
func (h *UserHandler) LoginUser(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid login request", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request format",
		})
		return
	}

	user, err := h.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		h.logger.Error("User not found", zap.Error(err))
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Invalid credentials",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		h.logger.Error("Invalid password", zap.Error(err))
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Invalid credentials",
		})
		return
	}

	// Parse JWT duration from config
	durationHours, err := strconv.Atoi(h.config.JWTDuration)
	if err != nil {
		durationHours = 168 // Default to 7 days if parsing fails
	}

	duration := time.Duration(durationHours) * time.Hour

	token, err := h.jwtManager.GenerateToken(user.ID, "user", duration)
	if err != nil {
		h.logger.Error("Failed to generate token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

// GetProfile retrieves the user profile
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		h.logger.Error("User ID not found in context")
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	user, err := h.userRepo.GetUserByID(userID.(string))
	if err != nil {
		h.logger.Error("Failed to get user profile", zap.Error(err))
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

// UpdateProfile handles user profile updates
func (h *UserHandler) UpdateUserProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		h.logger.Error("User ID not found in context")
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid update request", zap.Error(err))
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request format",
		})
		return
	}

	user, err := h.userRepo.GetUserByID(userID.(string))
	if err != nil {
		h.logger.Error("Failed to get user", zap.Error(err))
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "User not found",
		})
		return
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			h.logger.Error("Failed to hash password", zap.Error(err))
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to update password",
			})
			return
		}
		user.Password = string(hashedPassword)
	}

	if err := h.userRepo.UpdateUser(user); err != nil {
		h.logger.Error("Failed to update user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update profile",
		})
		return
	}

	c.JSON(http.StatusOK, dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

// GetServiceMetadata returns service metadata
func (h *UserHandler) GetServiceMetadata(c *gin.Context) {
	c.JSON(http.StatusOK, dto.NewMetadataResponse(
		h.config.ServiceName,
		h.config.ServiceVersion,
	))
}
