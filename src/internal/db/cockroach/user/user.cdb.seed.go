package user

var SQL = `
insert into public.users (id, updated_at, created_at, deleted_at, username, hash_password, use_default_password, role_id, admin_div_id, is_active)
values  ('09944b70-a52e-4f38-b856-d7a0d987e8ce', 1640002866414, 1640002866414, 0, 'citizen01010101', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, 'f508a982-2016-45e5-a429-ba1bd229dca4', '33f12b81-fa28-4a47-8beb-eb1ee352b89f', true),
        ('27e99822-65b0-4b97-9062-8d5f3ca7a6b7', 1639926636491, 1639926636491, 0, 'citizen0101', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '89c913de-9254-4e8c-a09b-492d5a248f28', '5e3363c6-0178-4fa9-a71c-1af2f9c0708d', true),
        ('622924ba-ec91-405f-aa61-851a295bb678', 1639926538564, 1639926538564, 0, 'citizen03', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '68a264e2-4f1b-42a5-a0ff-5c4ea50ef15f', 'b996248f-7fa8-44e1-b695-4eb8b821ec6d', true),
        ('65066f74-4a1a-4d6b-bfa1-1bee6f494060', 1639926483693, 1639926483693, 0, 'citizen01', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '68a264e2-4f1b-42a5-a0ff-5c4ea50ef15f', 'fbe83e5b-4416-4caa-96c4-2b825e687415', true),
        ('7c0a6259-b82a-4c29-a7a3-d18d6b2a4b87', 1639926692528, 1639926692528, 0, 'citizen010101', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '3157db78-b283-4a23-8b73-669b27be17fa', 'f5ccb72f-12bb-4b91-8b75-946696fd3322', true),
        ('8cf29dd8-eed0-4ade-823f-3a6fae9be067', 1639926529744, 1639926529744, 0, 'citizen02', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '68a264e2-4f1b-42a5-a0ff-5c4ea50ef15f', '7c7871d9-2d32-44c6-993b-ae1a67ac88fa', true),
        ('d531dc36-5e31-4941-9ca1-ec87eaca9b95', 1639650462801, 1639650462801, 0, 'admin', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '824db1b0-e20d-48cd-96b5-c055e502ec9f', '00000000-0000-0000-0000-000000000000', true);
`
