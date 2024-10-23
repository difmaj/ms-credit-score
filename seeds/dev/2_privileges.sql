INSERT IGNORE INTO 
    privileges (context, action, name, description) 
VALUES 
    ('list', 'assets', 'List assets', 'List assets'),
    ('read', 'assets', 'Read asset', 'Read asset'),
    ('create', 'assets', 'Create assets', 'Create assets'),
    ('update', 'assets', 'Update assets', 'Update assets'),
    ('delete', 'assets', 'Delete assets', 'Delete assets'),
    ('list', 'debts', 'List debts', 'List debts'),
    ('read', 'debt', 'Read debt', 'Read debt'),
    ('create', 'debts', 'Create debts', 'Create debts'),
    ('update', 'debts', 'Update debts', 'Update debts'),
    ('delete', 'debts', 'Delete debts', 'Delete debts'),
    ('read', 'scores', 'Read scores', 'Read scores');