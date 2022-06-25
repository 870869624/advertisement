create table if not exists mediaorganizations(
    id int(11) not null auto_increment primary key,
    mediaorname varchar(64) not null comment '媒介机构名称',
    type int(11) not null comment '类型',
    area_id int(11) not null unsigned comment '归属行政区域',
    created_at timestamp not null default now() comment '创建时间'
)