<a id="en"></a>

<div align="center">
  <a href="#en">🇬🇧 English</a> · 
  <a href="#fr">🇫🇷 français</a>
</div>

# InfraMusicStore API

A REST API for managing a music store database. Built with **Go-Fiber**, **MySQL**, and **Docker**.

## Project overview

This API provides complete CRUD operations for a music database with the following resources:
- **Genres** - music genres
- **Artists** - music artists/bands
- **Albums** - music albums  
- **Tracks** - individual songs

All data is persisted in MySQL and fully containerized with Docker.

## Architecture

```
┌────────────────────────────────────────────────┐
│         Docker Compose (3 Services)            │
├──────────────┬──────────────┬──────────────────┤
│  Go-Fiber    │   MySQL 8.0  │   Adminer        │
│  API         │   Database   │   (Admin UI)     │
│  :8080       │   :3306      │   :8081          │
└──────────────┴──────────────┴──────────────────┘
```

## 🚀 Quick start

### Prerequisites
- Docker & docker compose installed

### 1. Start the stack
```bash
docker-compose up -d
```

### 2. Wait for services
```bash
# check status
docker-compose ps

# view logs
docker-compose logs -f api
```

### 3. Test the API
```bash
# health check
curl http://localhost:8080/api/health

# get all genres
curl http://localhost:8080/api/genres

# get all tracks
curl http://localhost:8080/api/tracks
```

### 4. Access services
- **API**: http://localhost:8080
- **Adminer (database UI)**: http://localhost:8081
  - Server: `db`
  - Username: `inframusic_user`
  - Password: `inframusic_pass`
  - Database: `inframusic_db`

## API endpoints

### Genres
- `GET /api/genres` - list all genres
- `GET /api/genres/:id` - get genre by ID
- `POST /api/genres` - create new genre
- `PUT /api/genres/:id` - update genre
- `DELETE /api/genres/:id` - delete genre

### Artists
- `GET /api/artists` - list all artists
- `GET /api/artists/:id` - get artist by ID
- `POST /api/artists` - create new artist
- `PUT /api/artists/:id` - update artist
- `DELETE /api/artists/:id` - delete artist

### Albums
- `GET /api/albums` - list all albums
- `GET /api/albums/:id` - get album by ID
- `POST /api/albums` - create new album
- `PUT /api/albums/:id` - update album
- `DELETE /api/albums/:id` - delete album

### Tracks
- `GET /api/tracks` - list all tracks
- `GET /api/tracks/:id` - get track by ID
- `POST /api/tracks` - create new track
- `PUT /api/tracks/:id` - update track
- `DELETE /api/tracks/:id` - delete track

## Example requests

### Create a genre
```bash
curl -X POST http://localhost:8080/api/genres \
  -H "Content-Type: application/json" \
  -d '{"name": "Electronic"}'
```

### Create an Artist
```bash
curl -X POST http://localhost:8080/api/artists \
  -H "Content-Type: application/json" \
  -d '{"name": "Daft Punk"}'
```

### Create an Album
```bash
curl -X POST http://localhost:8080/api/albums \
  -H "Content-Type: application/json" \
  -d '{"title": "Random Access Memories", "artist_id": 1}'
```

### Create a Track
```bash
curl -X POST http://localhost:8080/api/tracks \
  -H "Content-Type: application/json" \
  -d '{"name": "Get Lucky", "album_id": 1, "genre_id": 1}'
```

## Docker Commands

### View Logs
```bash
# all services
docker-compose logs -f

# specific service
docker-compose logs -f api
docker-compose logs -f db
```

### Stop services
```bash
docker-compose down
```

### Rebuild images
```bash
docker-compose up -d --build
```

### Reset everything (including data)
```bash
docker-compose down -v
docker-compose up -d
```

### Access database shell
```bash
docker-compose exec db mysql -u inframusic_user -p inframusic_db
# password: inframusic_pass
```

## Database configuration

The database comes pre-loaded with:
- **25 genres**
- **275 artists**
- **347 albums**  
- **4100+ tracks**

Configuration via `.env`:
```
DB_HOST=db
DB_PORT=3306
DB_USER=inframusic_user
DB_PASSWORD=inframusic_pass
DB_NAME=inframusic_db
API_PORT=8080
```

## 📂 Project structure

```
├── src/                     # go source files
│   ├── main.go              # application entry point & routing
│   ├── models.go            # data structures (Genres, Artists, Albums, Tracks)
│   ├── db.go                # database connection & queries
│   └── handlers.go          # HTTP request handlers
├── db/                      # database scripts
│   ├── 01_init_struct.sql   # schema (tables & relationships)
│   └── 02_init_seed.sql     # test data (4100+ tracks)
├── Dockerfile               # go API container image
├── docker-compose.yml       # service orchestration
├── go.mod                   # go dependencies
└── README.md                # this file
```

