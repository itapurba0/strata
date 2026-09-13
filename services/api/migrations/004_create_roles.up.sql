CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    organization_id UUID
        REFERENCES organizations(id),

    name TEXT NOT NULL,

    description TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX roles_system_name_unique
    ON roles (name)
    WHERE organization_id IS NULL;

CREATE UNIQUE INDEX roles_org_name_unique
    ON roles (organization_id, name)
    WHERE organization_id IS NOT NULL;