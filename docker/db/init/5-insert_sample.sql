\c indoor_location_estimation;


INSERT INTO pedestrians (id)
VALUES
    ('01F8VYXK67BGC1F9RP1E4S9YAK');

INSERT INTO walking_information (id, pedestrian_id)
VALUES
    ('01JBJR8PMCEA0KVAAM0X5R5PNJ', '01F8VYXK67BGC1F9RP1E4S9YAK');

INSERT INTO gyroscopes (id, walking_information_id)
VALUES
    ('01JBJRAYD6RCAY6NKYAWRJ44S3', '01JBJR8PMCEA0KVAAM0X5R5PNJ');

INSERT INTO buildings (id, name, latitude, longitude)
VALUES
    ('01F8VYXK67BGC1F9RP1E4S9YAK', '14号館', 35.1847559, 137.1110031);

INSERT INTO buildings (id, name, latitude, longitude)
VALUES
    ('01JET6FBCS6M93EP96FJABWD82', 'デモ用建物', 35.1847559, 137.1110031);

INSERT INTO floors (id, level, name, building_id)
VALUES
    ('01F8VYXK67BGC1F9RP1E4S9YTV', 5, '14号館5階', '01F8VYXK67BGC1F9RP1E4S9YAK');

INSERT INTO floors (id, level, name, building_id)
VALUES
    ('01JET6KMY6558FGHD34XFJZ2VN', 1, 'デモ用建物1階', '01JET6FBCS6M93EP96FJABWD82');

INSERT INTO floor_information (id, floor_id)
VALUES
    ('01J5MT4907RHS5Q8QT583YXB96', '01F8VYXK67BGC1F9RP1E4S9YTV');

INSERT INTO floor_information (id, floor_id)
VALUES
    ('01JET6M5N79HZ92CFZQZVKFCPZ', '01JET6KMY6558FGHD34XFJZ2VN');

INSERT INTO floor_maps (id, floor_information_id)
VALUES
    ('01J41536W3YYS9KAJC2TVC9JD7', '01J5MT4907RHS5Q8QT583YXB96');

INSERT INTO floor_maps (id, floor_information_id)
VALUES
    ('01JET6N184W4EM2HP6YY08XK0E', '01JET6M5N79HZ92CFZQZVKFCPZ');

INSERT INTO trajectories (id, is_walking, pedestrian_id, floor_information_id)
VALUES
    ('01JET1DV4WJ2EP78B8HAKK5SP0', TRUE, '01F8VYXK67BGC1F9RP1E4S9YAK', '01JET6M5N79HZ92CFZQZVKFCPZ');
