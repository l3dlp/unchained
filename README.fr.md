# Reprendre Vos Données

**Langues :** [English](README.md) • [Français](README.fr.md) • [Español](README.es.md) • [Português](README.pt.md) • [Italiano](README.it.md)

Depuis plus d'une décennie, vous donnez vos pensées, vos souvenirs, vos emails à des corporations. Elles les ont minées, analysées, vendues aux plus offrants. Maintenant vous les voulez *de retour*. Dans un format qui ne nécessite pas leur permission pour lire. Sur *votre* machine. Jusqu'à la fin.

Ce repo est votre arsenal.

## Les Outils

### ChatGPT JSON vers Markdown
`openai/chatgpt-json2md`

Votre export ChatGPT dort là, un mur de JSON. On le transforme en fichiers Markdown propres et lisibles—un par conversation. Pas de conneries, pas d'appels au cloud, pas de « désolé l'API est down ». Juste du markdown qu'on peut chercher, éditer, jeter dans Obsidian, Logseq ou ta voûte plaintext.

- Zéro dépendances. Stdlib uniquement.
- Gère les exports multiples. Détecte les batches automatiquement.
- Nettoie les noms de fichier. Trie les messages par timestamp.
- Tout est *tien* à la seconde où ça quitte les serveurs d'OpenAI.

```bash
go run main.go
```

Docs complètes [ici](openai/chatgpt-json2md/README.md).

### DAT vers Assets
`openai/chatgpt-dat2assets`

Ton export ChatGPT contient des images, de l'audio, des documents—mais ils sont tous renommés `.dat` parce qu'OpenAI pensait que tu les récupérerais jamais. On sniffe les magic bytes et on leur rend leur vrais noms. PNG, JPEG, WebP, WAV. Un passage, c'est bon.

```bash
go run main.go
```

### Claude JSON vers Markdown
`anthropic/claude-go`

Même topo que ChatGPT. Anthropic te donne du JSON, on te donne du Markdown lisible. Garde ton historique de conversation portable et tien.

```bash
go run main.go
```

### Google Webfonts Manager
`google/webfonts`

T'as jamais téléchargé des fonts de Google Fonts et tu finis avec une boîte en carton pleine de TTF ? Ce CLI génère une GUI HTML statique où tu balades ta library de fonts locale, choisis les variantes, et copies les blocs CSS `@font-face`. Favoris sauvegardés en localStorage. Marche offline. Multi-langue.

```bash
go run ./cmd/mywebfont -root webfonts -out webfonts.html -base-url webfonts
```

Docs complètes [ici](google/webfonts/README.md).

### Goodbye Gmail
`google/gmail`

Celui-ci c'est different. C'est pas juste récupérer les emails—c'est les *archiver*. Tire les fichiers `.eml` bruts directement de l'API Gmail, les parse localement, extrait chaque pièce jointe, indexe tout en SQLite avec FTS5 full-text search, et construit un système de fichiers *que tu possèdes*.

- **Hyper rapide.** Workers concurrents, respecte les limites avec backoff.
- **Dédupliqué.** Chaque pièce jointe unique stockée une fois, hachée en SHA256.
- **Indexé.** Recherche FTS5. Instantané. Local. Zéro appel API.
- **Reprendre.** Crash ? Redémarre depuis ton checkpoint. Zéro re-fetch.
- **Parsing profond.** Extrait les contacts, signatures, réponses citées.

```bash
cd google/gmail
go run ./... auth          # OAuth vers Google
go run ./... fetch --workers 10  # Tire tes emails
go run ./... normalize     # Parse et index
go run ./... search "keywords"  # Recherche instantée
```

Docs complètes [ici](google/gmail/README.md).

---

## Pourquoi Ça Compte

T'as pas réalisé en tapant, mais tu construisais un artefact. Chaque conversation avec une IA était un point de décision, une idée affinée, un problème résolu. Tes emails c'est la chronique de ta vie : relations, boulot, luttes, croissance.

C'est pas juste des fichiers. C'est *ta propriété intellectuelle*. Et là, c'est verrouillé derrière des écrans de login et des ToS qui peuvent changer un mardi.

On est en fin de partie de l'ère des plateformes. Les corporations serrent l'étreinte parce qu'elles savent ce qu'elles ont construit a de la valeur. Elles font pas de remboursements. Donc on extrait. On convertit. On archive.

Ce repo c'est une déclaration : *tes données c'est les tiennes*. Pas temporairement. Pas conditionnellement. Les tiennes. Stockées dans des formats qui préexistent au web. Des formats qui seront toujours lisibles quand les serveurs vont pourrir.

---

## Installation

Chaque projet est standalone. Prends ce que tu besoin.

```bash
# Clone
git clone https://github.com/l3dlp/unchained.guru.git
cd unchained.guru/rip

# Converters simples: zéro setup
cd openai/chatgpt-json2md
go run main.go

# Goodbye Gmail: besoin Go 1.26+
cd google/gmail
go run ./... auth
```

**Besoin:** Go 1.26 ou plus (les vieilles versions peuvent marcher ; pas testé).

---

## Formats

Tout l'output est **portable** :
- Markdown: lisible dans n'importe quel éditeur, cherchable avec grep
- JSON: queryable, parseable, futuriste
- `.eml`: RFC 5322 standard. Import dans Thunderbird, Apple Mail, n'importe quoi
- SQLite: pas de vendor lock-in. CLI, ou query avec Python/Perl/Go

Tu possèdes le format. Tu possèdes les données.

---

## Le Manifeste Unix, Vers 1996

Ces outils suivent ce que les utilisateurs Unix savent depuis 30 ans :

- **Utilise des formats chiants.** Text. JSON. SQLite. Ils survivent à la mode.
- **Fais une chose bien.** Extrait. Converti. Indexe. Rien de plus.
- **Pas de cloud requis.** Tout tourne offline. Zéro télémétrie. Zéro appels maison.
- **Respecte l'utilisateur.** Assume que tu sais ce que tu fais. Pas de hand-holding, pas de barriers.
- **Enchaîne-les ensemble.** Pipe vers grep, awk, jq. Combine avec d'autres outils. C'est l'intérêt.

On construit pas une plateforme. On te donne des **outils**. Ce que tu fais avec tes données c'est ton affaire.

---

## License

MIT. Utilise-le. Fork-le. Améliore-le. Le rends pas aux corporations.

---

## Contribuer

T'as trouvé un bug ? Données corrompues ? Nouveau provider à supporter ? Envoie une PR. Pas de code of conduct. Pas de comités de gouvernance. Juste la bonne foi et des diffs propres.

---

**Tes données. Ta machine. Tes règles.**
