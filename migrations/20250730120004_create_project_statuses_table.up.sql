DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'project_statuses') THEN
        CREATE TABLE public.project_statuses (
            id varchar(50) PRIMARY KEY,
            project_id varchar(50) REFERENCES public.projects(id),
            name varchar(30) NOT NULL,
            created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
            created_by varchar(50),
            updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
            updated_by varchar(50),
            is_deleted boolean DEFAULT false
        );
    END IF;
END$$; 