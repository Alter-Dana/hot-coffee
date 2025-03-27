CREATE TABLE orders (
    order_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    status ORDER_STATUS_ENUM NOT NULL,
    created TIMESTAMPTZ DEFAULT NOW(),
    modified TIMESTAMPTZ DEFAULT NOW(),
    customer_id UUID REFERENCES customers(customer_id),
    preferences JSONB,
    total_amount NUMERIC(10,2) NOT NULL
);

CREATE TABLE order_items (
    order_id UUID REFERENCES orders(order_id) ON DELETE CASCADE,
    menu_item_id INT REFERENCES menu_items(menu_item_id) ON DELETE CASCADE,
    quantity INT CHECK (quantity > 0) NOT NULL,
    price_at_time_of_order NUMERIC(10,2) NOT NULL,
    customization JSONB, -- Stores optional customization details
    PRIMARY KEY (order_id, menu_item_id)
);

CREATE TABLE menu_items (
    menu_item_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    category TEXT NOT NULL,
    allergens TEXT[],
    price NUMERIC(10,2) NOT NULL,
    available BOOLEAN DEFAULT TRUE
);

CREATE TABLE menu_item_ingredients (
    menu_item_id INT REFERENCES menu_items(menu_item_id) ON DELETE CASCADE,
    ingredient_id INT REFERENCES inventory(ingredient_id) ON DELETE CASCADE,
    quantity NUMERIC NOT NULL, -- Allows decimals
    PRIMARY KEY (menu_item_id, ingredient_id)
);

CREATE TABLE inventory (
    inventory_id SERIAL PRIMARY KEY,
    ingredient_id INT UNIQUE NOT NULL, -- Ensures no duplicate ingredients
    name VARCHAR(255) NOT NULL,
    stock NUMERIC NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    unit_type VARCHAR(255) NOT NULL,
    last_updated TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE order_status_history (
    order_id UUID REFERENCES orders(order_id) ON DELETE CASCADE,
    previous_status VARCHAR(255) NOT NULL,
    last_status VARCHAR(255) NOT NULL,
    modified_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (order_id, modified_at)
);

CREATE TABLE price_history (
    menu_item_id INT REFERENCES menu_items(menu_item_id) ON DELETE CASCADE,
    old_price NUMERIC(10,2) NOT NULL,
    last_price NUMERIC(10,2) NOT NULL,
    modified_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (menu_item_id, modified_at)
);

CREATE TABLE inventory_transactions (
    inventory_id INT REFERENCES inventory(inventory_id) ON DELETE CASCADE,
    change_amount NUMERIC(10,2) NOT NULL,
    transaction_type VARCHAR(255) NOT NULL,
    modified_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (inventory_id, modified_at)
);

CREATE TABLE customers (
    customer_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    preferences JSONB DEFAULT '{}' -- Stores customer preferences in structured JSON format
);
