
CREATE TABLE public."Product"
(
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    id bigint NOT NULL DEFAULT nextval('"Product_id_seq"'::regclass),
    price numeric DEFAULT 0
)

TABLESPACE pg_default;

ALTER TABLE public."Product"
    OWNER to postgres;

COMMENT ON COLUMN public."Product".name
    IS '产品名称';

COMMENT ON COLUMN public."Product".id
    IS '编号';

COMMENT ON COLUMN public."Product".price
    IS '价格';

CREATE TABLE public."User"
(
    id bigint NOT NULL,
    name character varying(20) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "User_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public."User"
    OWNER to postgres;
COMMENT ON TABLE public."User"
    IS '用户信息表';

