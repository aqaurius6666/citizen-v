package citizen

var SQL = `
insert into public.citizens (id, updated_at, created_at, deleted_at, name, birthday, pid, gender, nationality, current_place, residence_place, hometown, job_name, religion, educational_level, father_name, father_pid, mother_name, mother_pid, admin_div_code, admin_div_id)
values  ('0203bec7-0e51-4abd-94db-f965fbdec66a', 1638279717616, 1638279717616, 0, 'Nguyễn Văn C', 1638279717616, '001201010023', 'male', 'Việt Nam', 'Hà Nội, Đông Anh, Nguyên Khê', 'Hà Nội, Đông Anh, Nguyên Khê', 'Hà Nội, Đông Anh, Nguyên Khê', 'Sinh viên', 'Không', '12/12', 'Nguyễn Ba C', '001201010001', 'Nguyễn Mẹ C', '001201010002', '06050101', '6e3890e5-4e89-4661-812d-b7f1987fa8d4'),
        ('1abebf8c-3af9-47a1-8907-40dc49d75da1', 1638279717616, 1638279717616, 0, 'Nguyễn Duy C', 1638279724474, '001201010015', 'male', 'Việt Nam', 'Hà Nội, Đông Anh, Liên Hà', 'Hà Nội, Đông Anh, Liên Hà', 'Hà Nội, Đông Anh, Liên Hà', 'Tài xế', 'Không', '12/12', 'Nguyễn Ba A', '001201010008', 'Nguyễn Mẹ A', '001201011002', '06050101', '6e3890e5-4e89-4661-812d-b7f1987fa8d4'),
        ('1d0bc669-462e-4479-95a6-18da17d47969', 1638279717616, 1638279717616, 0, 'Nguyễn Thanh K', 1638279724474, '001201010016', 'female', 'Việt Nam', 'Hà Nội, Đông Anh, Thị Trấn Đông Anh', 'Hà Nội, Đông Anh, Thị Trấn Đông Anh', 'Hà Nội, Đông Anh, Thị Trấn Đông Anh', 'Tài xế', 'Không', '12/12', 'Nguyễn Ba A', '001201010008', 'Nguyễn Mẹ A', '001201011002', '06050101', '6e3890e5-4e89-4661-812d-b7f1987fa8d4'),
        ('297a2206-72a6-40d5-84a2-114b896dabca', 1638279717616, 1638279717616, 0, 'Nguyễn Thị C', 1638279717616, '001201010022', 'female', 'Việt Nam', 'Hà Nội, Đông Anh, Dục Tú', 'Hà Nội, Đông Anh, Dục Tú', 'Hà Nội, Đông Anh, Dục Tú', 'Sinh viên', 'Không', '12/12', 'Nguyễn Ba C', '001201010001', 'Nguyễn Mẹ C', '001201010002', '06050101', '6e3890e5-4e89-4661-812d-b7f1987fa8d4'),
        ('7f1cd91c-c050-4ad1-ba95-d0394428242e', 1638279724474, 1638279724474, 0, 'Nguyễn Minh H', 1638279724474, '001201010017', 'male', 'Việt Nam', 'Hà Nội, Đông Anh, Tàm Xá', 'Hà Nội, Đông Anh, Tàm Xá', 'Hà Nội, Đông Anh, Tàm Xá', 'Công Nhân', 'Không', '12/12', 'Nguyễn Ba B', '001201010009', 'Nguyễn Mẹ B', '001201011001', '06050101', '6e3890e5-4e89-4661-812d-b7f1987fa8d4'),
        ('85b53467-1bb7-4793-8e76-ca6b29b7658b', 1638279717616, 1638279717616, 0, 'Nguyễn Văn D', 1638279717616, '001201010012', 'male', 'Việt Nam', 'Hà Nội, Sóc Sơn, Nội Bài', 'Hà Nội, Sóc Sơn, Nội Bài', 'Hà Nội, Sóc Sơn, Nội Bài', 'Sinh viên', 'Không', '12/12', 'Nguyễn Ba C', '001201010001', 'Nguyễn Mẹ C', '001201010002', '06050101', '6e3890e5-4e89-4661-812d-b7f1987fa8d4'),
        ('bca6e2a3-777c-4f65-8a49-03c91efffa48', 1638279717616, 1638279717616, 0, 'Nguyễn Văn B', 1638279724474, '001201010020', 'male', 'Việt Nam', 'Hà Nội, Đông Anh, Thụy Lâm', 'Hà Nội, Đông Anh, Thụy Lâm', 'Hà Nội, Đông Anh, Thụy Lâm', 'Tài xế', 'Không', '12/12', 'Nguyễn Ba A', '001201010008', 'Nguyễn Mẹ A', '001201011002', '06050101', '6e3890e5-4e89-4661-812d-b7f1987fa8d4'),
        ('c2592f14-7413-4c87-bd82-7d08f49e6c53', 1638279717616, 1638279717616, 0, 'Nguyễn Thị B', 1638279724474, '001201010019', 'female', 'Việt Nam', 'Hà Nội, Đông Anh, Thị Trấn Đông Anh', 'Hà Nội, Đông Anh, Thị Trấn Đông Anh', 'Hà Nội, Đông Anh, Thị Trấn Đông Anh', 'Tài xế', 'Không', '12/12', 'Nguyễn Ba A', '001201010008', 'Nguyễn Mẹ A', '001201011002', '06050101', '6e3890e5-4e89-4661-812d-b7f1987fa8d4'),
        ('fd88e9a4-3d59-4443-a917-10e0a6b85537', 1638279717616, 1638279717616, 0, 'Nguyễn Thị E', 1638279717616, '001201010013', 'female', 'Việt Nam', 'Hà Nội, Cầu Giấy, Xuân Đỉnh', 'Hà Nội, Cầu Giấy, Xuân Đỉnh', 'Hà Nội, Cầu Giấy, Xuân Đỉnh', 'Sinh viên', 'Không', '12/12', 'Nguyễn Ba C', '001201010001', 'Nguyễn Mẹ C', '001201010002', '06050101', '6e3890e5-4e89-4661-812d-b7f1987fa8d4'),
        ('fe92bdb5-ce1a-4714-8611-bae249dad72c', 1638279724474, 1638279724474, 0, 'Nguyễn Trần F', 1638279724474, '001201010014', 'female', 'Việt Nam', 'Hà Nội, Cầu Giấy, Yên Hòa', 'Hà Nội, Cầu Giấy, Yên Hòa', 'Hà Nội, Cầu Giấy, Yên Hòa', 'Công Nhân', 'Không', '12/12', 'Nguyễn Ba B', '001201010009', 'Nguyễn Mẹ B', '001201011001', '06050101', '6e3890e5-4e89-4661-812d-b7f1987fa8d4');
`
