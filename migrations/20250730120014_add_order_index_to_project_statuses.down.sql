DO $$
BEGIN
    -- Hapus index jika ada
    IF EXISTS (
        SELECT 1
        FROM pg_indexes
        WHERE schemaname = 'public'
          AND tablename = 'project_statuses'
          AND indexname = 'idx_project_statuses_project_order'
    ) THEN
        DROP INDEX public.idx_project_statuses_project_order;
    END IF;

    -- Hapus constraint unique jika ada
    IF EXISTS (
        SELECT 1
        FROM information_schema.table_constraints
        WHERE table_schema = 'public'
          AND table_name = 'project_statuses'
          AND constraint_name = 'unique_project_order'
    ) THEN
        ALTER TABLE public.project_statuses
        DROP CONSTRAINT unique_project_order;
    END IF;

    -- Hapus kolom order_index kalau ada
    IF EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_schema = 'public'
          AND table_name = 'project_statuses'
          AND column_name = 'order_index'
    ) THEN
        ALTER TABLE public.project_statuses
        DROP COLUMN order_index;
    END IF;
END$$;