## Understanding the code

**main.go** - application startup
- sets up Fiber web framework
- loads environment variables
- initializes database connection
- defines HTTP routes (GET, POST, PUT, DELETE)

**models.go** - data definitions
- struct definitions for Genres, Artists, Albums, Tracks
- JSON serialization tags for API responses

**db.go** - database operations
- database connection pooling
- CRUD functions for each resource
- SQL queries using parameterized statements

**handlers.go** - API request handling
- HTTP status codes (200, 201, 400, 404, 500)
- request body parsing (JSON)
- error handling and validation

### Go basics
- **Goroutines**: `go` keyword for concurrent operations
- **Channels**: communication between goroutines
- **Interfaces**: flexible type system (like duck typing)
- **Defer**: cleanup operations (like try/finally)
- **Error handling**: multiple return values

## 🐛 Troubleshooting

### API won't connect to database
- wait 10-15 seconds for MySQL to fully start
- check logs: `docker-compose logs api`
- verify network: `docker network ls`

### port already in use
- change ports in `docker-compose.yml` or `.env`
- or stop other services: `docker ps` and `docker stop <container>`

### database reset
```bash
docker-compose down -v  # removes volumes
docker-compose up -d    # fresh database with seed data
```

## ✅ success checklist

- ✅ Docker containers running without errors
- ✅ API responds to `/api/health`
- ✅ can retrieve data from `/api/genres`, `/api/artists`, etc.
- ✅ can create new records via POST requests
- ✅ Adminer accessible at http://localhost:8081

## MIT License

This is a student project for educational purposes created by Leila Wilde, Louis Cordier & Mathieu Auger at La Plateforme_ school in Marseille.

<a id="fr"></a>

<div align="center">
  <a href="#en">🇬🇧 English</a> · 
  <a href="#fr">🇫🇷 français</a>
</div>

# API InfraMusicStore

Une API REST pour gérer une base de données de magasin de musique. Construite avec **Go-Fiber**, **MySQL** et **Docker**.

## Aperçu du projet

Cette API fournit des opérations CRUD complètes pour une base de données musicale avec les ressources suivantes :
- **Genres** - genres musicaux
- **Artistes** - artistes/groupes musicaux
- **Albums** - albums musicaux
- **Pistes** - chansons individuelles

Toutes les données sont persistantes dans MySQL et entièrement conteneurisées avec Docker.

## Architecture

```
┌────────────────────────────────────────────────┐
│         Docker Compose (3 Services)            │
├──────────────┬──────────────┬──────────────────┤
│  Go-Fiber    │   MySQL 8.0  │   Adminer        │
│  API         │   Base de    │   (Interface     │
│  :8080       │   données    │    Admin)        │
│              │   :3306      │   :8081          │
└──────────────┴──────────────┴──────────────────┘
```

## 🚀 Démarrage rapide

### Prérequis
- Docker & docker compose installés

### 1. Démarrer la pile
```bash
docker-compose up -d
```

### 2. Attendre les services
```bash
# vérifier le statut
docker-compose ps

# voir les logs
docker-compose logs -f api
```

### 3. Tester l'API
```bash
# vérification de santé
curl http://localhost:8080/api/health

# obtenir tous les genres
curl http://localhost:8080/api/genres

# obtenir toutes les pistes
curl http://localhost:8080/api/tracks
```

### 4. Accéder aux services
- **API**: http://localhost:8080
- **Adminer (interface base de données)**: http://localhost:8081
  - Serveur: `db`
  - Nom d'utilisateur: `inframusic_user`
  - Mot de passe: `inframusic_pass`
  - Base de données: `inframusic_db`

## Points de terminaison API

### Genres
- `GET /api/genres` - lister tous les genres
- `GET /api/genres/:id` - obtenir un genre par ID
- `POST /api/genres` - créer un nouveau genre
- `PUT /api/genres/:id` - mettre à jour un genre
- `DELETE /api/genres/:id` - supprimer un genre

### Artistes
- `GET /api/artists` - lister tous les artistes
- `GET /api/artists/:id` - obtenir un artiste par ID
- `POST /api/artists` - créer un nouvel artiste
- `PUT /api/artists/:id` - mettre à jour un artiste
- `DELETE /api/artists/:id` - supprimer un artiste

### Albums
- `GET /api/albums` - lister tous les albums
- `GET /api/albums/:id` - obtenir un album par ID
- `POST /api/albums` - créer un nouvel album
- `PUT /api/albums/:id` - mettre à jour un album
- `DELETE /api/albums/:id` - supprimer un album

### Pistes
- `GET /api/tracks` - lister toutes les pistes
- `GET /api/tracks/:id` - obtenir une piste par ID
- `POST /api/tracks` - créer une nouvelle piste
- `PUT /api/tracks/:id` - mettre à jour une piste
- `DELETE /api/tracks/:id` - supprimer une piste

