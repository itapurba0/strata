INSERT INTO permissions (name, description)
VALUES
    ('organization:read', 'View organization information'),
    ('organization:update', 'Update organization information'),

    ('member:read', 'View organization members'),
    ('member:manage', 'Manage organization members'),

    ('role:read', 'View roles'),
    ('role:manage', 'Create, update, and assign roles'),

    ('job:create', 'Create job postings'),
    ('job:read', 'View job postings'),
    ('job:update', 'Update job postings'),
    ('job:delete', 'Delete job postings'),
    ('job:publish', 'Publish job postings'),

    ('application:read', 'View job applications'),
    ('application:update', 'Update application status'),

    ('candidate:read', 'View candidate information'),

    ('interview:read', 'View interviews'),
    ('interview:create', 'Create interviews'),
    ('interview:evaluate', 'Submit interview evaluations')
ON CONFLICT (name) DO NOTHING;