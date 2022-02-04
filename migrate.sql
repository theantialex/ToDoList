CREATE TABLE Tasks (
    id SERIAL NOT NULL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    done BOOLEAN DEFAULT false 
);