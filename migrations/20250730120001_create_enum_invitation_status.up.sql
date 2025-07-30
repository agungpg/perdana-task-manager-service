DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'invitation_status_enum') THEN
        CREATE TYPE public.invitation_status_enum AS ENUM (
            'pending',
            'accepted',
            'rejected'
        );
    END IF;
END$$; 