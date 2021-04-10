-- +migrate Up
create table node_scans (
    id integer not null primary key,
    created_at date not null,
    done integer not null
);
create table nodes (
    id integer not null primary key,
    created_at date not null,
    mac_address text not null,
    ip_address text not null,
    node_scan_id integer not null,
    foreign key (node_scan_id) references node_scans(id)
);
create table port_scans (
    id integer not null primary key,
    created_at date not null,
    done integer not null,
    node_id integer not null,
    foreign key (node_id) references nodes(id)
);
create table ports (
    id integer not null primary key,
    created_at date not null,
    port_number integer not null,
    transport_protocol text not null,
    port_scan_id integer not null,
    foreign key (port_scan_id) references port_scans(id)
);
-- +migrate Down
drop table node_scans;
drop table nodes;
drop table port_scans;
drop table ports;