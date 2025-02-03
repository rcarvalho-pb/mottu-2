-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
INSERT INTO tb_users
    (username, password, role, name, birth_date, cnpj, cnh, cnh_type, cnh_file_path)
VALUES
    ('rcarvalho', '$2a$10$g.Q7hSBL.1pftZv3RkE3nOu2nzfBuB9UuNnEMBnsS.VXg5XZFwMAW', 1, 'Ramon', '1994-07-22', '123123123123', '123123123123', 'A', 'TESTE');

INSERT INTO tb_users
    (username, password, name, birth_date, cnpj, cnh, cnh_type, cnh_file_path)
VALUES
    ('rcarvalhoo', '$2a$10$g.Q7hSBL.1pftZv3RkE3nOu2nzfBuB9UuNnEMBnsS.VXg5XZFwMAW', 'Ramon2', '1994-07-22', '1231231231232', '1231231231232', 'A', 'TESTE'),
    ('emillycs', '$2a$10$g.Q7hSBL.1pftZv3RkE3nOu2nzfBuB9UuNnEMBnsS.VXg5XZFwMAW', 'Emilly', '1993-09-23', '12312312312312', '12312312312312', 'B', 'TESTE');

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
