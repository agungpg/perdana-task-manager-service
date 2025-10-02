DO $$
BEGIN
    -- Change description type from varchar(100) to TEXT
    IF EXISTS (
        SELECT 1 FROM information_schema.columns 
        WHERE table_schema = 'public' 
        AND table_name = 'tasks' 
        AND column_name = 'description'
    ) THEN
        ALTER TABLE public.tasks ALTER COLUMN description TYPE TEXT;
    END IF;

    -- Rename status column to status_id
    IF EXISTS (
        SELECT 1 FROM information_schema.columns 
        WHERE table_schema = 'public' 
        AND table_name = 'tasks' 
        AND column_name = 'status'
    ) THEN
        ALTER TABLE public.tasks RENAME COLUMN status TO status_id;
    END IF;
END$$; 