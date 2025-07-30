DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'tasks') THEN
        CREATE TABLE public.tasks (
            id varchar(50) PRIMARY KEY,
            project_id varchar(50) REFERENCES public.projects(id),
            name varchar(100) NOT NULL,
            description varchar(100),
            thumbnail varchar(50),
            created_by varchar(50),
            created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
            updated_by varchar(50),
            updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
            is_deleted boolean DEFAULT false,
            assigned_to varchar(50),
            reported_to varchar(50),
            priority varchar(10),
            status varchar(50),
            due_date timestamptz,
            completed_date timestamptz,
            CONSTRAINT tasks_priority_check CHECK (priority IN ('low', 'medium', 'high'))
        );
    END IF;
END$$; 