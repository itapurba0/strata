INSERT INTO roles (name, description)
VALUES
    ('Organization Admin', 'Full administrative access within an organization'),
    ('HR Manager', 'Manage recruitment and hiring operations'),
    ('Recruiter', 'Manage recruiting and candidate pipelines'),
    ('Interviewer', 'Conduct and evaluate interviews')
ON CONFLICT DO NOTHING;