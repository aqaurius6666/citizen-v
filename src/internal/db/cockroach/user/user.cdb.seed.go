package user

var (
	SQL = `
	insert into public.users (id, updated_at, created_at, deleted_at, username, hash_password, use_default_password, role_id, admin_div_id, is_active)
values  ('23645fca-30b7-4b41-8cda-03a914fb6198', 1639477058794, 1639477020848, 0, 'citizen0110', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '4ae5030a-78f0-491c-9b9e-299c96b2ed82', '288d832f-2281-40ae-8023-4ffbb7527064', true),
        ('44ea3e96-9283-4a0d-98c5-850615e70425', 1639476919381, 1639476871807, 0, 'citizen01', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '68a264e2-4f1b-42a5-a0ff-5c4ea50ef15f', 'abb61a29-53f5-44a8-afdb-1ab9b9dab865', true),
        ('8993850e-9445-4aec-946c-fdb63698739d', 1639477133232, 1639477133232, 0, 'citizen011001', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '3157db78-b283-4a23-8b73-669b27be17fa', '6a31b531-1bb5-496a-b00f-d7cf6161bf6f', true);
	`
)
