DO $$
BEGIN
    -- Change description type back from TEXT to varchar(100)
    IF EXISTS (
        SELECT 1 FROM information_schema.columns 
        WHERE table_schema = 'public' 
        AND table_name = 'tasks' 
        AND column_name = 'description'
    ) THEN
        ALTER TABLE public.tasks ALTER COLUMN description TYPE varchar(100);
    END IF;

    -- Rename status_id column back to status
    IF EXISTS (
        SELECT 1 FROM information_schema.columns 
        WHERE table_schema = 'public' 
        AND table_name = 'tasks' 
        AND column_name = 'status_id'
    ) THEN
        ALTER TABLE public.tasks RENAME COLUMN status_id TO status;
    END IF;
END$$; 