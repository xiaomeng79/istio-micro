
-- 创建账户表
SET client_encoding = 'UTF8';

CREATE TABLE public.account (
    id serial PRIMARY KEY NOT NULL,
    user_id integer DEFAULT 0 NOT NULL,
    account_level smallint DEFAULT 0 NOT NULL,
    balance numeric(18,2) DEFAULT 0.00 NOT NULL,
    account_status smallint DEFAULT 0 NOT NULL
);

COMMENT ON TABLE public.account IS '账户表';

COMMENT ON COLUMN public.account.user_id IS '用户id';

COMMENT ON COLUMN public.account.account_level IS '账户级别';

COMMENT ON COLUMN public.account.balance IS '账户余额';

COMMENT ON COLUMN public.account.account_status IS '账户状态';
