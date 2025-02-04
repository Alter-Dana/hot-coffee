package domain

import "errors"

var (
	// Basic
	ErrInternalServer = errors.New("internal server error")
	//General
	ErrExistID = errors.New("such ID exists")

	// Inventory
	ErrInvalidInventory = errors.New("invalid inventory")
)
