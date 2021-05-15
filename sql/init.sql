-- Create Table system information
CREATE TABLE IF NOT EXISTS system (
    machineid VARCHAR UNIQUE PRIMARY KEY,
    attrs JSONB
);

CREATE TABLE IF NOT EXISTS packages (
    machineid VARCHAR UNIQUE PRIMARY KEY,
    attrs JSONB
);

CREATE TABLE IF NOT EXISTS systemuser (
    machineid VARCHAR UNIQUE PRIMARY KEY,
    attrs JSONB
);


CREATE INDEX IF NOT EXISTS idx_hostname_attrs ON system USING gin ((attrs->'node'->'hostname'));



-- ID serial PRIMARY KEY,