package handler

import (
	"database/sql"
	"errors"

	"project/internal/models"
	"project/internal/service"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service  *service.UserService
	validate *validator.Validate
}

func NewUserHandler(service *service.UserService) *UserHandler {
	v := validator.New()
	// Register custom validation for DOB
	v.RegisterValidation("dob_past", func(fl validator.FieldLevel) bool {
		dateStr := fl.Field().String()
		// Check against today (midnight) in the server's timezone to avoid UTC vs Local mismatch
		now := time.Now()
		dob, err := time.ParseInLocation("2006-01-02", dateStr, now.Location())
		if err != nil {
			return false
		}

		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

		return !dob.After(today)
	})

	return &UserHandler{
		service:  service,
		validate: v,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validate.Struct(req); err != nil {
		return h.handleValidationError(c, err)
	}

	resp, err := h.service.CreateUser(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	resp, err := h.service.GetUserByID(c.Context(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(resp)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	resp, err := h.service.ListUsers(c.Context(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validate.Struct(req); err != nil {
		return h.handleValidationError(c, err)
	}

	resp, err := h.service.UpdateUser(c.Context(), int32(id), req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(resp)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = h.service.DeleteUser(c.Context(), int32(id))
	if err != nil {
		// Delete usually returns 204 even if not found, but if we wanted 404 we'd check rows affected.
		// Requirement doesn't specify 404 for Delete, only for Update/Get.
		// Standard simple Delete is idempotent-ish.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) handleValidationError(c *fiber.Ctx, err error) error {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, e := range validationErrors {
			if e.Tag() == "dob_past" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "dob must be in the past"})
			}
		}
	}
	// Fallback for other validation errors
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
}
