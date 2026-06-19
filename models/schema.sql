create table if not exists users (
    id integer primary key autoincrement,
    name text not null,
    login text not null,
    password text not null,
    image_url text not null,
    enabled integer not null default 0
);

create table if not exists categories (
    id integer primary key autoincrement,
    description text not null
);

create table if not exists transactions (
    id integer primary key autoincrement,
    date text not null default (strftime('%d/%m/%Y', 'now')),
    amount int not null default 0,
    type text not null default 'Despesa',
    description text not null,
    category_id integer not null,
    user_id integer not null,
    foreign key (category_id) references categories(id),
    foreign key (user_id) references users(id)
);