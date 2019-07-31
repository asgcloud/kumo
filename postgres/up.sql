DROP TABLE IF EXISTS servers;
DROP TABLE IF EXISTS users;

CREATE TABLE users
(
    user_id uuid DEFAULT uuid_generate_v4 (),
    username VARCHAR(32) NOT NULL,
    hash VARCHAR(64) NOT NULL,
    PRIMARY KEY (user_id)
);

CREATE TABLE projects
(
    project_id uuid DEFAULT uuid_generate_v4 (),
    project_name VARCHAR(32) NOT NULL,
    host_platform VARCHAR(32) NOT NULL,
    host_ip VARCHAR(32),
    host_username VARCHAR(32),
    host_password VARCHAR(64),
    PRIMARY KEY (project_id)
);

INSERT INTO projects
    (project_name, host_platform)
VALUES
    ('Test Project', 'VMware');

UPDATE projects 
SET 
    project_id = 'f90e8097-ddac-49dd-978f-0e88bf90891e' 
WHERE 
    project_name = 'Test Project';

CREATE TABLE servers
(
    server_id uuid DEFAULT uuid_generate_v4 (),
    project_id uuid NOT NULL,
    server_name VARCHAR(32) NOT NULL,
    cpu INTEGER,
    ram INTEGER,
    storage NUMERIC,
    server_status VARCHAR(32),
    server_state VARCHAR(32),
    tenancy VARCHAR(32),
    host VARCHAR(32),
    PRIMARY KEY (server_id),
    FOREIGN KEY (project_id) REFERENCES projects(project_id)
);

INSERT INTO servers
    (project_id, server_name, cpu, ram, storage, server_status, server_state, tenancy, host)
VALUES
    ('f90e8097-ddac-49dd-978f-0e88bf90891e', 'CERTTR-PRD-01', 2, 8, 32.0, 'Active', 'Powered On', 'Production', 'Azure');
INSERT INTO servers
    (project_id, server_name, cpu, ram, storage, server_status, server_state, tenancy, host)
VALUES
    ('f90e8097-ddac-49dd-978f-0e88bf90891e', 'CARDAX-PRD_01', 2, 8, 32.0, 'Active', 'Powered On', 'Production', 'Azure');
INSERT INTO servers
    (project_id, server_name, cpu, ram, storage, server_status, server_state, tenancy, host)
VALUES
    ('f90e8097-ddac-49dd-978f-0e88bf90891e', 'SAPGDB-PRD-01', 2, 8, 32.0, 'Active', 'Powered On', 'Production', 'Azure');
INSERT INTO servers
    (project_id, server_name, cpu, ram, storage, server_status, server_state, tenancy, host)
VALUES
    ('f90e8097-ddac-49dd-978f-0e88bf90891e', 'TABLEU-PRD-01', 2, 8, 32.0, 'Active', 'Powered On', 'Production', 'Azure');