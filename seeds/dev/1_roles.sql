INSERT IGNORE INTO
    roles (id, name, description) 
VALUES 
    (UUID_TO_BIN('52dee553-8d4e-11ef-8e36-020017084e32'), 'admin', 'An admin can do anything.'),
    (UUID_TO_BIN('4aa9a814-8d4f-11ef-8e36-020017084e32'), 'user', 'A user can do anything except manage users.');
