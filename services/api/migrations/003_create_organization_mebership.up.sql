CREATE TYPE organization_membership_status AS ENUM (
    'pending',
    'active',
    'revoked'
);

CREATE TABLE organization_memberships (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    organization_id UUID NOT NULL
        REFERENCES organizations(id),

    user_id UUID NOT NULL
        REFERENCES users(id),

    status organization_membership_status NOT NULL DEFAULT 'pending',

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT organization_memberships_user_org_unique
        UNIQUE (user_id, organization_id)
);

CREATE INDEX idx_organization_memberships_organization_id
    ON organization_memberships (organization_id);

CREATE INDEX idx_organization_memberships_user_id
    ON organization_memberships (user_id);