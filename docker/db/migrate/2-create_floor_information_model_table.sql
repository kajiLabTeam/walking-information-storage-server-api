-- noqa: disable=all
\c indoor_location_estimation;
-- noqa: disable=all

CREATE TABLE coordinates (
    id SERIAL
    , x INTEGER NOT NULL
    , y INTEGER NOT NULL
    , is_walkable BOOLEAN NOT NULL
    , created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    , updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    , deleted_at TIMESTAMP WITH TIME ZONE
    , floor_id VARCHAR(26) REFERENCES floors (id)
    , PRIMARY KEY (id)
);

CREATE TABLE coordinate_information (
    id VARCHAR(26)
    , created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    , updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    , deleted_at TIMESTAMP WITH TIME ZONE
    , coordinate_id SERIAL REFERENCES coordinates (id)
    , PRIMARY KEY (id)
);
