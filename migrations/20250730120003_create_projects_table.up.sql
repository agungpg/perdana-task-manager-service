DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'projects') THEN
        CREATE TABLE public.projects (
            id varchar(50) PRIMARY KEY,
            thumbnail text,
            name varchar(100) NOT NULL,
            description text,
            created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
            created_by varchar(50),
            updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
            updated_by varchar(50),
            is_deleted boolean DEFAULT false
        );
    END IF;
END$$; 