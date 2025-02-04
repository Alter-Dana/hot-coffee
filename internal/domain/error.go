package domain

import "errors"

var (
	//General
	ErrExistID = errors.New("such ID exists")

	// Inventory
	ErrInvalidInventory = errors.New("invalid inventory")
)
