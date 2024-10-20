\c indoor_location_estimation;

CREATE TABLE particles (
    id VARCHAR(26) PRIMARY KEY,
    x DECIMAL NOT NULL,
    y DECIMAL NOT NULL,
    weight DECIMAL NOT NULL,
    direction DECIMAL NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    walking_sample_id VARCHAR(26) REFERENCES walking_samples(id)
);

CREATE TABLE estimated_positions (
    id VARCHAR(26) PRIMARY KEY,
    x DECIMAL NOT NULL,
    y DECIMAL NOT NULL,
    direction DECIMAL NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    walking_sample_id VARCHAR(26) UNIQUE REFERENCES walking_samples(id)
);
