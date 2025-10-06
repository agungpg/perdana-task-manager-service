DO $$
BEGIN
    -- Tambah kolom order_index kalau belum ada
    IF NOT EXISTS (
        SELECT 1 
        FROM information_schema.columns 
        WHERE table_schema = 'public'
          AND table_name = 'project_statuses'
          AND column_name = 'order_index'
    ) THEN
        ALTER TABLE public.project_statuses
        ADD COLUMN order_index INT DEFAULT 0;
    END IF;

    -- Update order_index supaya unik per project_id
    UPDATE public.project_statuses ps
    SET order_index = ordered.new_order
    FROM (
        SELECT id, project_id,
               ROW_NUMBER() OVER (PARTITION BY project_id ORDER BY created_at ASC, id ASC) - 1 AS new_order
        FROM public.project_statuses
    ) AS ordered
    WHERE ps.id = ordered.id;

    -- Tambah constraint unique (project_id, order_index) kalau belum ada
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.table_constraints
        WHERE table_schema = 'public'
          AND table_name = 'project_statuses'
          AND constraint_name = 'unique_project_order'
    ) THEN
        ALTER TABLE public.project_statuses
        ADD CONSTRAINT unique_project_order UNIQUE (project_id, order_index);
    END IF;

    -- Tambah index untuk mempercepat query ORDER BY project_id, order_index
    IF NOT EXISTS (
        SELECT 1
        FROM pg_indexes
        WHERE schemaname = 'public'
          AND tablename = 'project_statuses'
          AND indexname = 'idx_project_statuses_project_order'
    ) THEN
        CREATE INDEX idx_project_statuses_project_order
        ON public.project_statuses (project_id, order_index);
    END IF;
END$$;
