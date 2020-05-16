-- Table: public."Product"

-- DROP TABLE public."Product";

CREATE TABLE public."Product"
(
    name   character varying(50) COLLATE pg_catalog."default" NOT NULL,
    id     bigint                                             NOT NULL DEFAULT nextval('"Product_id_seq"'::regclass),
    price  numeric                                                     DEFAULT 0,
    groups character varying[] COLLATE pg_catalog."default",
    CONSTRAINT "Product_pkey" PRIMARY KEY (id)
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

COMMENT ON COLUMN public."Product".groups
    IS '属于的组';