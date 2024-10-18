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
    r.name = 'admin'
    AND (p.context, p.action) IN (
        ('create', 'assets'),
        ('read', 'assets'),
        ('update', 'assets'),
        ('delete', 'assets'),
        ('read', 'scores')
    );

