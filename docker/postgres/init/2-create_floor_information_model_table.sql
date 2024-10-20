\c indoor_location_estimation;

CREATE TABLE floor_maps (
    id VARCHAR(26),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    floor_information_id VARCHAR(26) UNIQUE REFERENCES floor_information(id),
    PRIMARY KEY (id, floor_information_id)
);


CREATE TABLE fp_models (
    id VARCHAR(26),
    mac_address VARCHAR(17),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    floor_information_id VARCHAR(26) UNIQUE REFERENCES floor_information(id),
    PRIMARY KEY (id, floor_information_id)
);

