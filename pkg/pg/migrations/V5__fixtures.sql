INSERT INTO users (id, first_name, last_name) VALUES
  ('00000000-0000-0000-0000-000000000001', 'Luke', 'Skywalkwer'),
  ('00000000-0000-0000-0000-000000000002', 'Dark', 'Vador')
;

INSERT INTO emails (id, user_id, email, principal, verified) VALUES
  ('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000001', 'lskywalker@test.local', true, true),
  ('00000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000001', 'luke.skywalker@test.local', false, true),
  ('00000000-0000-0000-0000-000000000003', '00000000-0000-0000-0000-000000000002', 'dvador@test.local', true, false)
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