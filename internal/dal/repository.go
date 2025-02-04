package dal

type Repository struct {
	Directory     string
	FileOrder     string
	FileMenu      string
	FileInventory string
}

func NewRepository(directory string) *Repository {
	return &Repository{
		Directory:     directory,
		FileOrder:     "orders.json",
		FileMenu:      "menu_items.json",
		FileInventory: "inventory.json",
	}
}
