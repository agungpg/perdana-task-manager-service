DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.columns 
        WHERE table_schema = 'public' 
        AND table_name = 'projects' 
        AND column_name = 'status_default'
    ) THEN
        ALTER TABLE public.projects
        ADD COLUMN status_default varchar(50) REFERENCES public.project_statuses(id);
    END IF;
END$$; 