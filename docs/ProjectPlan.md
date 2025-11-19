<a id="fr"></a>

<div align="center">
  <a href="#en">🇬🇧 English</a> · 
  <a href="#fr">🇫🇷 français</a>
</div>

# InfraMusicStore - Plan de projet 5 jours

**Équipe :** Leila Wilde, Louis Cordier, Mathieu Auger  
**Durée :** 5 jours  
**Date de début :** 2025-11-17 
**Date de fin :** 2025-11-21

---

## Aperçu du projet

Construire une API REST conteneurisée pour un magasin de disques en ligne avec base de données MySQL, entièrement déployée via Docker Compose avec automatisation CI/CD.

---

## Jour 1 : base de données et configuration du projet (2025-11-17)

### matin
- [ ] Initialiser le référentiel Git et GitHub Actions
- [ ] Créer le squelette Docker Compose avec 3 services (API, BD, Admin)
- [ ] Configurer `.env.example` et les fichiers de configuration

### après-midi
- [ ] Concevoir et créer le schéma de base de données MySQL (basé sur Chinook)
   - Tables Artistes, Albums, Pistes, Genres
   - Clés étrangères et relations
- [ ] Créer le DDL (scripts d'initialisation SQL)
- [ ] Générer les données de test par programmation (~20 entrées par 4 tables)

**Livrable :** Base de données en exécution dans Docker, remplie avec les données de test

---

## Jour 2 : cœur de l'API REST (2025-11-18)

### matin
- [ ] Configurer la structure du projet Go-Fiber
- [ ] Implémenter la connexion à la base de données avec les variables d'environnement
- [ ] Créer les points de terminaison CRUD pour la ressource Pistes
   - GET /api/tracks
   - GET /api/tracks/:id
   - POST /api/tracks
   - PUT /api/tracks/:id
   - DELETE /api/tracks/:id

### après-midi
- [ ] Implémenter la gestion des erreurs HTTP (403, 404, 500)
- [ ] Tester les points de terminaison Pistes localement
- [ ] Créer le Dockerfile pour le service API

**Livrable :** API CRUD Pistes fonctionnelle dans Docker

---

## Jour 3 : ressources supplémentaires et Swagger (2025-11-19)

### matin
- [ ] Implémenter CRUD pour 3 ressources supplémentaires :
   - Artistes
   - Albums
   - Genres

### après-midi
- [ ] Configurer l'intégration Swagger UI
- [ ] Créer swagger.yml/swagger.json avec tous les points de terminaison et exemples
- [ ] Tester tous les points de terminaison via Swagger UI

**Livrable :** 4 ressources CRUD complètes avec documentation Swagger

---

## Jour 4 : Docker Compose et CI/CD (2025-11-20)

### matin
- [ ] Terminer docker-compose.yml avec les 3 services :
   - API (Go-Fiber)
   - Base de données MySQL
   - Adminer (Interface Admin)
- [ ] Configurer les réseaux et volumes Docker
- [ ] Tester la pile complète localement : `docker-compose up`

### après-midi
- [ ] Configurer le flux de travail GitHub Actions pour CI/CD
   - Construction automatique des images Docker lors des envois
   - Déploiement automatique à l'environnement
- [ ] Créer `.env.example` avec toutes les variables requises
- [ ] Ajouter les commandes Docker à la documentation

**Livrable :** Application entièrement conteneurisée, prête à être déployée

---

## Jour 5 : documentation et derniers ajustements (2025-11-21)

### matin
- [ ] Compléter le README.md avec :
   - aperçu de l'architecture du projet
   - instructions d'installation
   - commandes Docker
   - URL des services API
   - schéma de la base de données (optionnel)
- [ ] Vérifier que tous les liens et exemples de documentation fonctionnent

### après-midi
- [ ] Test final : démarrer une nouvelle pile Docker, vérifier tous les points de terminaison
- [ ] Révision rapide du code et nettoyage
- [ ] Valider les modifications finales
- [ ] Pousser vers GitHub

**Livrable :** soumission complète du projet sur GitHub

---

## Critères de succès

- ✅ Base de données : 4 tables normalisées avec données de test
- ✅ API : toutes les opérations CRUD fonctionnelles pour 4+ ressources
- ✅ Swagger : documentation API complète
- ✅ Docker : pile de 3 services s'exécutant correctement
- ✅ CI/CD : flux de travail GitHub Actions fonctionnel
- ✅ Documentation : README.md complet et clair

---

## Priorités clés (par ordre)

1. **base de données** - doit être solide avant l'API
2. **API CRUD de base** - faire fonctionner les Pistes en premier, puis répliquer
3. **Docker Compose** - doit fonctionner de manière fiable pour le déploiement
4. **Documentation Swagger** - requise pour la spécification de l'API
5. **pipeline CI/CD** - automatise le déploiement
6. **derniers ajustements** - code propre et documentation

---

## Référence des commandes rapides

```bash
# Démarrer le développement
docker-compose up -d

# Afficher les journaux
docker-compose logs -f api

# Accès à la base de données
docker-compose exec db mysql -u root -p

# Arrêter tout
docker-compose down

# Reconstruction complète
docker-compose down && docker-compose up --build
```

---

## notes

- garder l'authentification simple ou l'ignorer si le temps est limité
- pagination/filtrage peut être ajouté après la date limite
- se concentrer sur la fonctionnalité principale plutôt que sur la perfection
- points de contrôle quotidiens pour rester sur la bonne voie

<a id="en"></a>

<div align="center">
  <a href="#en">🇬🇧 English</a> · 
  <a href="#fr">🇫🇷 français</a>
</div>

# InfraMusicStore - 5 day project plan

**Team:** Leila Wilde, Louis Cordier, Mathieu Auger  
**Duration:** 5 Days  
**Start Date:** 2025-11-17 
**End Date:** 2025-11-21

---

## Project overview

Build a containerized REST API for an online record store with MySQL database, all deployed via Docker Compose with CI/CD automation.

---

## Day 1: database & project setup (2025-11-17)

### morning
- [ ] Initialize Git repository and GitHub Actions
- [ ] Create Docker Compose skeleton with 3 services (API, DB, Admin)
- [ ] Set up `.env.example` and configuration files

### afternoon
- [ ] Design and create MySQL database schema (Chinook-based)
  - Artists, Albums, Tracks, Genres tables
  - Foreign keys and relationships
- [ ] Create DDL (SQL initialization scripts)
- [ ] Generate test data programmatically (~20 entries per 4 tables)

**deliverable:** Database running in Docker, populated with test data

---

## Day 2: REST API core (2025-11-18)

### morning
- [ ] Set up Go-Fiber project structure
- [ ] Implement database connection with environment variables
- [ ] Create CRUD endpoints for Tracks resource
  - GET /api/tracks
  - GET /api/tracks/:id
  - POST /api/tracks
  - PUT /api/tracks/:id
  - DELETE /api/tracks/:id

### afternoon
- [ ] implement HTTP error handling (403, 404, 500)
- [ ] Test Tracks endpoints locally
- [ ] create Dockerfile for API service

**Deliverable:** functional Tracks CRUD API in Docker

---

## Day 3: additional resources & Swagger (2025-11-19)

### morning
- [ ] implement CRUD for 3 more resources:
  - Artists
  - Albums
  - Genres

### afternoon
- [ ] set up Swagger UI integration
- [ ] create swagger.yml/swagger.json with all endpoints and examples
- [ ] test all endpoints via Swagger UI

**deliverable:** 4 complete CRUD resources with Swagger documentation

---

## Day 4: Docker compose & CI/CD (2025-11-20)

### morning
- [ ] complete docker-compose.yml with all 3 services:
  - API (Go-Fiber)
  - MySQL Database
  - Adminer (Admin UI)
- [ ] configure Docker networks and volumes
- [ ] test full stack locally: `docker-compose up`

### afternoon
- [ ] set up GitHub Actions workflow for CI/CD
  - Auto-build Docker images on push
  - Auto-deploy to environment
- [ ] create `.env.example` with all required variables
- [ ] add docker commands to documentation

**deliverable:** Fully containerized application, ready to deploy

---

## Day 5: documentation & final touches (2025-11-21)

### morning
- [ ] complete README.md with:
  - project architecture overview
  - installation instructions
  - Docker commands
  - API service URLs
  - database schema (optional)
- [ ] verify all documentation links and examples work

### afternoon
- [ ] final testing: start fresh Docker stack, verify all endpoints
- [ ] quick code review and cleanup
- [ ] commit final changes
- [ ] push to GitHub

**deliverable:** complete project submission on GitHub

---

## Success criteria

- ✅ Database: 4 normalized tables with test data
- ✅ API: all CRUD operations working for 4+ resources
- ✅ Swagger: complete API documentation
- ✅ Docker: 3-service stack running smoothly
- ✅ CI/CD: GitHub Actions workflow functional
- ✅ Documentation: README.md complete and clear

---

## Key priorities (in order)

1. **database** - must be solid before API
2. **core CRUD API** - get Tracks working first, then replicate
3. **Docker compose** - must work reliably for deployment
4. **Swagger docs** - required for API specification
5. **CI/CD pipeline** - automates deployment
6. **final touches** - clean code and documentation

---

## Quick commands reference

```bash
# Start development
docker-compose up -d

# View logs
docker-compose logs -f api

# Database access
docker-compose exec db mysql -u root -p

# Stop everything
docker-compose down

# Full rebuild
docker-compose down && docker-compose up --build
```

---

## notes

- keep authentication simple or skip if time is tight
- pagination/filtering can be added post-deadline
- focus on core functionality over perfection
- daily check-ins to stay on track
