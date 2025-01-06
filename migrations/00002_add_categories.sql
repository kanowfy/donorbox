-- +goose Up
INSERT INTO categories(name, description, cover_picture) VALUES
    ('medical', 'Help funding for medical causes', 'https://images.pexels.com/photos/40568/medical-appointment-doctor-healthcare-40568.jpeg'),
    ('emergency', 'Help those with urgent needs', 'https://images.pexels.com/photos/263402/pexels-photo-263402.jpeg'),
    ('education', 'Help funding for the future', 'https://images.pexels.com/photos/247819/pexels-photo-247819.jpeg'),
    ('animals', 'Help with animal causes', 'https://images.pexels.com/photos/1299391/pexels-photo-1299391.jpeg'),
    ('competition', 'Help funding for competition', 'https://images.pexels.com/photos/3755440/pexels-photo-3755440.jpeg'),
    ('event', 'Help raising for events', 'https://images.pexels.com/photos/433452/pexels-photo-433452.jpeg'),
    ('environment', 'Raise funds for environmental causes', 'https://images.pexels.com/photos/1002703/pexels-photo-1002703.jpeg'),
    ('travel', 'Help with travelling expenses', 'https://images.pexels.com/photos/1271619/pexels-photo-1271619.jpeg'),
    ('business', 'Raise funds for business causes', 'https://images.pexels.com/photos/265087/pexels-photo-265087.jpeg');

-- +goose Down
--DELETE FROM categories WHERE TRUE;