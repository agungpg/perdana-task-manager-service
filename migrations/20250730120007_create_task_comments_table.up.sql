DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'task_comments') THEN
        CREATE TABLE public.task_comments (
            id varchar(50) PRIMARY KEY,
            task_id varchar(50) REFERENCES public.tasks(id) ON DELETE CASCADE,
            content text NOT NULL,
            created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
            created_by varchar(50) REFERENCES public.users(id),
            updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
            updated_by varchar(50),
            is_deleted boolean DEFAULT false
        );
    END IF;
END$$; 