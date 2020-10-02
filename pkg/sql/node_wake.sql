-- Remove old tables if they exist
drop table if exists nodes;
drop table if exists node_wakes;
drop table if exists node_wakes_nodes;
-- Create new tables
create table nodes (
    mac_address text not null primary key,
    powered_on integer not null
);
create table node_wakes (
    id integer not null primary key,
    created_at date not null,
    done integer not null
);
-- Create join tables
create table node_wakes_nodes (
    id integer not null primary key,
    created_at date not null,
    node_id text not null,
    node_wakes_id integer not null
);