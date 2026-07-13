# Riprendi I Tuoi Dati

**Lingue:** [English](README.md) • [Français](README.fr.md) • [Español](README.es.md) • [Português](README.pt.md) • [Italiano](README.it.md)

Ormai da più di un decennio stai dando i tuoi pensieri, i tuoi ricordi, le tue email alle corporation. Le hanno scavate, analizzate, vendute al miglior offerente. Ora li vuoi *indietro*. In un formato che non richiede il loro permesso per leggerli. Sulla *tua* macchina. Per sempre.

Questo repo è il tuo arsenale.

## Gli Strumenti

### ChatGPT JSON a Markdown
`openai/chatgpt-json2md`

Il tuo export da ChatGPT sta lì, una muraglia di JSON. Lo trasformiamo in file Markdown puliti e leggibili—uno per conversazione. Niente cazzate, niente chiamate al cloud, niente "scusa l'API è down". Solo markdown che puoi cercare, editare, e buttare in Obsidian, Logseq, o il tuo vault plaintext.

- Zero dipendenze. Solo stdlib.
- Gestisce exports multipli. Rileva batch automaticamente.
- Pulisce i nomi dei file. Ordina i messaggi per timestamp.
- Tutto è *tuo* nel secondo in cui esce dai server di OpenAI.

```bash
go run main.go
```

Docs completo [qui](openai/chatgpt-json2md/README.md).

### DAT ad Assets
`openai/chatgpt-dat2assets`

Il tuo export da ChatGPT contiene immagini, audio, documenti—ma tutti rinominati `.dat` perché OpenAI pensava che non li avresti mai recuperati. Annusiamo i magic bytes e restituiamo i loro nomi reali. PNG, JPEG, WebP, WAV. Una passata, fatto.

```bash
go run main.go
```

### Claude JSON a Markdown
`anthropic/claude-go`

Stesso piano di ChatGPT. Anthropic ti dà JSON, ti diamo Markdown leggibile. Tieni la tua cronologia di conversazione portabile e tua.

```bash
go run main.go
```

### Google Webfonts Manager
`google/webfonts`

Non hai mai scaricato font da Google Fonts e finito con una scatola di TTF sparsi? Questo CLI genera una GUI HTML statica dove navighi la tua libreria locale di font, scegli varianti, e copi i blocchi CSS `@font-face`. Preferiti salvati in localStorage. Funziona offline. Multi-lingua.

```bash
go run ./cmd/mywebfont -root webfonts -out webfonts.html -base-url webfonts
```

Docs completo [qui](google/webfonts/README.md).

### Goodbye Gmail
`google/gmail`

Questo è diverso. Non solo recupera le tue email—le *archivia*. Tira file `.eml` grezzi direttamente dall'API di Gmail, li parsea localmente, estrae ogni allegato, indicizza tutto in SQLite con ricerca full-text FTS5, e costruisce un filesystem che *possiedi*.

- **Fulmineo.** Worker concorrenti, rispetta i limiti con backoff.
- **Deduplicato.** Ogni allegato unico memorizzato una volta, hashato in SHA256.
- **Indicizzato.** Ricerca FTS5. Istantanea. Locale. Senza chiamate API.
- **Riprendibile.** Crash? Riavvia dal tuo checkpoint. Niente re-fetch.
- **Parsing profondo.** Estrae contatti, firme, risposte citate.

```bash
cd google/gmail
go run ./... auth          # OAuth a Google
go run ./... fetch --workers 10  # Tira le tue email
go run ./... normalize     # Parsea e indicizza
go run ./... search "keywords"  # Ricerca istantanea
```

Docs completo [qui](google/gmail/README.md).

---

## Perché Importa

Non te ne sei accorto mentre scrivevi, ma stavi costruendo un artefatto. Ogni conversazione con un'IA era un punto di decisione, un'idea affinata, un problema risolto. Le tue email sono la cronaca della tua vita: relazioni, lavoro, lotte, crescita.

Non sono solo file. Sono *la tua proprietà intellettuale*. E adesso sono intrappolati dietro schermi di login e termini di servizio che possono cambiare un martedì.

Siamo all'endgame dell'era delle piattaforme. Le corporation stringono la presa perché sanno che quello che hanno costruito ha valore. Non fanno rimborsi. Allora estraiamo. Convertiamo. Archiviamo.

Questo repo è una dichiarazione: *i tuoi dati sono tuoi*. Non temporaneamente. Non condizionatamente. Tuoi. Memorizzati in formati che preesistono al web. Formati che saranno ancora leggibili quando i server marciranno.

---

## Installazione

Ogni progetto è standalone. Prendi quello che ti serve.

```bash
# Clona
git clone https://github.com/l3dlp/unchained.guru.git
cd unchained.guru/rip

# Converter semplici: zero setup
cd openai/chatgpt-json2md
go run main.go

# Goodbye Gmail: ha bisogno di Go 1.26+
cd google/gmail
go run ./... auth
```

**Richiede:** Go 1.26 o successivo (versioni più vecchie potrebbero funzionare; non testate).

---

## Formati

Tutto l'output è **portabile**:
- Markdown: leggibile in qualsiasi editor, cercabile con grep
- JSON: interrogabile, parseable, a prova di futuro
- `.eml`: RFC 5322 standard. Importa in Thunderbird, Apple Mail, qualsiasi cosa
- SQLite: niente vendor lock-in. CLI, o interroga con Python/Perl/Go

Tu possiedi il formato. Tu possiedi i dati.

---

## Il Manifesto Unix, Verso il 1996

Questi strumenti seguono quello che gli utenti Unix sanno da 30 anni:

- **Usa formati noioso.** Text. JSON. SQLite. Sopravvivono alla moda.
- **Fai una cosa bene.** Estrai. Converti. Indicizza. Nient'altro.
- **Nessun cloud richiesto.** Tutto funziona offline. Nessuna telemetria. Nessuna chiamata a casa.
- **Rispetta l'utente.** Assumi che sai cosa stai facendo. Nessun hand-holding, nessun ostacolo.
- **Concatenali insieme.** Pipe a grep, awk, jq. Combina con altri strumenti. Questo è il punto.

Non stiamo costruendo una piattaforma. Ti stiamo dando **strumenti**. Cosa fai con i tuoi dati è affar tuo.

---

## Licenza

MIT. Usalo. Forkalo. Miglioralo. Non ridarla alle corporation.

---

## Contribuisci

Trovato un bug? Dati corrotti? Nuovo provider da supportare? Invia una PR. Niente codice di condotta. Niente comitati di governance. Solo buona fede e diff puliti.

---

**I tuoi dati. La tua macchina. Le tue regole.**
