-- 创建数据库
create database test with template template0 lc_collate "zh_CN.UTF8" lc_ctype "zh_CN.UTF8" encoding 'UTF8';
\c test

-- 增加系统表
CREATE TABLE public.sys_info
(
    id serial PRIMARY KEY NOT NULL,
    name varchar(50) DEFAULT '' NOT NULL,
    version varchar(20) DEFAULT '0.0.0' NOT NULL
);
COMMENT ON COLUMN public.sys_info.name IS '网站名称';
COMMENT ON COLUMN public.sys_info.version IS '版本号';
COMMENT ON TABLE public.sys_info IS '系统信息表';

-- 添加版本号
INSERT INTO public.sys_info(name,version) values('账户接口','0.0.0')