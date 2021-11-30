package citizen

var SQL = `insert into public.citizens (id, updated_at, created_at, deleted_at, name, birthday, pid, gender, nationality, father_name, father_pid, mother_name, mother_pid, current_place, job_name)
values  ('0203bec7-0e51-4abd-94db-f965fbdec66a', 1638279717616, 1638279717616, 0, 'Nguyễn Văn C', 1638279717616, '001201010023', 'male', 'Việt Nam', 'Nguyễn Ba C', '001201010001', 'Nguyễn Mẹ C', '001201010002', 'Hà Nội, Đông Anh, Nguyên Khê', 'Sinh viên'),
        ('1abebf8c-3af9-47a1-8907-40dc49d75da1', 1638279717616, 1638279717616, 0, 'Nguyễn Duy C', 1638279724474, '001201010015', 'male', 'Việt Nam', 'Nguyễn Ba A', '001201010008', 'Nguyễn Mẹ A', '001201011002', 'Hà Nội, Đông Anh, Liên Hà', 'Tài xế'),
        ('1d0bc669-462e-4479-95a6-18da17d47969', 1638279717616, 1638279717616, 0, 'Nguyễn Thanh K', 1638279724474, '001201010016', 'female', 'Việt Nam', 'Nguyễn Ba A', '001201010008', 'Nguyễn Mẹ A', '001201011002', 'Hà Nội, Đông Anh, Thị Trấn Đông Anh', 'Tài xế'),
        ('297a2206-72a6-40d5-84a2-114b896dabca', 1638279717616, 1638279717616, 0, 'Nguyễn Thị C', 1638279717616, '001201010022', 'female', 'Việt Nam', 'Nguyễn Ba C', '001201010001', 'Nguyễn Mẹ C', '001201010002', 'Hà Nội, Đông Anh, Dục Tú', 'Sinh viên'),
        ('7f1cd91c-c050-4ad1-ba95-d0394428242e', 1638279724474, 1638279724474, 0, 'Nguyễn Minh H', 1638279724474, '001201010017', 'male', 'Việt Nam', 'Nguyễn Ba B', '001201010009', 'Nguyễn Mẹ B', '001201011001', 'Hà Nội, Đông Anh, Tàm Xá', 'Công Nhân'),
        ('85b53467-1bb7-4793-8e76-ca6b29b7658b', 1638279717616, 1638279717616, 0, 'Nguyễn Văn D', 1638279717616, '001201010012', 'male', 'Việt Nam', 'Nguyễn Ba C', '001201010001', 'Nguyễn Mẹ C', '001201010002', 'Hà Nội, Sóc Sơn, Nội Bài', 'Sinh viên'),
        ('bbe5245b-eb70-42bc-ac5a-4dca05c2047e', 1638279724474, 1638279724474, 0, 'Nguyên Thị A', 1638279724474, '001201010021', 'female', 'Việt Nam', 'Nguyễn Ba B', '001201010009', 'Nguyễn Mẹ B', '001201011001', 'Hà Nội, Đông Anh, Xuân Nộn', 'Công Nhân'),
        ('bca6e2a3-777c-4f65-8a49-03c91efffa48', 1638279717616, 1638279717616, 0, 'Nguyễn Văn B', 1638279724474, '001201010020', 'male', 'Việt Nam', 'Nguyễn Ba A', '001201010008', 'Nguyễn Mẹ A', '001201011002', 'Hà Nội, Đông Anh, Thụy Lâm', 'Tài xế'),
        ('c2592f14-7413-4c87-bd82-7d08f49e6c53', 1638279717616, 1638279717616, 0, 'Nguyễn Thị B', 1638279724474, '001201010019', 'female', 'Việt Nam', 'Nguyễn Ba A', '001201010008', 'Nguyễn Mẹ A', '001201011002', 'Hà Nội, Đông Anh, Thị Trấn Đông Anh', 'Tài xế'),
        ('d62cc8d4-9cec-41d1-b700-a87ee14b7e3e', 1638279724474, 1638279724474, 0, 'Nguyên Văn A', 1638279724474, '001201010018', 'male', 'Việt Nam', 'Nguyễn Ba B', '001201010009', 'Nguyễn Mẹ B', '001201011001', 'Hà Nội, Cầu Giấy, Yên Hòa', 'Công Nhân'),
        ('fd88e9a4-3d59-4443-a917-10e0a6b85537', 1638279717616, 1638279717616, 0, 'Nguyễn Thị E', 1638279717616, '001201010013', 'female', 'Việt Nam', 'Nguyễn Ba C', '001201010001', 'Nguyễn Mẹ C', '001201010002', 'Hà Nội, Cầu Giấy, Xuân Đỉnh', 'Sinh viên'),
        ('fe92bdb5-ce1a-4714-8611-bae249dad72c', 1638279724474, 1638279724474, 0, 'Nguyễn Trần F', 1638279724474, '001201010014', 'female', 'Việt Nam', 'Nguyễn Ba B', '001201010009', 'Nguyễn Mẹ B', '001201011001', 'Hà Nội, Cầu Giấy, Yên Hòa', 'Công Nhân');`
