create table if not exists areas(
    id int(11) not null auto_increment primary key,
    name varchar(64) not null comment '行政区域名称',
    level int(11) not null comment'行政区域等级',
    pid int(11) not null comment '所属行政区域',
    left_ int(11) not null comment '所属范围',
    right_ int(11) not null comment '所属范围',
    INDEX(left_),
    INDEX(right_),
    created_at timestamp not null default now() comment '创建时间'
)
