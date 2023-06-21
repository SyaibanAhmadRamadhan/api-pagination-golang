package http

import (
	"context"
	"strconv"
	"time"

	"github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/internal/http-protocol/exception"
	"github.com/SyaibanAhmadRamadhan/technical-test-pt-zahir-international/src/modules/profile/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (h *HttpHandlerImpl) InsertProfile(c *fiber.Ctx) error {
	payload := new(dto.CreateProfileRequest)

	if err := c.BodyParser(payload); err != nil {
		log.Warn().Msgf("cannot parser : %v", err)
		return exception.Err(c, err)
	}

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	profile, err := h.ProfileService.Insert(ctx, payload)
	if err != nil {
		return exception.Err(c, err)
	}

	return c.Status(201).JSON(map[string]dto.ProfileResponse{
		"data": *profile,
	})
}

func (h *HttpHandlerImpl) UpdateProfile(c *fiber.Ctx) error {
	payload := new(dto.UpdateProfileRequest)

	if err := c.BodyParser(payload); err != nil {
		log.Warn().Msgf("cannot parser : %v", err)
		return exception.Err(c, err)
	}

	if c.Params("id", "") == "" {
		return exception.Err(c, exception.NotFound("NOT FOUND"))
	}
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	profile, err := h.ProfileService.Update(ctx, c.Params("id"), payload)
	if err != nil {
		return exception.Err(c, err)
	}

	return c.Status(201).JSON(map[string]dto.ProfileResponse{
		"data": *profile,
	})
}

func (h *HttpHandlerImpl) FindByID(c *fiber.Ctx) error {
	if c.Params("id", "") == "" {
		return exception.Err(c, exception.NotFound("NOT FOUND"))
	}
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	profile, err := h.ProfileService.FindByID(ctx, c.Params("id"))
	if err != nil {
		return exception.Err(c, err)
	}

	return c.Status(201).JSON(map[string]dto.ProfileResponse{
		"data": *profile,
	})
}

func (h *HttpHandlerImpl) FindAll(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	// if c.Get("page") == "" {
	// }
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return exception.Err(c, exception.BadRequest(map[string]map[string]string{
			"page": {
				"invalid_page": "page invalid",
			},
		}))
	}
	profile, pagination, err := h.ProfileService.FindAll(ctx, page, c.Query("sort_by", "DESC"), c.Query("date", ""))
	if err != nil {
		return exception.Err(c, err)
	}

	return c.Status(201).JSON(map[string]any{
		"data":       profile,
		"pagination": pagination,
	})
}

func (h *HttpHandlerImpl) DeleteById(c *fiber.Ctx) error {
	if c.Params("id", "") == "" {
		return exception.Err(c, exception.NotFound("NOT FOUND"))
	}
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	err := h.ProfileService.Delete(ctx, c.Params("id"))
	if err != nil {
		return exception.Err(c, err)
	}

	return c.SendStatus(200)
}
