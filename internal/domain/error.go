package domain

import "errors"

var (
	// Basic
	ErrInternalServer = errors.New("internal server error")

	//General
	ErrExistID = errors.New("such ID is not found")

	// Inventory
	ErrInvalidInventory = errors.New("invalid inventory")
	ErrInvalidOrder     = errors.New("invalid order")
	ErrInvalidMenu      = errors.New("invalid menu")
)
