CREATE TABLE membership_roles (
    membership_id UUID NOT NULL
        REFERENCES organization_memberships(id),

    role_id UUID NOT NULL
        REFERENCES roles(id),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (membership_id, role_id)
);

CREATE INDEX idx_membership_roles_role_id
    ON membership_roles (role_id);