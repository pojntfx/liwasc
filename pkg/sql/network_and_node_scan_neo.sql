-- Remove old tables if they exist
drop table if exists network_scans;
drop table if exists nodes;
drop table if exists node_scans;
drop table if exists services;
-- Create new tables
create table network_scans (
    id integer not null primary key,
    created_at date not null,
    done integer not null
);
create table nodes (
    id integer not null primary key,
    created_at date not null,
    mac_address text not null,
    network_scan_id integer not null,
    foreign key (network_scan_id) references network_scans(id)
);
create table node_scans (
    id integer not null primary key,
    created_at date not null,
    done integer not null,
    node_id integer not null,
    foreign key (node_id) references nodes(id)
);
create table services (
    id integer not null primary key,
    created_at date not null,
    port_number integer not null,
    transport_protocol text not null,
    node_scan_id integer not null,
    foreign key (node_scan_id) references node_scans(id)
);