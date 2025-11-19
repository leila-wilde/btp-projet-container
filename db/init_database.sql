-- ============================================
-- Script de création de la base de données
-- Magasin de musique en ligne
-- ============================================

-- Utiliser la base de données
USE music_store;

-- Supprimer les tables si elles existent déjà (pour pouvoir relancer le script)
DROP TABLE IF EXISTS lignes_commande;
DROP TABLE IF EXISTS commandes;
DROP TABLE IF EXISTS clients;
DROP TABLE IF EXISTS produits;
DROP TABLE IF EXISTS categories;

-- ============================================
-- Table : categories
-- ============================================
CREATE TABLE categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nom VARCHAR(100) NOT NULL,
    description TEXT,
    UNIQUE KEY unique_nom (nom)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Table : produits
-- ============================================
CREATE TABLE produits (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nom VARCHAR(200) NOT NULL,
    description TEXT,
    prix DECIMAL(10, 2) NOT NULL CHECK (prix >= 0),
    quantite_stock INT NOT NULL DEFAULT 0 CHECK (quantite_stock >= 0),
    categorie_id INT NOT NULL,
    image_url VARCHAR(500),
    date_ajout DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (categorie_id) REFERENCES categories(id) ON DELETE RESTRICT,
    INDEX idx_categorie (categorie_id),
    INDEX idx_nom (nom)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Table : clients
-- ============================================
CREATE TABLE clients (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nom VARCHAR(100) NOT NULL,
    prenom VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL,
    telephone VARCHAR(20),
    adresse TEXT NOT NULL,
    date_inscription DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY unique_email (email),
    INDEX idx_nom_prenom (nom, prenom)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Table : commandes
-- ============================================
CREATE TABLE commandes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    client_id INT NOT NULL,
    date_commande DATETIME DEFAULT CURRENT_TIMESTAMP,
    montant_total DECIMAL(10, 2) NOT NULL CHECK (montant_total >= 0),
    adresse_livraison TEXT NOT NULL,
    FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE RESTRICT,
    INDEX idx_client (client_id),
    INDEX idx_date (date_commande)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Table : lignes_commande
-- ============================================
CREATE TABLE lignes_commande (
    id INT AUTO_INCREMENT PRIMARY KEY,
    commande_id INT NOT NULL,
    produit_id INT NOT NULL,
    quantite INT NOT NULL CHECK (quantite > 0),
    prix_unitaire DECIMAL(10, 2) NOT NULL CHECK (prix_unitaire >= 0),
    FOREIGN KEY (commande_id) REFERENCES commandes(id) ON DELETE CASCADE,
    FOREIGN KEY (produit_id) REFERENCES produits(id) ON DELETE RESTRICT,
    INDEX idx_commande (commande_id),
    INDEX idx_produit (produit_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ============================================
-- Données de test : Categories
-- ============================================
INSERT INTO categories (nom, description) VALUES
('Instruments', 'Instruments de musique de tous types'),
('Vinyles', 'Disques vinyles et collectors'),
('CD', 'Compact Discs audio'),
('Accessoires', 'Accessoires pour musiciens');

-- ============================================
-- Données de test : Produits
-- ============================================
INSERT INTO produits (nom, description, prix, quantite_stock, categorie_id) VALUES
-- Instruments
('Guitare Fender Stratocaster', 'Guitare électrique légendaire, corps en aulne, manche en érable', 1299.99, 5, 1),
('Piano Yamaha P-45', 'Piano numérique 88 touches, idéal débutants', 499.99, 10, 1),
('Batterie Pearl Export', 'Batterie acoustique complète 5 fûts avec cymbales', 899.99, 3, 1),
('Saxophone Alto Selmer', 'Saxophone professionnel en mi bémol', 2499.99, 2, 1),

-- Vinyles
('Pink Floyd - The Dark Side of the Moon', 'Album vinyle remasterisé 180g', 29.99, 50, 2),
('The Beatles - Abbey Road', 'Vinyle original pressage 2019', 34.99, 30, 2),
('Daft Punk - Random Access Memories', 'Double vinyle 33 tours', 39.99, 20, 2),

-- CD
('Radiohead - OK Computer', 'CD album édition remasterisée', 14.99, 100, 3),
('Miles Davis - Kind of Blue', 'CD jazz classique', 12.99, 75, 3),

-- Accessoires
('Câble Jack 3m', 'Câble instrument blindé haute qualité', 19.99, 200, 4),
('Médiators Dunlop Pack 12', 'Pack de 12 médiators variés', 5.99, 500, 4),
('Pied de microphone Shure', 'Support micro réglable professionnel', 89.99, 25, 4),
('Accordeur chromatique Korg', 'Accordeur précis tous instruments', 24.99, 80, 4);

-- ============================================
-- Données de test : Clients
-- ============================================
INSERT INTO clients (nom, prenom, email, telephone, adresse) VALUES
('Dupont', 'Jean', 'jean.dupont@email.com', '0601020304', '12 Rue de la Musique, 75001 Paris'),
('Martin', 'Sophie', 'sophie.martin@email.com', '0612345678', '45 Avenue des Arts, 69002 Lyon'),
('Bernard', 'Luc', 'luc.bernard@email.com', '0698765432', '8 Boulevard du Rock, 13001 Marseille');

-- ============================================
-- Données de test : Commandes et lignes_commande
-- ============================================

-- Commande 1 : Jean Dupont achète une guitare et un câble
INSERT INTO commandes (client_id, montant_total, adresse_livraison) VALUES
(1, 1319.98, '12 Rue de la Musique, 75001 Paris');

INSERT INTO lignes_commande (commande_id, produit_id, quantite, prix_unitaire) VALUES
(1, 1, 1, 1299.99),  -- Guitare Fender
(1, 10, 1, 19.99);   -- Câble Jack

-- Commande 2 : Sophie Martin achète 2 vinyles
INSERT INTO commandes (client_id, montant_total, adresse_livraison) VALUES
(2, 64.98, '45 Avenue des Arts, 69002 Lyon');

INSERT INTO lignes_commande (commande_id, produit_id, quantite, prix_unitaire) VALUES
(2, 5, 1, 29.99),    -- Pink Floyd
(2, 6, 1, 34.99);    -- The Beatles

-- Commande 3 : Luc Bernard achète un piano et des accessoires
INSERT INTO commandes (client_id, montant_total, adresse_livraison) VALUES
(3, 614.97, '8 Boulevard du Rock, 13001 Marseille');

INSERT INTO lignes_commande (commande_id, produit_id, quantite, prix_unitaire) VALUES
(3, 2, 1, 499.99),   -- Piano Yamaha
(3, 11, 10, 5.99),   -- Pack médiators x10
(3, 13, 1, 24.99);   -- Accordeur

-- ============================================
-- Vérification : Afficher le contenu des tables
-- ============================================
SELECT '=== CATEGORIES ===' AS '';
SELECT * FROM categories;

SELECT '=== PRODUITS ===' AS '';
SELECT p.id, p.nom, p.prix, p.quantite_stock, c.nom as categorie 
FROM produits p 
JOIN categories c ON p.categorie_id = c.id;

SELECT '=== CLIENTS ===' AS '';
SELECT * FROM clients;

SELECT '=== COMMANDES ===' AS '';
SELECT cmd.id, CONCAT(cl.prenom, ' ', cl.nom) as client, cmd.date_commande, cmd.montant_total
FROM commandes cmd
JOIN clients cl ON cmd.client_id = cl.id;

SELECT '=== LIGNES COMMANDE ===' AS '';
SELECT lc.id, lc.commande_id, p.nom as produit, lc.quantite, lc.prix_unitaire
FROM lignes_commande lc
JOIN produits p ON lc.produit_id = p.id;