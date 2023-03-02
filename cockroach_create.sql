CREATE TABLE todos(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    description VARCHAR(100),
    priority INT,
    status VARCHAR(100)
);