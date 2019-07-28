DROP TABLE IF EXISTS servers;

CREATE TABLE servers
(
    server_id uuid DEFAULT uuid_generate_v4 (),
    server_name TEXT NOT NULL,
    cpu INTEGER,
    memory INTEGER,
    storage NUMERIC,
    PRIMARY KEY (server_id)
);

INSERT INTO servers
    (server_name)
VALUES
    ('CERTTR-PRD-01');
INSERT INTO servers
    (server_name)
VALUES
    ('CARDAX-PRD_01');
INSERT INTO servers
    (server_name)
VALUES
    ('SAPGDB-PRD-01');
INSERT INTO servers
    (server_name)
VALUES
    ('TABLEU-PRD-01');