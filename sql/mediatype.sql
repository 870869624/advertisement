create table if not exists mediatype(
    id int(11) not null auto_increment primary key,
    name varchar(64) not null  comment '类型名称',
    created_at timestamp not null default now() comment '创建时间'
)