## Exemples de requêtes

### Créer un genre
```bash
curl -X POST http://localhost:8080/api/genres \
  -H "Content-Type: application/json" \
  -d '{"name": "Electronique"}'
```

### Créer un artiste
```bash
curl -X POST http://localhost:8080/api/artists \
  -H "Content-Type: application/json" \
  -d '{"name": "Daft Punk"}'
```

### Créer un album
```bash
curl -X POST http://localhost:8080/api/albums \
  -H "Content-Type: application/json" \
  -d '{"title": "Random Access Memories", "artist_id": 1}'
```

### Créer une piste
```bash
curl -X POST http://localhost:8080/api/tracks \
  -H "Content-Type: application/json" \
  -d '{"name": "Get Lucky", "album_id": 1, "genre_id": 1}'
```

## Commandes Docker

### Afficher les logs
```bash
# tous les services
docker-compose logs -f

# service spécifique
docker-compose logs -f api
docker-compose logs -f db
```

### Arrêter les services
```bash
docker-compose down
```

### Reconstruire les images
```bash
docker-compose up -d --build
```

### Réinitialiser tout (y compris les données)
```bash
docker-compose down -v
docker-compose up -d
```

### Accéder au shell de la base de données
```bash
docker-compose exec db mysql -u inframusic_user -p inframusic_db
# mot de passe: inframusic_pass
```

## Configuration de la base de données

La base de données est pré-chargée avec :
- **25 genres**
- **275 artistes**
- **347 albums**
- **4100+ pistes**

Configuration via `.env`:
```
DB_HOST=db
DB_PORT=3306
DB_USER=inframusic_user
DB_PASSWORD=inframusic_pass
DB_NAME=inframusic_db
API_PORT=8080
```

## 📂 Structure du projet

```
├── src/                     # fichiers source go
│   ├── main.go              # point d'entrée de l'application & routage
│   ├── models.go            # structures de données (Genres, Artistes, Albums, Pistes)
│   ├── db.go                # connexion à la base de données & requêtes
│   └── handlers.go          # gestionnaires de requêtes HTTP
├── db/                      # scripts de base de données
│   ├── 01_init_struct.sql   # schéma (tableaux & relations)
│   └── 02_init_seed.sql     # données de test (4100+ pistes)
├── Dockerfile               # image du conteneur API go
├── docker-compose.yml       # orchestration des services
├── go.mod                   # dépendances go
└── README.md                # ce fichier
```

## Comprendre le code

**main.go** - démarrage de l'application
- configure le framework web Fiber
- charge les variables d'environnement
- initialise la connexion à la base de données
- définit les routes HTTP (GET, POST, PUT, DELETE)

**models.go** - définitions de données
- définitions de struct pour Genres, Artistes, Albums, Pistes
- étiquettes de sérialisation JSON pour les réponses API

**db.go** - opérations de base de données
- pool de connexions de base de données
- fonctions CRUD pour chaque ressource
- requêtes SQL utilisant des déclarations paramétrées

**handlers.go** - gestion des requêtes API
- codes de statut HTTP (200, 201, 400, 404, 500)
- analyse du corps de la requête (JSON)
- gestion des erreurs et validation

### Concepts Go
- **Goroutines**: mot-clé `go` pour les opérations concurrentes
- **Canaux**: communication entre goroutines
- **Interfaces**: système de type flexible (comme le duck typing)
- **Defer**: opérations de nettoyage (comme try/finally)
- **Gestion d'erreurs**: valeurs de retour multiples

## 🐛 Dépannage

### L'API ne peut pas se connecter à la base de données
- attendre 10-15 secondes pour que MySQL démarre complètement
- vérifier les logs: `docker-compose logs api`
- vérifier le réseau: `docker network ls`

### Port déjà utilisé
- changer les ports dans `docker-compose.yml` ou `.env`
- ou arrêter d'autres services: `docker ps` et `docker stop <conteneur>`

### Réinitialiser la base de données
```bash
docker-compose down -v  # supprime les volumes
docker-compose up -d    # nouvelle base de données avec données de semence
```

## ✅ Liste de contrôle de succès

- ✅ Conteneurs Docker fonctionnant sans erreurs
- ✅ L'API répond à `/api/health`
- ✅ Peut récupérer les données de `/api/genres`, `/api/artists`, etc.
- ✅ Peut créer de nouveaux enregistrements via des requêtes POST
- ✅ Adminer accessible à http://localhost:8081

## Licence MIT

Ceci est un projet étudiant à des fins éducatives créé par Leila Wilde, Louis Cordier & Mathieu Auger à l'école La Plateforme_ à Marseille.