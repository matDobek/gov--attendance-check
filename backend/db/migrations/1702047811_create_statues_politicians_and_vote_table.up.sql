begin;

create table politicians (
  id integer GENERATED ALWAYS AS IDENTITY (START WITH 1000) PRIMARY KEY,
  name varchar(255),
  party varchar(255)
);

create table statues (
  id integer GENERATED ALWAYS AS IDENTITY (START WITH 1000) PRIMARY KEY,
  voting_number  integer not null,
  session_number integer not null,
  term_number    integer not null,
  title          varchar(255) not null
);

create table votes (
  id integer GENERATED ALWAYS AS IDENTITY (START WITH 1000) PRIMARY KEY,
  politician_id integer not null references politicians(id) on delete cascade,
  statue_id     integer not null references statues(id)     on delete cascade,
  response varchar(255)
);

commit;
