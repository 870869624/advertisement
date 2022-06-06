create table if not exists users (
    id int(11) not null auto_increment primary key,
    username varchar(64) not null comment '账号',
    name varchar (64) not null comment '姓名',
    password varchar(64) not null  comment '密码',
    telephone varchar(64) not null comment '电话号码',
    jurisdiction varchar(64) not null comment '管理机构',
    created_at timestamp not null default now() comment '创建时间'
)