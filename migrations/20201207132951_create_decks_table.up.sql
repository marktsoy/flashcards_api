CREATE TABLE decks (
    id  CHAR(32) PRIMARY KEY,
    user_id INTEGER  NOT NULL ,
    name VARCHAR (255) NOT NULL,

    FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE
)