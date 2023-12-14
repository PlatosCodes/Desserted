-- Seed file for Desserted game

-- Inserting ingredient cards
INSERT INTO cards (type, name) VALUES 
('ingredient', 'Flour'),
('ingredient', 'Sugar'),
('ingredient', 'Eggs'),
('ingredient', 'Butter'),
('ingredient', 'Milk'),
('ingredient', 'Chocolate'),
('ingredient', 'Vanilla'),
('ingredient', 'Berries'),
('ingredient', 'Nuts'),
('ingredient', 'Cream Cheese'),
('ingredient', 'Saffron'),
('ingredient', 'Honey'),
('ingredient', 'Dark Chocolate'),
('ingredient', 'Matcha Powder'),
('ingredient', 'Edible Gold Leaf');

-- Inserting dessert cards
INSERT INTO desserts (type, name, points) VALUES 
('dessert', 'Cake', 10),
('dessert', 'Pie', 15),
('dessert', 'Chocolate Chip Cookies', 20),
('dessert', 'Cheesecake', 25),
('dessert', 'Tiramisu', 30),
('dessert', 'Matcha Cake', 35),
('dessert', 'Saffron Panna Cotta', 40),
('dessert', 'Gourmet Truffles', 45),
('dessert', 'Gold Leaf Cupcakes', 50);

-- Inserting special cards
INSERT INTO cards (type, name) VALUES 
('special', 'Wildcard Ingredient'),
('special', 'Steal Card'),
('special', 'Double Points'),
('special', 'Refresh Hand'),
('special', 'Instant Bake'),
('special', 'Mystery Flavor'),
('special', 'Sabotage');
