create table if not exists mediaorganization(
    id int(11) not null auto_increment primary key,
    mediaorname varchar(64) not null comment '媒介机构名称',
    type varchar(64) not null comment '类型',
    division varchar(64) not null comment '归属行政区域',
    created_at timestamp not null default now() comment '创建时间'
)