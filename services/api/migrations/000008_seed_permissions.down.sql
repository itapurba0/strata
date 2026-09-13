DELETE FROM permissions
WHERE name IN (
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

    'candidate:read',

    'interview:read',
    'interview:create',
    'interview:evaluate'
);