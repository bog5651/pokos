CREATE TABLE clients
(
    "id"   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
    "name" TEXT
);

CREATE TABLE "cash_desk"
(
    "id"   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
    "name" TEXT
);

CREATE TABLE "kkm_registers"
(
    "id"                   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
    "client_id"            INTEGER,
    "cash_desk_id"         INTEGER,
    "serial_number"        TEXT,
    "register_date"        TEXT,
    "ofd"                  TEXT,
    "isExcise"             INTEGER,
    "system_no"            TEXT,
    "type"                 TEXT,
    "fn"                   TEXT,
    "address"              TEXT,
    "end_date_fn"          TEXT,
    "end_date_ofd"         TEXT,
    "inspection_day_count" INTEGER,
    "comment"              TEXT
);
-- select query
SELECT k.id,
       c.name,
       cd.name,
       k.client_id,
       k.cash_desk_id,
       k.serial_number,
       k.register_date,
       k.ofd,
       k.isExcise,
       k.system_no,
       k.type,
       k.fn,
       k.address,
       k.end_date_fn,
       k.end_date_ofd,
       k.inspection_day_count,
       k.comment
FROM kkm_registers k
         LEFT JOIN clients AS c ON c.id = k.client_id
         LEFT JOIN cash_desk AS cd ON cd.id = k.cash_desk_id;

SELECT id, name
FROM cash_desk;

SELECT id, name
FROM clients;
