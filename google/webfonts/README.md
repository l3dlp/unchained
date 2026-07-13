# mywebfont

CLI Go minimaliste qui scanne une arborescence `webfonts/<Famille>/*.ttf` et genere une GUI statique pour copier le CSS `@font-face`.

```sh
go run ./cmd/mywebfont -root /chemin/vers/webfonts -out webfonts.html -base-url webfonts
```

Options utiles :

```sh
go run ./cmd/mywebfont -lang fr -root webfonts -out webfonts.html -base-url webfonts
```

- `-base-url` vaut `webfonts` par defaut si non renseigne.
- `-lang` force la langue du CLI : `fr`, `en`, `es`, `pt`, `it`, `de`, `pl`, `ru`, `uk`, `he`, `ar`.
- Si `-root` n'existe pas, le CLI demande s'il faut creer le repertoire.
- Les favoris sont stockes en `localStorage` et permettent de copier un CSS global.

Pose `webfonts.html` a cote du dossier `webfonts`, ouvre-le, cherche une fonte, coche les variantes, puis copie le bloc voulu.
