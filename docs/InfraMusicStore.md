<a id="fr"></a>

<div align="center">
  <a href="#en">🇬🇧 English</a> · 
  <a href="#fr">🇫🇷 français</a>
</div>

# Bachelor TP - InfraMusicStore

Développer une API pour un disquaire en ligne, avec une stack entièrement containerisée

## Introduction du sujet

Ce projet d’une semaine a vocation à développer l’API de gestion d'un disquaire en ligne, dans un environnement Docker complet.

Vous allez construire un outil de gestion pour un magasin de musique en ligne.

## Mise en situation professionnelle 

Vous rejoignez l’équipe numérique du magasin **InfraMusicStore**, qui va diffuser de la musique en ligne, et promouvoir des artistes.

En fonction de votre parcours (Web, Logiciel, Data/IA, …) vous maîtrisez certaines techniques et langages de programmation (JS, C++, Python, SQL …), et vous êtes capables de développer le même projet avec des méthodologies différentes.

Vous livrerez votre projet par équipe de 2 à 3 personnes.

Votre manager vous demande d’implémenter une solution API, en construisant au préalable une base de données qui sera structurée comme
la base Chinook [disponible en ligne](https://github.com/lerocha/chinook-database).

## Étapes du projet

On ne vous demande pas la réalisation d’un site de gestion complet, ce qui serait trop long, mais simplement de **mettre en place une API REST qui permette d’accéder aux données de votre base relationnelle Chinook**, en lecture/écriture : toutes les opérations CRUD doivent être supportées.

Vous devrez donc déployer une API mettant à disposition les données de votre base de données, dans un environnement entièrement dockerisé, et qui utilise un **pipeline CI/CD automatisé**.

### Étape 1 : la base de données 

Vous devez mettre en place une **base de données relationnelle comportant plusieurs tables interconnectées** : artistes, albums, pistes, genres… ces tables sont fournies dans le dépôt Chinook. Ce sera à vous de les importer dans votre système et de vous assurer que les liaisons sont bien faites.

![ER diagram](./tables.png "illustration of tables")

Les attentes sont les suivantes :

- Définition claire des **relations et clés étrangères**,
- Fourniture des scripts SQL d’initialisation (DDL),
- La DB recommandée est **MariaDB**, mais d’autres moteurs de BDD relationnelles sont acceptés dans la mesure où ils sont Open Source.
- Vous devez vous assurer que vos tables sont bien **normalisées**.
- Vous devez intégrer des données de test qui ne sont pas dans les tables de base (au minimum 20 entrées pour 4 tables). Ces données devront être ajoutées de façon **programmatique** : vous ne pouvez pas
faire d’ajout manuel via un GUI !
    - Les étudiants des parcours Web et Logiciel peuvent s’inspirer d’un jeu de données [fourni](https://github.com/lerocha/chinook-database/releases) en JSON ou en SQL avec le dépôt Chinook. 
    Vous pouvez vous inspirer de ces exemples pour intégrer vos données de test.
    - Les étudiants en Data/IA sont invités à intégrer le dataset fourni avec le dépôt Chinook, et à produire un notebook Jupyter d’Analyse Exploratoire des Données.

### Étape 2 : API REST

L’application doit exposer une **API REST** complète, connectée à la base de données via des **variables d’environnement**.

Les endpoints minimums attendus sont ceux d’un CRUD. Par exemple, pour les tracks :

- **GET** /api/tracks — toutes les tracks
- **GET** /api/tracks/:id —détail d’une track
- **POST** /api/tracks — crée une track
- **PUT** /api/tracks/:id — modifie une track
- **DELETE** /api/tracks/:id — supprime une track.

Ce schéma est à avoir pour **au moins 4 ressources**.

Autres exigences :

- Toutes vos réponses doivent être au format **JSON**.
- Optionnel : mise en place d’un système **d’authentification** : le client de l’API doit être authentifié, via le système de votre choix (faites simple !). Si le client n’est pas authentifié, votre API devra retourner **le bon code HTTP**.
- Vous devez gérer correctement les **erreurs HTTP basiques** (403, 404, 500, etc.).
- Pour votre API, vous pouvez utilisez les technos suivantes : **Web** -> Express / [Node.js](http://node.js), **Logiciel** -> Go-Fiber, **Data/IA** -> Flask
- Optionnel : gestion de la **pagination**, des **filtres**, du **tri**.

### Étape 3 : Architecture Docker

Maintenant que votre API est fonctionnelle, vous pouvez commencer la conteneurisation. L’environnement de votre application doit être entièrement conteneurisé avec **Docker**. Vous devez utiliser **Docker Compose** avec au minimum **trois services distincts** :

1. L’**application principale** (votre API)
2. La **base de données**
3. Un outil d’administration de la base de données (Adminer, phpMyAdmin…)
4. (optionnel) Swagger UI pour la documentation interactive de l’API.

Les points suivants doivent être présents :

- Un **Dockerfile** spécifique à l’API,
- Une configuration **Docker Compose** complète avec les trois services,
- Une gestion correcte des **réseaux Docker** pour la communication entre
conteneurs,
- Optionnel : une configuration des **volumes** pour assurer la **persistance des données** de la base

### Étape 4 : Pipeline GIT + CI/CD

Une fois votre application entièrement dockerisée, on attend de vous que vous mettiez en place une véritable **intégration continue** de votre projet. En gros, le déploiement de votre application doit se faire de façon automatisée !

La gestion de version doivent être assurées via **GitHub** :

- Vous devez utiliser le système de **Pull Requests**,
- Vous devez mettre en place un **workflow GitHub Actions** permettant le **déploiement automatique** de votre projet.

### Étape 5 : Documentation Swagger

Vous devrez proposer une documentation complète de votre API sous la forme d’un fichier swagger.yml ou swagger.json. Ce fichier doit :

- Documenter **tous les endpoints** de l’API,
- Inclure des **exemples de requêtes et de réponses**,
- Présenter les **schémas de votre modèle de données**,
- Être accessible depuis un endpoint type http://localhost:8080/docs via **Swagger UI**.

### Étape 6 : Documentation projet

Enfin, on vous demande de documenter entièrement votre projet, en utilisant les fichiers standard du développement.

Notamment, vous devrez avoir un fichier README.md avec :

- L’architecture de votre projet,
- Les instructions d'installation,
- Quelques commandes Docker utiles,
- Éventuellement, le schéma de la base de données,
- Les URLs des services de votre application

Vous fournirez aussi un fichier .env.example pour les variables.

## Compétences visées

- Bases relationnelles
- API REST
- Docker & Docker Compose
- CI/CD
- Swagger

## Rendu

Le projet est à rendre sur votre github :
https://github.com/prenom-nom/btp-projet-container

## Ressources

- [Swagger](https://en.wikipedia.org/wiki/Swagger_(software))

<a id="en"></a>

<div align="center">
  <a href="#en">🇬🇧 English</a> · 
  <a href="#fr">🇫🇷 français</a>
</div>

# Bachelor TP - InfraMusicStore

Develop an API for an online record store, with a fully containerized stack

## Introduction

This week-long project aims to develop the management API for an online record store in a complete Docker environment.

You will build a management tool for an online music store.

## Professional Context

You are joining the digital team of the **InfraMusicStore** store, which will distribute music online and promote artists.

Depending on your background (Web, Software, Data/AI, etc.), you master certain techniques and programming languages (JS, C++, Python, SQL, etc.), and you are capable of developing the same project with different methodologies.

You will deliver your project as a team of 2 to 3 people.

Your manager asks you to implement an API solution by first building a database that will be structured like the Chinook database [available online](https://github.com/lerocha/chinook-database).

## Project Steps

We are not asking you to build a complete management website, which would be too time-consuming, but simply to **set up a REST API that allows access to your Chinook relational database data**, with read/write capabilities: all CRUD operations must be supported.

You will need to deploy an API making your database data available, in a fully dockerized environment, and that uses an **automated CI/CD pipeline**.

### Step 1: The Database

You must set up a **relational database with multiple interconnected tables**: artists, albums, tracks, genres... these tables are provided in the Chinook repository. It will be up to you to import them into your system and ensure that the connections are properly made.

![ER diagram](./tables.png "illustration of tables")

The expectations are as follows:

- Clear definition of **relationships and foreign keys**,
- Provision of SQL initialization scripts (DDL),
- The recommended DB is **MariaDB**, but other relational database engines are accepted as long as they are Open Source.
- You must ensure your tables are properly **normalized**.
- You must integrate test data that is not in the base tables (at least 20 entries for 4 tables). This data must be added **programmatically**: you cannot make manual additions via a GUI!
    - Web and Software students can be inspired by a dataset [provided](https://github.com/lerocha/chinook-database/releases) in JSON or SQL with the Chinook repository.
    You can use these examples as inspiration to integrate your test data.
    - Data/AI students are invited to integrate the dataset provided with the Chinook repository, and to produce a Jupyter notebook for Exploratory Data Analysis.

### Step 2: REST API

The application must expose a **complete REST API**, connected to the database via **environment variables**.

The minimum expected endpoints are those of a CRUD. For example, for tracks:

- **GET** /api/tracks — all tracks
- **GET** /api/tracks/:id — track details
- **POST** /api/tracks — creates a track
- **PUT** /api/tracks/:id — modifies a track
- **DELETE** /api/tracks/:id — deletes a track.

This pattern should apply to **at least 4 resources**.

Other requirements:

- All your responses must be in **JSON** format.
- Optional: implementation of an **authentication system**: the API client must be authenticated, via the system of your choice (keep it simple!). If the client is not authenticated, your API must return **the correct HTTP code**.
- You must properly handle **basic HTTP errors** (403, 404, 500, etc.).
- For your API, you can use the following technologies: **Web** -> Express / [Node.js](http://node.js), **Software** -> Go-Fiber, **Data/AI** -> Flask
- Optional: management of **pagination**, **filters**, **sorting**.

### Step 3: Docker Architecture

Now that your API is functional, you can begin containerization. Your application's environment must be fully containerized with **Docker**. You must use **Docker Compose** with at least **three distinct services**:

1. The **main application** (your API)
2. The **database**
3. A database administration tool (Adminer, phpMyAdmin, etc.)
4. (optional) Swagger UI for interactive API documentation.

The following points must be present:

- A **Dockerfile** specific to the API,
- A complete **Docker Compose** configuration with the three services,
- Proper management of **Docker networks** for communication between containers,
- Optional: configuration of **volumes** to ensure **data persistence** for the database

### Step 4: GIT Pipeline + CI/CD

Once your application is fully dockerized, we expect you to set up real **continuous integration** for your project. In short, your application deployment must happen automatically!

Version control must be managed via **GitHub**:

- You must use the **Pull Requests** system,
- You must implement a **GitHub Actions workflow** enabling **automatic deployment** of your project.

### Step 5: Swagger Documentation

You will need to provide complete documentation of your API in the form of a swagger.yml or swagger.json file. This file must:

- Document **all API endpoints**,
- Include **examples of requests and responses**,
- Present the **schemas of your data model**,
- Be accessible from an endpoint like http://localhost:8080/docs via **Swagger UI**.

### Step 6: Project Documentation

Finally, you are asked to fully document your project using standard development files.

In particular, you must have a README.md file with:

- Your project architecture,
- Installation instructions,
- Some useful Docker commands,
- Optionally, the database schema,
- The URLs of your application services

You will also provide a .env.example file for variables.

## Targeted Skills

- Relational databases
- REST API
- Docker & Docker Compose
- CI/CD
- Swagger

## Submission

The project must be submitted on your github:
https://github.com/first-name-last-name/btp-projet-container

## Resources

- [Swagger](https://en.wikipedia.org/wiki/Swagger_(software))
