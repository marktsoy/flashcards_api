CREATE TABLE cards(
    id SERIAL PRIMARY KEY,
    question TEXT,
    answer TEXT,
    
    created_at timestamp NOT NULL,
    
    deck_id CHAR(32) NOT NULL,


    FOREIGN KEY (deck_id)
        REFERENCES decks (id)
        ON DELETE CASCADE
)