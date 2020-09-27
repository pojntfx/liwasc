drop table if exists nodes;
drop table if exists scans;
create table nodes (
    id integer not null primary key,
    scan_id integer not null references scans(id),
    powered_on integer not null,
    mac_address text not null,
    ip_address text not null,
    vendor text not null,
    registry text not null,
    organization text not null,
    address text not null,
    visible integer not null
);
create table scans (id integer not null primary key);