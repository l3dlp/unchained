# Recupera Tus Datos

**Idiomas:** [English](README.md) • [Français](README.fr.md) • [Español](README.es.md) • [Português](README.pt.md) • [Italiano](README.it.md)

Hace más de una década le das tus pensamientos, tus memorias, tus emails a corporaciones. Las han minado, analizadas, vendidas al mejor postor. Ahora los quieres *de vuelta*. En un formato que no requiere su permiso para leer. En *tu* máquina. Para siempre.

Este repo es tu arsenal.

## Las Herramientas

### ChatGPT JSON a Markdown
`openai/chatgpt-json2md`

Tu export de ChatGPT está ahí, un muro de JSON. Lo convertimos en ficheros Markdown limpios y legibles—uno por conversación. Sin gilipolleces, sin llamadas a la nube, sin "disculpa, la API está caída". Solo markdown que puedes buscar, editar, y tirar a Obsidian, Logseq, o tu bóveda plaintext.

- Cero dependencias. Solo stdlib.
- Maneja exports múltiples. Detecta batch automáticamente.
- Limpia nombres de ficheros. Ordena mensajes por timestamp.
- Todo es *tuyo* en el segundo en que sale de los servidores de OpenAI.

```bash
go run main.go
```

Docs completo [aquí](openai/chatgpt-json2md/README.md).

### DAT a Assets
`openai/chatgpt-dat2assets`

Tu export de ChatGPT contiene imágenes, audio, documentos—pero todos renombrados `.dat` porque OpenAI pensaba que nunca los recuperarías. Olfateamos los magic bytes y les devolvemos sus nombres reales. PNG, JPEG, WebP, WAV. Un paso, listo.

```bash
go run main.go
```

### Claude JSON a Markdown
`anthropic/claude-go`

Mismo plan que ChatGPT. Anthropic te da JSON, te damos Markdown legible. Mantén tu historial de conversación portable y tuyo.

```bash
go run main.go
```

### Google Webfonts Manager
`google/webfonts`

¿Nunca descargaste fonts de Google Fonts y acabas con una caja de TTF sueltos? Este CLI genera una GUI HTML estática donde navegas tu library local de fonts, escoges variantes, y copias los bloques CSS `@font-face`. Favoritos guardados en localStorage. Funciona offline. Multiidioma.

```bash
go run ./cmd/mywebfont -root webfonts -out webfonts.html -base-url webfonts
```

Docs completo [aquí](google/webfonts/README.md).

### Goodbye Gmail
`google/gmail`

Este es diferente. No solo recupera tus emails—los *archiva*. Saca ficheros `.eml` crudos directamente de la API de Gmail, los parsea localmente, extrae cada attachment, indexa todo en SQLite con búsqueda full-text FTS5, y construye un filesystem que *posees*.

- **Rápido como el rayo.** Workers concurrentes, respeta límites con backoff.
- **Deduplicado.** Cada attachment único almacenado una vez, hasheado en SHA256.
- **Indexado.** Búsqueda FTS5. Instantánea. Local. Sin llamadas API.
- **Reanudable.** ¿Crash? Reinicia desde tu checkpoint. Sin re-fetch.
- **Parsing profundo.** Extrae contactos, firmas, respuestas citadas.

```bash
cd google/gmail
go run ./... auth          # OAuth a Google
go run ./... fetch --workers 10  # Saca tus emails
go run ./... normalize     # Parsea e indexa
go run ./... search "keywords"  # Búsqueda instantánea
```

Docs completo [aquí](google/gmail/README.md).

---

## Por Qué Importa

No te diste cuenta al escribir, pero estabas construyendo un artefacto. Cada conversación con una IA era un punto de decisión, una idea refinada, un problema resuelto. Tus emails son la crónica de tu vida: relaciones, trabajo, luchas, crecimiento.

No son solo ficheros. Son *tu propiedad intelectual*. Y ahora están encerrados tras pantallas de login y términos de servicio que pueden cambiar un martes.

Estamos en el fin del juego de la era de las plataformas. Las corporaciones aprietan su agarre porque saben qué han construido es valioso. No dan reembolsos. Así que extraemos. Convertimos. Archivamos.

Este repo es una declaración: *tus datos son tuyos*. No temporalmente. No condicionalmente. Tuyos. Almacenados en formatos que preexisten a la web. Formatos que seguirán siendo legibles cuando los servidores se pudran.

---

## Instalación

Cada proyecto es standalone. Coge lo que necesites.

```bash
# Clona
git clone https://github.com/l3dlp/unchained.guru.git
cd unchained.guru/rip

# Converters simples: sin setup
cd openai/chatgpt-json2md
go run main.go

# Goodbye Gmail: necesita Go 1.26+
cd google/gmail
go run ./... auth
```

**Requiere:** Go 1.26 o posterior (versiones más viejas podrían funcionar; no probadas).

---

## Formatos

Todo output es **portable**:
- Markdown: legible en cualquier editor, buscable con grep
- JSON: consultable, parseble, a prueba de futuro
- `.eml`: RFC 5322 standard. Importa en Thunderbird, Apple Mail, cualquier cosa
- SQLite: sin vendor lock-in. CLI, o consulta con Python/Perl/Go

Posees el formato. Posees los datos.

---

## El Manifiesto Unix, Hacia 1996

Estas herramientas siguen lo que los usuarios de Unix saben desde hace 30 años:

- **Usa formatos aburridos.** Text. JSON. SQLite. Sobreviven a la moda.
- **Haz una cosa bien.** Extrae. Convierte. Indexa. Nada más.
- **Sin nube requerida.** Todo funciona offline. Sin telemetría. Sin llamadas a casa.
- **Respeta al usuario.** Asume que sabes lo que haces. Sin hand-holding, sin barreras.
- **Encadénalas.** Pipe a grep, awk, jq. Combina con otras herramientas. Ese es el punto.

No estamos construyendo una plataforma. Te damos **herramientas**. Qué haces con tus datos es tu asunto.

---

## Licencia

MIT. Úsalo. Forkéalo. Mejóralo. No se lo devuelvas a las corporaciones.

---

## Contribuye

¿Encontraste un bug? ¿Datos corruptos? ¿Nuevo provider para soportar? Envía un PR. Sin código de conducta. Sin comités de gobernanza. Solo buena fe y diffs limpios.

---

**Tus datos. Tu máquina. Tus reglas.**
