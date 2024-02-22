create table if not exists medias(
    id int(11) not null auto_increment primary key,
    medianame varchar(64) not null comment '媒体名称',
    mediatype int(11) not null comment '媒体类型',
    testor varchar(64) not null comment '检测机构',
    jurisdiction varchar(64) not null comment '管辖机构',
    division varchar(64) not null comment '行政区域',
    mediaorganization varchar(64) not null comment'媒介机构',
    distributor varchar(64) not null comment'派发人',
    deadline varchar (64) not null default 0 comment '截止时间',
    created_at timestamp not null default now() comment '创建时间'  
)