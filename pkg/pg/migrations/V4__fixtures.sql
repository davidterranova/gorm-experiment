INSERT INTO users (id, first_name, last_name, email) VALUES
  ('00000000-0000-0000-0000-000000000001', 'Luke', 'Skywalkwer', 'lskywalker@test.local'),
  ('00000000-0000-0000-0000-000000000002', 'Dark', 'Vador', 'dvador@test.local')
;

INSERT INTO roles (id, name) VALUES
  ('00000000-0000-0000-0000-000000000001', 'admin'),
  ('00000000-0000-0000-0000-000000000002', 'user')
;

INSERT INTO users_roles (user_id, role_id, notes) VALUES
  ('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000001', 'general admin'),
  ('00000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000002', 'default role for new users'),
  ('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000002', 'default role for new users')
;
