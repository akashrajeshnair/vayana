package dto

import (
	"time"
)

// Request DTOs

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateProfileRequest struct {
	Name     string `json:"name" binding:"omitempty,min=2,max=100"`
	Password string `json:"password" binding:"omitempty,min=8"`
}

// Response DTOs

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Metadata response
type MetadataResponse struct {
	ServiceName    string    `json:"service_name"`
	ServiceVersion string    `json:"service_version"`
	ServerTime     time.Time `json:"server_time"`
}

func NewMetadataResponse(serviceName, serviceVersion string) MetadataResponse {
	return MetadataResponse{
		ServiceName:    serviceName,
		ServiceVersion: serviceVersion,
		ServerTime:     time.Now().UTC(),
	}
}

// List response wrapper
type ListResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	TotalPages int         `json:"total_pages"`
}

func NewListResponse(data interface{}, total int64, page, perPage int) ListResponse {
	totalPages := int(total) / perPage
	if int(total)%perPage != 0 {
		totalPages++
	}

	return ListResponse{
		Data:       data,
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
	}
}
