package delivery

import (
	"errors"
	"log"
	"skeprogz/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type SepedaHandler struct {
	SepedaUseCase domain.SepedaUseCase
}

func NewSepedaHandler(app fiber.Router, uc domain.SepedaUseCase) {
	handler := &SepedaHandler{
		SepedaUseCase: uc,
	}

	app.Post("/sepeda/add", handler.CreateSepeda)
	app.Get("/sepeda/:id", handler.GetSepedaByID)
	app.Put("/sepeda/:id", handler.UpdateSepeda)
	app.Delete("/sepeda/:id", handler.DeleteSepeda)
	app.Get("/sepeda", handler.ListSepeda)
}

func (h *SepedaHandler) CreateSepeda(c *fiber.Ctx) error {
	var sepeda domain.Sepeda
	if err := c.BodyParser(&sepeda); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"status":  false,
			"error":   err.Error(),
		})
	}

	// Validate data types
	if err := validateSepeda(sepeda); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data type",
			"status":  false,
			"error":   err.Error(),
		})
	}

	if err := h.SepedaUseCase.CreateUC(&sepeda); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create sepeda",
			"status":  false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success",
		"data":    sepeda,
		"status":  true,
	})
}

func (h *SepedaHandler) GetSepedaByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
			"status":  false,
			"error":   err.Error(),
		})
	}

	sepeda, err := h.SepedaUseCase.GetByIDUC(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get sepeda by ID",
			"status":  false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    sepeda,
		"status":  true,
	})
}

func (h *SepedaHandler) UpdateSepeda(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
			"status":  false,
			"error":   err.Error(),
		})
	}

	var sepeda domain.Sepeda
	if err := c.BodyParser(&sepeda); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
			"status":  false,
			"error":   err.Error(),
		})
	}

	// Validate data types
	if err := validateSepeda(sepeda); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data type",
			"status":  false,
			"error":   err.Error(),
		})
	}

	sepeda.ID = uint(id)
	if err := h.SepedaUseCase.UpdateUC(&sepeda); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update sepeda",
			"status":  false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    sepeda,
		"status":  true,
	})
}

func (h *SepedaHandler) DeleteSepeda(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
			"status":  false,
			"error":   err.Error(),
		})
	}

	if err := h.SepedaUseCase.DeleteUC(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete sepeda",
			"status":  false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"status":  true,
	})
}

func (h *SepedaHandler) ListSepeda(c *fiber.Ctx) error {
	sepedaList, err := h.SepedaUseCase.GetAllUC()
	if err != nil {
		log.Printf("Error in GetAllUC: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to list sepeda",
			"status":  false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    sepedaList,
		"status":  true,
	})
}

// Validation function to ensure data types match the struct
func validateSepeda(sepeda domain.Sepeda) error {
	if sepeda.Brand == "" {
		return errors.New("Brand is required")
	}
	if sepeda.Size <= 0 {
		return errors.New("Size must be a positive integer")
	}
	if sepeda.Type == "" {
		return errors.New("Type is required")
	}
	if sepeda.Quantity <= 0 {
		return errors.New("Quantity must be a positive integer")
	}
	return nil
}
