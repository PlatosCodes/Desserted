-- Seed file for Desserted game

-- Inserting ingredient cards
INSERT INTO cards (type, name) VALUES 
('ingredient', 'Flour'),
('ingredient', 'Flour'),
('ingredient', 'Flour'),
('ingredient', 'Flour'),
('ingredient', 'Flour'),
('ingredient', 'Flour'),
('ingredient', 'Sugar'),
('ingredient', 'Sugar'),
('ingredient', 'Sugar'),
('ingredient', 'Sugar'),
('ingredient', 'Sugar'),
('ingredient', 'Eggs'),
('ingredient', 'Eggs'),
('ingredient', 'Eggs'),
('ingredient', 'Eggs'),
('ingredient', 'Butter'),
('ingredient', 'Butter'),
('ingredient', 'Butter'),
('ingredient', 'Cocoa'),
('ingredient', 'Vanilla'),
('ingredient', 'Berries'),
('ingredient', 'Cream Cheese'),
('ingredient', 'Cream Cheese'),
('ingredient', 'Saffron'),
('ingredient', 'Honey'),
('ingredient', 'Dark Chocolate'),
('ingredient', 'Dark Chocolate'),
('ingredient', 'Matcha Powder'),
('ingredient', 'Edible Gold Leaf');

-- Inserting dessert cards
INSERT INTO desserts (name, points) VALUES 
('Cake', 10),
('Pie', 15),
('Chocolate Chip Cookies', 20),
('Cheesecake', 25),
('Tiramisu', 30),
('Matcha Cake', 35),
('Saffron Panna Cotta', 40),
('Gourmet Truffles', 45),
('Gold Leaf Cupcakes', 50);

-- Inserting special cards
INSERT INTO cards (type, name) VALUES 
('special', 'Wildcard Ingredient'),
('special', 'Wildcard Ingredient'),
('special', 'Wildcard Ingredient'),
('special', 'Wildcard Ingredient'),
('special', 'Steal Card'),
('special', 'Steal Card'),
('special', 'Steal Card'),
('special', 'Double Points'),
('special', 'Double Points'),
('special', 'Refresh Hand'),
('special', 'Refresh Hand'),
('special', 'Instant Bake'),
('special', 'Instant Bake'),
('special', 'Mystery Flavor'),
('special', 'Mystery Flavor'),
('special', 'Sabotage'),
('special', 'Sabotage'),
('special', 'Glass of Milk'),
('special', 'Glass of Milk'),
('special', 'Glass of Milk'),
('special', 'Glass of Milk');
