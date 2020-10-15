-- Remove old tables if they exist
drop table if exists node_wakes;
-- Create new tables
create table node_wakes (
    id integer not null primary key,
    created_at date not null,
    done integer not null,
    mac_address text not null,
    powered_on integer not null
);