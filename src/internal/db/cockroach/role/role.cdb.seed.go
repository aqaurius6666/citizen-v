package role

var SQL = `
insert into public.roles (id, updated_at, created_at, deleted_at, name)
values  ('3157db78-b283-4a23-8b73-669b27be17fa', 1637828766653, 1637828766653, 0, 'role-b1'),
        ('68a264e2-4f1b-42a5-a0ff-5c4ea50ef15f', 1637828716389, 1637828716389, 0, 'role-a2'),
        ('824db1b0-e20d-48cd-96b5-c055e502ec9f', 1637831958338, 1637831958338, 0, 'role-a1'),
        ('89c913de-9254-4e8c-a09b-492d5a248f28', 1637828746206, 1637828746206, 0, 'role-a3'),
        ('f508a982-2016-45e5-a429-ba1bd229dca4', 1637828784984, 1637828784984, 0, 'role-b2');
`
