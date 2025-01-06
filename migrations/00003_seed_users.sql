-- +goose Up
INSERT INTO users
    (email, first_name, last_name, hashed_password, activated, verification_status, verification_document_url)
VALUES
    ('user01@gmail.com', 'John', 'Doe', '$2a$10$d4YogSzgFpWBE/G4SHY/NOBhljdX//IFeZKsk4cg0Eze42sQWMk9K', true, 'unverified', NULL),
    ('user02@gmail.com', 'Jane', 'Doe', '$2a$10$d4YogSzgFpWBE/G4SHY/NOBhljdX//IFeZKsk4cg0Eze42sQWMk9K', true, 'verified', 'https://www.dropbox.com/scl/fi/hmmrp6vystgzw87pphfzj/userdoc_1.pdf?rlkey=0j493u4rcmpsrwnjfymdug4v4&dl=0'),
    ('user03@gmail.com', 'Steven', 'Dubner', '$2a$10$d4YogSzgFpWBE/G4SHY/NOBhljdX//IFeZKsk4cg0Eze42sQWMk9K', true, 'unverified', NULL),
    ('user04@gmail.com', 'Jamal', 'Smith', '$2a$10$d4YogSzgFpWBE/G4SHY/NOBhljdX//IFeZKsk4cg0Eze42sQWMk9K', true, 'unverified', NULL);

INSERT INTO escrow_users
    (email, hashed_password)
VALUES
    ('escrow01@donorbox.com', '$2a$10$JLQE/YoItW3utQQRmL7C3eU7foPOfOdpIlOl14pt.Gp7kb4Tw.7re'),
    ('escrow02@donorbox.com', '$2a$10$JLQE/YoItW3utQQRmL7C3eU7foPOfOdpIlOl14pt.Gp7kb4Tw.7re');
-- +goose Down
--DELETE FROM users WHERE TRUE;
--DELETE FROM escrow_users WHERE TRUE;