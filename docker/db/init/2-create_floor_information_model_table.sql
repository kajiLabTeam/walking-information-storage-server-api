\c indoor_location_estimation;

CREATE TABLE floor_maps (
    id VARCHAR(26),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    floor_information_id VARCHAR(26) UNIQUE REFERENCES floor_information(id),
    PRIMARY KEY (id, floor_information_id)
);

CREATE TABLE coordinates (
    id VARCHAR(26),
    x DECIMAL NOT NULL,
    y DECIMAL NOT NULL,
    is_walkable BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    floor_id VARCHAR(26) REFERENCES floors(id),
    PRIMARY KEY (id)
);

CREATE TABLE geomagnetic_fingerprinting (
    id VARCHAR(26),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    coordinate_id VARCHAR(26) REFERENCES coordinates(id),
    PRIMARY KEY (id)
);

CREATE TABLE wifi_fingerprinting (
    id VARCHAR(26),
    mac_address VARCHAR(17),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    coordinate_id VARCHAR(26) REFERENCES coordinates(id),
    PRIMARY KEY (id)
);
