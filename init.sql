create table "user"
(
    id         serial primary key not null,
    username   varchar unique     not null,
    created_at timestamp default now()
);

create table "chat"
(
    id         serial primary key not null,
    name       varchar            not null,
    created_at timestamp default now()
);

create table "user_to_chat"
(
    user_id int references "user" (id),
    chat_id int references "chat" (id)
);

create table "message"
(
    id         serial primary key not null,
    chat       int references "chat" (id),
    author     int references "user" (id),
    text       text,
    created_at timestamp default now()
);

create or replace function chats(uid int)
    returns table
            (
                id         int,
                name       varchar,
                created_at timestamp
            )
    language sql
as
$$
with user_chats as
         (
             select chat.id, chat.name, chat.created_at
             from user_to_chat as utc
                      join chat on utc.chat_id = chat.id
             where utc.user_id = uid
         ),
     msg as
         (
             select chat, max(created_at) as last_msg_ts
             from "message"
             where author = uid
             group by chat
         )
select user_chats.*
from user_chats
         left join msg on user_chats.id = msg.chat
order by msg.last_msg_ts desc;
$$;

create or replace function messages(chat_id int)
    returns table
            (
                id         int,
                chat       varchar,
                author     varchar,
                text       text,
                created_at timestamp
            )
    language sql
as
$$
with m_u as
         (
             select m.id, m.chat, u.username, m.text, m.created_at
             from "message" as m
                      join "user" as u on m.author = u.id
             where m.chat = chat_id
         )

select m_u.id, chat.name, m_u.username, m_u.text, m_u.created_at
from m_u
         join chat on m_u.chat = chat.id
order by m_u.created_at desc;
$$;