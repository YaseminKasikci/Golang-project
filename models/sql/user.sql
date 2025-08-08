-- INNER JOIN: element existant sur les 2 tableaux
-- SELECT * FROM users
--     INNER JOIN sessions ON users.id = sessions.user_id

-- INNER LEFT: permet de lister tous les résultats de la table de gauche (left = gauche) même s’il n’y a pas de correspondance dans la deuxième tables.
-- SELECT * FROM users
--     LEFT JOIN sessions ON users.id = sessions.user_id

-- INNER RIGHT: permet de retourner tous les enregistrements de la table de droite (right = droite) même s’il n’y a pas de correspondance avec la table de gauche. 
-- SELECT * FROM users
--     RIGHT JOIN sessions ON users.id = sessions.user_id

--
-- INNER OUTER / FULL: opération qui renverra tous les enregistrements des deux tables, qu'il existe ou non un enregistrement lié
-- SELECT * FROM users
--     FULL OUTER JOIN sessions ON users.id = sessions.user_id

SELECT * FROM users
    LEFT JOIN sessions ON users.id = sessions.user_id

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);