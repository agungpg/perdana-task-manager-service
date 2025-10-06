DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 
        FROM information_schema.tables 
        WHERE table_schema = 'public' 
          AND table_name = 'task_change_logs'
    ) THEN
        CREATE TABLE public.task_change_logs (
            id varchar(50) PRIMARY KEY,
            task_id varchar(50) NOT NULL REFERENCES public.tasks(id) ON DELETE CASCADE,
            field_name VARCHAR(100) NOT NULL,
            old_value TEXT,
            new_value TEXT,
            changed_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
            changed_by varchar(50) REFERENCES public.users(id) ON DELETE SET NULL
        );
    END IF;
END$$;
