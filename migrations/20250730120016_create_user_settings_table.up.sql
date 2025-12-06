DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM information_schema.tables
        WHERE table_schema = 'public'
          AND table_name = 'user_settings'
    ) THEN
        CREATE TABLE public.user_settings (
            id varchar(50) PRIMARY KEY,
            user_id varchar(50),
            active_project_id varchar(50),
            theme varchar(32) NOT NULL DEFAULT 'light',
            language varchar(8) NOT NULL DEFAULT 'en',
            created_at timestamptz NOT NULL DEFAULT now(),
            updated_at timestamptz NOT NULL DEFAULT now(),
            CONSTRAINT fk_user_settings_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE,
            CONSTRAINT fk_user_settings_project FOREIGN KEY (active_project_id) REFERENCES public.projects(id),
            CONSTRAINT user_settings_theme_check CHECK (theme IN ('light', 'dark'))
        );
    END IF;
END$$;
