drop table if exists nodes;
drop table if exists scans;
drop table if exists scans_nodes;
create table nodes (
    mac_address text not null primary key,
    ip_address text not null,
    vendor text not null,
    registry text not null,
    organization text not null,
    address text not null,
    visible integer not null
);
create table scans (id integer not null primary key);
create table scans_nodes (
    id integer not null primary key,
    node_id text not null,
    scan_id integer not null
)