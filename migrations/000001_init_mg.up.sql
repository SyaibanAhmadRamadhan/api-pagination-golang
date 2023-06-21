CREATE TABLE profiles(
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    gender ENUM('male', 'female'),
    phone VARCHAR(15) NOT NULL,
    email VARCHAR(100) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);