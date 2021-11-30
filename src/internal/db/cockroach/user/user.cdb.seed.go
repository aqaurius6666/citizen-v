package user

var (
	SQL = `
	insert into public.users (id, updated_at, created_at, deleted_at, username, hash_password, role_id)
	values  ('090ed02e-42c4-4880-8f7d-d761f4711035', 1638283533683, 1638283533683, 0, 'citizen01', '123456', '68a264e2-4f1b-42a5-a0ff-5c4ea50ef15f'),
			('afefa5a6-3d04-4287-8428-c21c1908075d', 1638283537496, 1638283537496, 0, 'citizen02', '123456', '68a264e2-4f1b-42a5-a0ff-5c4ea50ef15f'),
			('f49f0714-3578-4ca4-b01f-400fa40a8f32', 1638283541037, 1638283541037, 0, 'citizen03', '123456', '68a264e2-4f1b-42a5-a0ff-5c4ea50ef15f');
	`
)
