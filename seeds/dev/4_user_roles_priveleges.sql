-- Admin privileges
INSERT IGNORE INTO
    roles_privileges (role_id, privilege_id)
SELECT
    r.id AS role_id,
    p.id AS privilege_id
FROM
    roles r,
    privileges p
WHERE
    r.name = 'user'
    AND (p.context, p.action) IN (
        ('read', 'debts'),
        ('create', 'debts'),
        ('update', 'debts'),
        ('delete', 'debts'),
        ('read', 'scores')
    );