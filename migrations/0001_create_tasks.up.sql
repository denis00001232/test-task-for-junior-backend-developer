CREATE TABLE IF NOT EXISTS tasks (
	id BIGSERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT NOT NULL DEFAULT '',
	status TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	periodicity_type TEXT,
	daily_interval INTEGER,
	monthly_days INTEGER[],
	specific_dates TIMESTAMPTZ[],
	odd_even_type TEXT
);

CREATE INDEX IF NOT EXISTS idx_tasks_status ON tasks (status);
