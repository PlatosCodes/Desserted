-- Seed file for Desserted game

-- Inserting ingredient cards
INSERT INTO cards (type, name, points) VALUES 
('ingredient', 'Flour', 1),
('ingredient', 'Sugar', 1),
('ingredient', 'Eggs', 1),
('ingredient', 'Butter', 1),
('ingredient', 'Milk', 1),
('ingredient', 'Chocolate', 3),
('ingredient', 'Vanilla', 3),
('ingredient', 'Berries', 3),
('ingredient', 'Nuts', 3),
('ingredient', 'Cream Cheese', 5),
('ingredient', 'Saffron', 5),
('ingredient', 'Honey', 5),
('ingredient', 'Dark Chocolate', 5),
('ingredient', 'Matcha Powder', 5),
('ingredient', 'Edible Gold Leaf', 5);

-- Inserting special cards
INSERT INTO cards (type, name, points) VALUES 
('special', 'Wildcard Ingredient', 0),
('special', 'Steal Card', 0),
('special', 'Double Points', 0),
('special', 'Refresh Hand', 0),
('special', 'Instant Bake', 0),
('special', 'Mystery Flavor', 0),
('special', 'Sabotage', 0);
