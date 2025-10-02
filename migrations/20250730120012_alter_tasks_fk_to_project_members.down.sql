DO $$
BEGIN
    -- Drop project_members foreign key constraints if they exist
    IF EXISTS (
        SELECT 1 FROM information_schema.table_constraints 
        WHERE constraint_name = 'tasks_assigned_to_project_member_fkey' 
        AND table_name = 'tasks'
    ) THEN
        ALTER TABLE public.tasks DROP CONSTRAINT tasks_assigned_to_project_member_fkey;
    END IF;

    IF EXISTS (
        SELECT 1 FROM information_schema.table_constraints 
        WHERE constraint_name = 'tasks_reported_to_project_member_fkey' 
        AND table_name = 'tasks'
    ) THEN
        ALTER TABLE public.tasks DROP CONSTRAINT tasks_reported_to_project_member_fkey;
    END IF;

    -- Add back foreign key constraints to users
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints 
        WHERE constraint_name = 'tasks_assigned_to_fkey' 
        AND table_name = 'tasks'
    ) THEN
        ALTER TABLE public.tasks 
        ADD CONSTRAINT tasks_assigned_to_fkey 
        FOREIGN KEY (assigned_to) REFERENCES public.users(id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.table_constraints 
        WHERE constraint_name = 'tasks_reported_to_fkey' 
        AND table_name = 'tasks'
    ) THEN
        ALTER TABLE public.tasks 
        ADD CONSTRAINT tasks_reported_to_fkey 
        FOREIGN KEY (reported_to) REFERENCES public.users(id);
    END IF;
END$$; 