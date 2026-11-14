# Planning TAP

## 1. Protocole

| Partie                         | Langage | Dev |
| ------------------------------ | ------- | --- |
| Lecture RFC + liste des écarts | Doc     |     |
| Format des commandes           | Doc     |     |
| Format des réponses            | Doc     |     |
| Format des événements          | Doc     |     |
| Liste des codes d'erreur       | Doc     |     |

## 2. Serveur back — socle

| Partie                                             | Langage | Dev       |
| -------------------------------------------------- | ------- | --------- |
| Listener TCP + goroutine par client                | C       | Theo      |
| Machine d'états de connexion                       | Go      | Noam      |
| Greeting initial                                   | Go      | Noam      |
| Authentification                                   | Go      | Noam      |
| Déconnexion propre                                 | Go/C    | Theo/Noam |
| Dispatcher de commandes                            | Go      | Noam      |
| Chargement `commands.json`                         | Go      | Noam      |
| Réponses d'erreur standard                         | Go      | Noam      |
| Système d'événements (broadcast)                   | Go      | Noam      |
| Registre des joueurs                               | Go      | Noam      |
| Validation des entrées                             | Go/C    | Theo/Noam |
| Limites de ressources (anti-flood, taille message) | C       | Theo      |
| Logging serveur                                    | C       | Theo      |

## 3. CLI back

| Partie                                     | Langage | Dev  |
| ------------------------------------------ | ------- | ---- |
| Connexion TCP au serveur                   | Go      | Noam |
| Envoi des commandes                        | Go      | Noam |
| Lecture réponses + événements en parallèle | Go      | Noam |
| Parsing des messages du protocole          | Go      | Noam |
| Gestion déconnexion / reconnexion          | Go      | Noam |

## 4. CLI front

| Partie                                              | Langage | Dev  |
| --------------------------------------------------- | ------- | ---- |
| Prompt + bannière                                   | Go      | Noam |
| Commandes locales (HELP, CLEAR, QUIT)               | Go      | Noam |
| Affichage des réponses                              | Go      | Noam |
| Affichage des événements (chat, combat, salle)      | Go      | Noam |
| Couleurs / mise en forme                            | Go      | Noam |
| Affichage des parties (tour, timer, mise, résultat) | Go      | Noam |

## 5. Serveur back — gameplay

| Partie                                                          | Langage | Dev  |
| --------------------------------------------------------------- | ------- | ---- |
| Hub (Lobby) + déplacements entre salles                         | Go      | Noam |
| LOOK / description des salles                                   | Go      | Noam |
| Économie : jetons, `BALANCE`                                    | Go      | Noam |
| Mises : `BET`, gains selon la cote, perte de la mise            | Go      | Noam |
| Dette quand fauché (prêt du Boss)                               | Go      | Noam |
| Moteur tour par tour commun (phase de choix, timer, résolution) | Go      | Noam |
| Initiative + action par défaut si pas de réponse                | Go      | Noam |
| Tables : rejoindre, quitter, PNJ qui complètent                 | Go      | Noam |
| Commande générique `PLAY <action>`                              | Go      | Noam |
| Événements `EVT TURN` / `EVT GAME` / `EVT TOKENS`               | Go      | Noam |
| Salle Roulette Truquée                                          | Go      | Noam |
| Salle Fight Club Poudré (`ATTACK`, `DEFEND`, `FLEE`, `USE`)     | Go      | Noam |
| Combat : formules de dégâts + effets (poudre, chaise…)          | Go      | Noam |
| Salle Course de Hamsters Dopés                                  | Go      | Noam |
| Salle Bouton Rouge                                              | Go      | Noam |
| Salle Blackjack du Tricheur (`PLAY accuse`)                     | Go      | Noam |
| Toilettes : marché noir (`BUY`)                                 | Go      | Noam |
| Bureau du Boss : accès VIP + duel final                         | Go      | Noam |
| HP, KO + infirmerie clandestine                                 | Go      | Noam |
| PNJ + dialogues (`TALK`)                                        | Go      | Noam |
| Items + inventaire (`TAKE`, `DROP`, `USE`, `INVENTORY`)         | Go      | Noam |
| Ressources dynamiques (unicité, cycle de vie)                   | Go      | Noam |
| Noms multi-mots                                                 | Go      | Noam |
| Chat (`CHAT GLOBAL/ROOM/GROUP`)                                 | Go      | Noam |
| Groupes : mise commune, pas de coups entre membres              | Go      | Noam |
| Quêtes : progression                                            | Go      | Noam |
| Quêtes : validation                                             | Go      | Noam |
| Quêtes : récompenses + chaînes (accès VIP)                      | Go      | Noam |

