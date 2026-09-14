INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
CROSS JOIN permissions p
WHERE r.name = 'Organization Admin'
  AND r.organization_id IS NULL
  AND p.name IN (
      'organization:read',
      'organization:update',
      'member:read',
      'member:manage',
      'role:read',
      'role:manage',
      'job:create',
      'job:read',
      'job:update',
      'job:delete',
      'job:publish',
      'application:read',
      'application:update',
      'candidate:read'
  )
ON CONFLICT DO NOTHING;