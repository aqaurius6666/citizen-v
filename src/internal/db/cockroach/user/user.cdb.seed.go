package user

var (
	SQL = `
	insert into public.users (id, updated_at, created_at, deleted_at, username, hash_password, use_default_password, role_id, admin_div_id)
values  ('23645fca-30b7-4b41-8cda-03a914fb6198', 1639477058794, 1639477020848, 0, 'citizen0111', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '4ae5030a-78f0-491c-9b9e-299c96b2ed82', '6c4894c7-c784-4ff9-9fe8-14078b81d2f4'),
        ('44ea3e96-9283-4a0d-98c5-850615e70425', 1639476919381, 1639476871807, 0, 'citizen01', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '68a264e2-4f1b-42a5-a0ff-5c4ea50ef15f', 'ad9c97f8-65b8-4bd9-9dbb-2bb44cca1410'),
        ('8993850e-9445-4aec-946c-fdb63698739d', 1639477133232, 1639477133232, 0, 'citizen011101', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '3157db78-b283-4a23-8b73-669b27be17fa', 'd54603d5-224b-4b5a-8f47-5bdc3e8621b8');`
)
