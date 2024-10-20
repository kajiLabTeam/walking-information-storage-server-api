\c indoor_location_estimation;

CREATE TABLE buildings (
    id VARCHAR(26) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE floors (
    id VARCHAR(26) PRIMARY KEY,
    level INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    building_id VARCHAR(26) REFERENCES buildings(id)
);

CREATE TABLE floor_information (
    id VARCHAR(26) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    floor_id VARCHAR(26) REFERENCES floors(id)
);

CREATE TABLE pedestrians (
    id VARCHAR(26) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE walking_information (
    id VARCHAR(26) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    pedestrian_id VARCHAR(26) REFERENCES pedestrians(id)
);

CREATE TABLE trajectories (
    id VARCHAR(26) PRIMARY KEY,
    is_walking BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    pedestrian_id VARCHAR(26) REFERENCES pedestrians(id),
    floor_information_id VARCHAR(26) REFERENCES floor_information(id)
);

CREATE TABLE correct_positions (
    id VARCHAR(26) PRIMARY KEY,
    x DECIMAL NOT NULL,
    y DECIMAL NOT NULL,
    direction DECIMAL NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    trajectory_id VARCHAR(26) REFERENCES trajectories(id)
);

CREATE TABLE walking_samples (
    id VARCHAR(26) PRIMARY KEY,
    is_converged BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    trajectory_id VARCHAR(26) REFERENCES trajectories(id),
    walking_information_id VARCHAR(26) UNIQUE REFERENCES walking_information(id)
);
