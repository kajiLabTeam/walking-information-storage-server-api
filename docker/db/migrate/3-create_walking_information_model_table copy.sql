-- noqa: disable=all
\c indoor_location_estimation;
-- noqa: disable=all

CREATE TABLE estimated_positions (
    id VARCHAR(26) PRIMARY KEY
    , x INTEGER NOT NULL
    , y INTEGER NOT NULL
    , is_converged BOOLEAN NOT NULL
    , direction INTEGER NOT NULL
    , created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    , updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    , deleted_at TIMESTAMP WITH TIME ZONE
    , trajectory_id VARCHAR(26) REFERENCES trajectories (id)
    , walking_information_id VARCHAR(26) REFERENCES walking_information (id)
);
