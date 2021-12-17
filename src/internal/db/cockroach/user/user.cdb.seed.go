package user

var (
	SQL = `
    insert into public.users (id, updated_at, created_at, deleted_at, username, hash_password, use_default_password, role_id, admin_div_id, is_active)
    values  ('128999d9-f46c-41f1-b65c-fcbbcd563d4c', 1639651538666, 1639651538666, 0, 'citizen0301', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '4ae5030a-78f0-491c-9b9e-299c96b2ed82', '2eed7d62-9a92-445f-8b8c-e109550682dd', true),
            ('15829907-83c2-4b9a-84aa-97e278ae753c', 1639650848995, 1639650848995, 0, 'citizen03', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '68a264e2-4f1b-42a5-a0ff-5c4ea50ef15f', '35bc69a6-1a46-4060-ae5e-3352fe21426b', true),
            ('23645fca-30b7-4b41-8cda-03a914fb6198', 1639733471667, 1639477020848, 0, 'citizen0110', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '4ae5030a-78f0-491c-9b9e-299c96b2ed82', '288d832f-2281-40ae-8023-4ffbb7527064', true),
            ('306c875c-3145-499a-aa33-9d3ed850152c', 1639728265599, 1639728265599, 0, 'citizen0109', 'uMc+Nh6J8uImcmf0GkbQ1DFPJLfkuiFfitENz29QzZ4=', true, '4ae5030a-78f0-491c-9b9e-299c96b2ed82', '18917370-7f8c-4de4-b9e7-1c0bb004badc', true),
            ('44ea3e96-9283-4a0d-98c5-850615e70425', 1639476919381, 1639476871807, 0, 'citizen01', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '68a264e2-4f1b-42a5-a0ff-5c4ea50ef15f', 'abb61a29-53f5-44a8-afdb-1ab9b9dab865', true),
            ('8993850e-9445-4aec-946c-fdb63698739d', 1639477133232, 1639477133232, 0, 'citizen011001', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '3157db78-b283-4a23-8b73-669b27be17fa', '6a31b531-1bb5-496a-b00f-d7cf6161bf6f', true),
            ('b3a8e813-df81-46e6-9037-b4cf39db1a73', 1639650462801, 1639650462801, 0, 'citizen10', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '68a264e2-4f1b-42a5-a0ff-5c4ea50ef15f', 'd4c3a365-c073-42af-ba0a-598b90256647', true),
            ('d531dc36-5e31-4941-9ca1-ec87eaca9b95', 1639650462801, 1639650462801, 0, 'admin', 'Ep1wn1V67h4ytvjI6ZCf7wcIrEf8rjfAmKAGdcfCfSU=', false, '824db1b0-e20d-48cd-96b5-c055e502ec9f', '00000000-0000-0000-0000-000000000000', true);    
    `
)
