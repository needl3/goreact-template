CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  fname character varying(64) NOT NULL,
  lname character varying(64) NOT NULL,
  avatar bpchar NOT NULL,
  email character varying(64) NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE OR REPLACE FUNCTION update_modified_timestamp()   
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;   
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_modtime BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE  update_modified_timestamp();
