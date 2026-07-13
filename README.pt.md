# Recupere Seus Dados

**Idiomas:** [English](README.md) • [Français](README.fr.md) • [Español](README.es.md) • [Português](README.pt.md) • [Italiano](README.it.md)

Há mais de uma década você está dando seus pensamentos, suas memórias, seus emails para corporações. Elas os mineraram, analisaram, venderam para o maior lance. Agora você quer *de volta*. Em um formato que não requer sua permissão para ler. Na *sua* máquina. Para sempre.

Este repo é seu arsenal.

## As Ferramentas

### ChatGPT JSON para Markdown
`openai/chatgpt-json2md`

Seu export do ChatGPT está lá, uma parede de JSON. Transformamos em arquivos Markdown limpos e legíveis—um por conversa. Sem besteiras, sem chamadas para a nuvem, sem "desculpe a API está fora". Apenas markdown que você pode buscar, editar, e jogar em Obsidian, Logseq, ou seu cofre plaintext.

- Zero dependências. Apenas stdlib.
- Lida com múltiplos exports. Detecta batch automaticamente.
- Limpa nomes de arquivos. Ordena mensagens por timestamp.
- Tudo é *seu* no segundo em que sai dos servidores da OpenAI.

```bash
go run main.go
```

Docs completo [aqui](openai/chatgpt-json2md/README.md).

### DAT para Assets
`openai/chatgpt-dat2assets`

Seu export do ChatGPT contém imagens, áudio, documentos—mas todos renomeados `.dat` porque a OpenAI achou que você nunca os recuperaria. Farejamos os magic bytes e devolvemos seus nomes reais. PNG, JPEG, WebP, WAV. Uma passada, pronto.

```bash
go run main.go
```

### Claude JSON para Markdown
`anthropic/claude-go`

Mesmo esquema que ChatGPT. Anthropic te dá JSON, te damos Markdown legível. Mantenha seu histórico de conversa portável e seu.

```bash
go run main.go
```

### Google Webfonts Manager
`google/webfonts`

Você nunca baixou fonts do Google Fonts e acabou com uma caixa de TTF soltos? Este CLI gera uma GUI HTML estática onde você navega sua library local de fonts, escolhe variantes, e copia os blocos CSS `@font-face`. Favoritos salvos em localStorage. Funciona offline. Multi-idioma.

```bash
go run ./cmd/mywebfont -root webfonts -out webfonts.html -base-url webfonts
```

Docs completo [aqui](google/webfonts/README.md).

### Goodbye Gmail
`google/gmail`

Este é diferente. Não apenas recupera seus emails—os *arquiva*. Puxa arquivos `.eml` brutos diretamente da API do Gmail, parseia localmente, extrai cada anexo, indexa tudo em SQLite com busca full-text FTS5, e constrói um filesystem que você *possui*.

- **Rápido como o raio.** Workers concorrentes, respeita limites com backoff.
- **Deduplicado.** Cada anexo único armazenado uma vez, hasheado em SHA256.
- **Indexado.** Busca FTS5. Instantânea. Local. Sem chamadas API.
- **Retomável.** Crash? Reinicia do seu checkpoint. Sem re-fetch.
- **Parsing profundo.** Extrai contatos, assinaturas, respostas citadas.

```bash
cd google/gmail
go run ./... auth          # OAuth no Google
go run ./... fetch --workers 10  # Puxa seus emails
go run ./... normalize     # Parseia e indexa
go run ./... search "keywords"  # Busca instantânea
```

Docs completo [aqui](google/gmail/README.md).

---

## Por Que Importa

Você não percebeu ao digitar, mas estava construindo um artefato. Cada conversa com uma IA era um ponto de decisão, uma ideia refinada, um problema resolvido. Seus emails são a crônica de sua vida: relacionamentos, trabalho, lutas, crescimento.

Não são apenas arquivos. São *sua propriedade intelectual*. E agora estão trancafiados atrás de telas de login e termos de serviço que podem mudar numa terça-feira.

Estamos no endgame da era das plataformas. As corporações estão apertando o cerco porque sabem o que construíram é valioso. Elas não oferecem reembolsos. Então extraímos. Convertemos. Arquivamos.

Este repo é uma declaração: *seus dados são seus*. Não temporariamente. Não condicionalmente. Seus. Armazenados em formatos que preexistem a web. Formatos que ainda serão legíveis quando os servidores apodrecerem.

---

## Instalação

Cada projeto é standalone. Pegue o que você precisa.

```bash
# Clone
git clone https://github.com/l3dlp/unchained.guru.git
cd unchained.guru/rip

# Converters simples: sem setup
cd openai/chatgpt-json2md
go run main.go

# Goodbye Gmail: precisa Go 1.26+
cd google/gmail
go run ./... auth
```

**Requer:** Go 1.26 ou posterior (versões mais antigas podem funcionar; não testadas).

---

## Formatos

Todo output é **portável**:
- Markdown: legível em qualquer editor, buscável com grep
- JSON: consultável, parseável, à prova de futuro
- `.eml`: RFC 5322 padrão. Importe em Thunderbird, Apple Mail, qualquer coisa
- SQLite: sem vendor lock-in. CLI, ou consulta com Python/Perl/Go

Você possui o formato. Você possui os dados.

---

## O Manifesto Unix, Circa 1996

Estas ferramentas seguem o que usuários Unix sabem há 30 anos:

- **Use formatos entediantes.** Text. JSON. SQLite. Eles sobrevivem à moda.
- **Faça uma coisa bem.** Extraia. Converta. Indexe. Nada mais.
- **Sem nuvem necessária.** Tudo funciona offline. Sem telemetria. Sem ligações para casa.
- **Respeite o usuário.** Assuma que você sabe o que está fazendo. Sem hand-holding, sem barreiras.
- **Encadeie-os.** Pipe para grep, awk, jq. Combine com outras ferramentas. Esse é o ponto.

Não estamos construindo uma plataforma. Estamos lhe dando **ferramentas**. O que você faz com seus dados é seu negócio.

---

## Licença

MIT. Use-o. Fork-o. Melhore-o. Não o devolva às corporações.

---

## Contribua

Encontrou um bug? Dados corrompidos? Novo provedor para suportar? Envie um PR. Sem código de conduta. Sem comitês de governança. Apenas boa fé e diffs limpos.

---

**Seus dados. Sua máquina. Suas regras.**
