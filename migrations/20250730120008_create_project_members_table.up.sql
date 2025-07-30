DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'project_members') THEN
        CREATE TABLE public.project_members (
            id varchar(50) PRIMARY KEY,
            project_id varchar(50) REFERENCES public.projects(id),
            user_id varchar(50) REFERENCES public.users(id),
            is_admin boolean DEFAULT false,
            created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
            created_by varchar(50),
            updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
            updated_by varchar(50),
            is_deleted boolean DEFAULT false,
            invitation_status public.invitation_status_enum DEFAULT 'pending'
        );
    END IF;
END$$; 