## 6. Data / monde

| Partie                                        | Langage | Dev |
| --------------------------------------------- | ------- | --- |
| Carte : hub + salles + sorties (`world.yaml`) | YAML    |     |
| Règles des salles (mise min/max, cotes)       | YAML    |     |
| PNJ + dialogues                               | YAML    |     |
| Items (effets, prix)                          | YAML    |     |
| Quêtes                                        | YAML    |     |
| Stats combat (joueurs, PNJ)                   | YAML    |     |

## 7. GUI back

| Partie                                         | Langage | Dev  |
| ---------------------------------------------- | ------- | ---- |
| Client TCP                                     | Rust    | Fiti |
| Parsing du protocole                           | Rust    | Fiti |
| Mise à jour de l'état (`state.rs`)             | Rust    | Fiti |
| Gestion déconnexion                            | Rust    | Fiti |
| Parsing `EVT TURN` / `EVT GAME` / `EVT TOKENS` | Rust    | Fiti |

## 8. GUI front

| Partie                                          | Langage | Dev  |
| ----------------------------------------------- | ------- | ---- |
| Panneau top (statut)                            | Rust    | Fiti |
| Panneau room (salle)                            | Rust    | Fiti |
| Panneau side (joueur, inventaire, groupe)       | Rust    | Fiti |
| Panneau input (saisie commandes)                | Rust    | Fiti |
| Palette / thème                                 | Rust    | Fiti |
| Affichage combat                                | Rust    | Fiti |
| Affichage chat                                  | Rust    | Fiti |
| Dessin 2D du Hub                                | Rust    | Fiti |
| Dessin 2D Roulette (roue qui tourne)            | Rust    | Fiti |
| Dessin 2D Fight Club (ring, KO)                 | Rust    | Fiti |
| Dessin 2D Course de Hamsters (piste)            | Rust    | Fiti |
| Dessin 2D Bouton Rouge (explosion)              | Rust    | Fiti |
| Dessin 2D Blackjack (table, cartes)             | Rust    | Fiti |
| Dessin 2D Toilettes, Bureau du Boss, Infirmerie | Rust    | Fiti |
| Affichage du tour (timer, à qui de jouer)       | Rust    | Fiti |
| Affichage jetons + mise en cours                | Rust    | Fiti |

## 9. Tests

| Partie                                | Langage | Dev  |
| ------------------------------------- | ------- | ---- |
| Tests unitaires serveur               | Go      | Noam |
| Tests multijoueur (plusieurs clients) | Go      | Noam |
| Tests combat                          | Go      | Noam |
| Tests quêtes                          | Go      | Noam |
| Tests coupure réseau                  | Go      | Noam |
| Tests GUI                             | Rust    | Fiti |
| Tests moteur tour par tour            | Go      | Noam |
| Tests économie (mises, gains, dette)  | Go      | Noam |
| Tests de chaque salle                 | Go      | Noam |

## 10. Build / rendu

| Partie                                        | Langage  | Dev |
| --------------------------------------------- | -------- | --- |
| Makefile (server, client, gui)                | Make     |     |
| README : Description, Instructions, Resources | Markdown |     |
| README : Architecture                         | Markdown |     |
| README : Protocol Implementation              | Markdown |     |
| README : Combat System                        | Markdown |     |
| README : Quest System                         | Markdown |     |
| README : World Design                         | Markdown |     |
| README : Server Logging                       | Markdown |     |
| README : Group Contributions                  | Markdown |     |
| README : Building and Running                 | Markdown |     |
| README : Testing                              | Markdown |     |
| Fusion des branches dans `main`               | Git      |     |
