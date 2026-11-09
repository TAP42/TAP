# Le Dernier Jeton

Nom provisoire. Casino clandestin complètement absurde, en 2D et en tour par tour.
Chaque salle est un jeu différent, plus débile et plus border que la précédente.
Tu arrives avec 100 jetons. Le but : finir riche, ou au moins pas mort.

## Règles générales

### Jetons
- Départ : **100 jetons**.
- `BALANCE` affiche ton solde.
- Pour jouer dans une salle : `BET <montant>`. Mise min/max propre à chaque salle.
- Gagné : mise × cote de la salle. Perdu : mise perdue, comme au casino.

### Fauché
- À 0 jeton, le Boss te prête 50 jetons. Tu as maintenant une **dette**.
- Tant que la dette n'est pas remboursée : accès VIP fermé, et Gérard le videur te regarde mal.

### Points de vie
- Certaines salles font mal (Fight Club, Bouton Rouge, gifle du croupier).
- HP à 0 = **KO** : réveil à l'infirmerie clandestine, mise perdue.

### Tour par tour (commun à toutes les salles)
1. Une partie démarre quand assez de joueurs ont misé. S'il manque du monde, des PNJ complètent la table.
2. **Phase de choix** : chaque joueur a ~20 s pour envoyer son action (`PLAY <action>`, `ATTACK`, `DEFEND`…).
3. **Résolution** : les actions sont jouées dans l'ordre d'initiative.
4. Le serveur envoie le résultat à tous les joueurs de la salle (`EVT`).
5. Pas d'action envoyée à temps = action par défaut (passer, ou `DEFEND` en combat).
6. On recommence jusqu'à la fin de la partie, puis gains et pertes sont distribués.

## Plan

```
                 [Bureau du Boss] (VIP)
                         |
 [Toilettes] ——— [Hub : Le Lobby] ——— [Roulette Truquée]
                    |          |
             [Fight Club]  [Couloir] ——— [Course de Hamsters]
                               |
                 [Bouton Rouge] ——— [Blackjack du Tricheur]

 [Infirmerie clandestine] : on y arrive après un KO, sortie vers le Hub
```

## Les salles

### Hub — Le Lobby
Point de départ et de retour.
- **Gérard le videur** : tutoriel, première quête.
- **Josiane la barmaid** : vend des boissons et des objets (`BUY`).
- **Le caissier** : solde, remboursement de la dette.
- Tableau des quêtes.

### La Roulette Truquée
- Tout le monde mise sur rouge, noir, pair, impair ou un numéro.
- **Monsieur Lucien**, croupier, lance la roue.
- Cases spéciales :
  - **0 Maison** : le croupier rafle toutes les mises.
  - **Case clown** : le croupier gifle un joueur au hasard (−10 HP).

### Le Fight Club Poudré
- Battle royale en tour par tour : `ATTACK <cible>`, `DEFEND`, `FLEE`, `USE <item>`.
- Le dernier debout rafle toute la cagnotte.
- `FLEE` : tu sors vivant, mais tu perds ta mise.
- Objets :
  - **Poudre suspecte** : +50 % de dégâts, −50 % de défense, pendant 3 tours.
  - **Chaise pliante** : gros dégâts, se casse après 2 coups.
  - **Bouteille** : soigne ou blesse, 50/50.
- PNJ :
  - **Dédé la Narine** : combattant PNJ.
  - **L'arbitre aveugle** : ne voit jamais les coups bas.

### La Course de Hamsters Dopés
- Tu paries sur un hamster parmi 5.
- À chaque tour, chaque hamster avance, et des événements au hasard tombent : crise cardiaque, raccourci, un hamster en mange un autre.
- Pendant la course, tu peux dépenser des jetons pour `PLAY boost <hamster>` ou `PLAY sabote <hamster>`.

### Le Bouton Rouge
- Jeu « push your luck » : une cagnotte commune grossit à chaque tour.
- Chaque tour, chacun choisit `PLAY appuie` (reste dans la partie) ou `PLAY sors` (encaisse sa part).
- Le risque d'explosion augmente à chaque tour. Si ça explose, ceux qui sont restés perdent tout, plus 30 HP.

### Le Blackjack du Tricheur
- Blackjack classique contre un croupier PNJ qui triche ouvertement (il sort un as de sa manche).
- `PLAY accuse` :
  - s'il trichait vraiment ce tour-ci, gain × 3 ;
  - sinon, le vigile te tabasse.

### Les Toilettes
- Pas de jeu. Marché noir.
- **Le Fournisseur** vend des objets douteux (`BUY`) et donne des quêtes louches.

### Bureau du Boss (VIP)
- Accès : 1000 jetons **ou** la quête « Survivant ».
- Partie finale : quitte ou double contre le Boss, en duel tour par tour.

### Infirmerie clandestine
- Réveil après un KO. Un docteur sans diplôme te remet sur pied (HP au max, contre des jetons si tu en as).

## Quêtes

| Quête | Donnée par | Objectif | Récompense |
|---|---|---|---|
| Premier pas | Gérard | Jouer 1 partie, n'importe quelle salle | 20 jetons |
| Dette | Le Boss (automatique) | Rembourser le prêt | Accès VIP rétabli |
| Livraison | Le Fournisseur | Porter un paquet à Josiane | Poudre suspecte |
| Chuchoteur de hamsters | Josiane | Gagner 3 courses | 150 jetons |
| Survivant | Dédé la Narine | Gagner 1 Fight Club | Accès Bureau du Boss |
| Balance | Le caissier | Démasquer 3 tricheries au Blackjack | 200 jetons |

## Groupes
- Un groupe peut faire une **mise commune** et partager les gains.
- Au Fight Club, les membres d'un même groupe ne peuvent pas se frapper entre eux, tant que le groupe tient.

## Commandes

Commandes de la RFC utilisées : `CONNECT`, `LOOK`, `MOVE`, `QUIT`, `CHAT`, `WHO`, `GROUP ...`, `TAKE`, `DROP`, `INVENTORY`, `TALK`, `ATTACK`, `STATUS`, `QUEST`, `QUESTS`.

Ajouts à documenter dans la section « Protocol Implementation » du README :

| Commande | Rôle |
|---|---|
| `BALANCE` | Solde de jetons et dette |
| `BET <montant>` | Miser pour rejoindre la partie de la salle |
| `PLAY <action>` | Action de jeu propre à la salle (`rouge`, `appuie`, `accuse`…) |
| `DEFEND` | Se protéger pendant le tour |
| `FLEE` | Fuir le combat (mise perdue) |
| `USE <item>` | Utiliser un objet |
| `BUY <item>` | Acheter à un PNJ marchand |

Événements à ajouter :

| Événement | Rôle |
|---|---|
| `EVT TURN` | Début ou fin d'un tour, temps restant, à qui de jouer |
| `EVT GAME` | Début et résultat d'une partie, gains et pertes |
| `EVT TOKENS` | Changement de solde |

## 2D
- Le serveur n'envoie **jamais d'image**, seulement l'état du jeu (`OK {json}` et `EVT`).
- La GUI (Rust) dessine chaque salle à partir de son `id` : roue, ring, piste de hamsters, bouton, table…
- La CLI (Go) affiche les mêmes informations en texte.
