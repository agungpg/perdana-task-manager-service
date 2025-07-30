DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'sub_tasks') THEN
        CREATE TABLE public.sub_tasks (
            id varchar(50) PRIMARY KEY,
            task_id varchar(50) REFERENCES public.tasks(id),
            name varchar(50),
            created_by varchar(50),
            created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
            updated_by varchar(50),
            updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
            is_deleted boolean DEFAULT false,
            assigned_to varchar(50) REFERENCES public.users(id),
            completed_date timestamptz,
            is_completed boolean DEFAULT false
        );
    END IF;
END$$; 