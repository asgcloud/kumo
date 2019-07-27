DROP TABLE IF EXISTS servers;

CREATE TABLE servers
(
    server_id TEXT NOT NULL PRIMARY KEY,
    server_name TEXT NOT NULL
);

INSERT INTO servers
VALUES
    (
        "2d9647bc-1124-453c-a59d-1abcccbf67d7",
        "CERTTR-PRD-01"
);

INSERT INTO servers
VALUES
    (
        "68a0db62-0c1f-4af3-9520-d5051e9375ce",
        "CARDAX-PRD_01"
);

INSERT INTO servers
VALUES
    (
        "a5cfc35e-0a85-4f51-b242-fc82b4e834f4",
        "SAPGDB-PRD-01"
);

INSERT INTO servers
VALUES
    (
        "d93527bd-4233-4d2a-814c-99e557e2c4fb",
        "TABLEU-PRD-01"
);