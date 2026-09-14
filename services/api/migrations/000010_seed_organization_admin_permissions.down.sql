DELETE FROM role_permissions
WHERE role_id = (
    SELECT id
    FROM roles
    WHERE name = 'Organization Admin'
      AND organization_id IS NULL
);