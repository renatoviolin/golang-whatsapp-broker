CREATE SEQUENCE agent_id_seq;
CREATE TABLE agents (
    id bigint PRIMARY KEY NOT NULL DEFAULT nextval('agent_id_seq'),
    name text,
    read_topic text,
    write_topic text,
    error_topic text,
    active integer DEFAULT 1
);

INSERT INTO agents (name, read_topic, write_topic, error_topic) values ('Agent A', 'agent_A-in', 'agent_A-out', 'agent_A-error');