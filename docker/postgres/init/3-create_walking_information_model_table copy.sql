\c indoor_location_estimation;

CREATE TABLE gyroscopes (
    id VARCHAR(26) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    walking_information_id VARCHAR(26) UNIQUE REFERENCES walking_information(id)
);

CREATE TABLE accelerometers (
    id VARCHAR(26) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    walking_information_id VARCHAR(26) UNIQUE REFERENCES walking_information(id)
);

CREATE TABLE atmospheric_pressures (
    id VARCHAR(26) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    walking_information_id VARCHAR(26) UNIQUE REFERENCES walking_information(id)
);

CREATE TABLE ratio_waves (
    id VARCHAR(26) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    walking_information_id VARCHAR(26) UNIQUE REFERENCES walking_information(id)
);

CREATE TABLE gps (
    id VARCHAR(26) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    walking_information_id VARCHAR(26) UNIQUE REFERENCES walking_information(id)
);