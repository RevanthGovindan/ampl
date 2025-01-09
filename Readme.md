SELECT typname FROM pg_type WHERE typtype = 'e';
CREATE TYPE status_type AS ENUM('pending', 'in-progress', 'completed');