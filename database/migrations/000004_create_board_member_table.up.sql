-- Tabel untuk menyimpan relasi antara board dan user (many-to-many)
CREATE TABLE board_members (
    board_internal_id BIGINT NOT NULL REFERENCES boards(internal_id) ON DELETE CASCADE, -- ID board, hapus relasi jika board dihapus
    user_internal_id BIGINT NOT NULL REFERENCES users(internal_id) ON DELETE CASCADE,   -- ID user, hapus relasi jika user dihapus
    joined_at TIMESTAMP NOT NULL DEFAULT NOW(),                                         -- Waktu user bergabung ke board
    PRIMARY KEY (board_internal_id, user_internal_id)                                   -- Satu user hanya bisa sekali di tiap board
);