INSERT IGNORE INTO users (name, email, password_hash, role_id) VALUES 
    ('Admin User', 'admin@example.com', '$2a$10$yGb6qFYZwj.gCth4PY8n0.wTiydEYhbuMuBUajtR5dp6DYHxxMZCq', (SELECT id FROM roles WHERE name = 'admin')),
    ('User', 'user@example.com', '$2a$10$yGb6qFYZwj.gCth4PY8n0.wTiydEYhbuMuBUajtR5dp6DYHxxMZCq', (SELECT id FROM roles WHERE name = 'user')),
    ('User 2', 'user2@example.com', '$2a$10$yGb6qFYZwj.gCth4PY8n0.wTiydEYhbuMuBUajtR5dp6DYHxxMZCq', (SELECT id FROM roles WHERE name = 'user'));