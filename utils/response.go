package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status       string      `json:"status"`
	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Error        string      `json:"error,omitempty"`
}

type ResponsePagination struct {
	Status       string         `json:"status"`
	ResponseCode int            `json:"response_code"`
	Message      string         `json:"message,omitempty"`
	Data         interface{}    `json:"data,omitempty"`
	Error        string         `json:"error,omitempty"`
	Meta         PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	Page      int    `json:"page" example:"1"`
	Limit     int    `json:"limit" example:"10"`
	Total     int    `json:"total" example:"100"`
	TotalPage int    `json:"total_pagess" example:"10"`
	Filter    string `json:"filter" example:"nama=umam"`
	Sort      string `json:"sort" example:"-id"`
}

func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:       "Success",
		ResponseCode: fiber.StatusOK,
		Message:      message,
		Data:         data,
	})
}

func SuccessPagination(c *fiber.Ctx, message string, data interface{}, meta PaginationMeta) error {
	return c.Status(fiber.StatusOK).JSON(ResponsePagination{
		Status:       "Success",
		ResponseCode: fiber.StatusOK,
		Message:      message,
		Data:         data,
		Meta:         meta,
	})
}

func NotFoundPagination(c *fiber.Ctx, message string, data interface{}, meta PaginationMeta) error {
	return c.Status(fiber.StatusNotFound).JSON(ResponsePagination{
		Status:       "Not Found",
		ResponseCode: fiber.StatusNotFound,
		Message:      message,
		Data:         data,
		Meta:         meta,
	})
}

func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Status:       "Created",
		ResponseCode: fiber.StatusCreated,
		Message:      message,
		Data:         data,
	})
}

func BadRequest(c *fiber.Ctx, message string, err string) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Status:       "Error Bad Request",
		ResponseCode: fiber.StatusBadRequest,
		Message:      message,
		Error:        err,
	})
}

func NotFound(c *fiber.Ctx, message string, err string) error {
	return c.Status(fiber.StatusNotFound).JSON(Response{
		Status:       "Error Not Found",
		ResponseCode: fiber.StatusNotFound,
		Message:      message,
		Error:        err,
	})
}

func Unauthorized(c *fiber.Ctx, message string, err string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Response{
		Status:       "Unauthorized",
		ResponseCode: fiber.StatusUnauthorized,
		Message:      message,
		Error:        err,
	})
}
