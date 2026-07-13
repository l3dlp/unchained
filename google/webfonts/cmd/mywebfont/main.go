package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"mywebfont"
)

func main() {
	initialLang := preselectedLang(os.Args[1:], detectLang())
	tr := messagesFor(initialLang)
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	lang := fs.String("lang", initialLang, tr["flagLang"])
	root := fs.String("root", "webfonts", tr["flagRoot"])
	out := fs.String("out", "webfonts.html", tr["flagOut"])
	baseURL := fs.String("base-url", "webfonts", tr["flagBaseURL"])
	title := fs.String("title", "My Webfonts", tr["flagTitle"])
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\n\n", tr["usage"])
		printDefaults(os.Stderr, fs, tr)
	}
	if err := fs.Parse(os.Args[1:]); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "%s: %s\n", tr["errorPrefix"], tr["parseError"])
		os.Exit(2)
	}
	tr = messagesFor(*lang)

	if err := ensureRootDir(*root, tr); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", tr["errorPrefix"], err)
		os.Exit(1)
	}

	if err := mywebfont.Generate(mywebfont.Options{
		Root:    *root,
		Out:     *out,
		BaseURL: *baseURL,
		Title:   *title,
	}); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", tr["errorPrefix"], err)
		os.Exit(1)
	}

	fmt.Printf(tr["generated"]+"\n", *out)
}

func ensureRootDir(root string, tr map[string]string) error {
	info, err := os.Stat(root)
	if err == nil {
		if !info.IsDir() {
			return fmt.Errorf(tr["notDir"], root)
		}
		return nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	fmt.Fprintf(os.Stderr, tr["createQuestion"]+" ", root)
	answer, readErr := bufio.NewReader(os.Stdin).ReadString('\n')
	if readErr != nil && strings.TrimSpace(answer) == "" {
		return fmt.Errorf(tr["aborted"])
	}
	if !isYes(answer) {
		return fmt.Errorf(tr["aborted"])
	}
	if err := os.MkdirAll(root, 0o755); err != nil {
		return fmt.Errorf(tr["createFailed"], err)
	}
	fmt.Fprintf(os.Stderr, tr["created"]+"\n", root)
	return nil
}

func isYes(answer string) bool {
	switch strings.ToLower(strings.TrimSpace(answer)) {
	case "y", "yes", "o", "oui", "s", "si", "sí", "sim", "j", "ja", "t", "tak", "д", "да", "так", "כן", "نعم":
		return true
	default:
		return false
	}
}

func printDefaults(w io.Writer, fs *flag.FlagSet, tr map[string]string) {
	fs.VisitAll(func(f *flag.Flag) {
		fmt.Fprintf(w, "  -%s\n    \t%s", f.Name, f.Usage)
		if f.DefValue != "" {
			fmt.Fprintf(w, " (%s %q)", tr["defaultLabel"], f.DefValue)
		}
		fmt.Fprintln(w)
	})
}

func preselectedLang(args []string, fallback string) string {
	for i, arg := range args {
		switch {
		case arg == "-lang" && i+1 < len(args):
			return normalizeLang(args[i+1])
		case arg == "--lang" && i+1 < len(args):
			return normalizeLang(args[i+1])
		case strings.HasPrefix(arg, "-lang="):
			return normalizeLang(strings.TrimPrefix(arg, "-lang="))
		case strings.HasPrefix(arg, "--lang="):
			return normalizeLang(strings.TrimPrefix(arg, "--lang="))
		}
	}
	return normalizeLang(fallback)
}

func detectLang() string {
	for _, key := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		if value := os.Getenv(key); value != "" {
			return normalizeLang(value)
		}
	}
	return "en"
}

func normalizeLang(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	value = strings.ReplaceAll(value, "_", "-")
	if value == "" || value == "c" || value == "posix" {
		return "en"
	}
	if strings.HasPrefix(value, "pt") {
		return "pt"
	}
	if strings.HasPrefix(value, "uk") {
		return "uk"
	}
	for _, lang := range []string{"fr", "en", "es", "it", "de", "pl", "ru", "he", "ar"} {
		if value == lang || strings.HasPrefix(value, lang+"-") {
			return lang
		}
	}
	return "en"
}

func messagesFor(lang string) map[string]string {
	messages, ok := translations[normalizeLang(lang)]
	if !ok {
		return translations["en"]
	}
	return messages
}

var translations = map[string]map[string]string{
	"fr": {
		"usage":          "Utilisation: mywebfont [options]",
		"flagLang":       "langue du CLI (fr, en, es, pt, it, de, pl, ru, uk, he, ar)",
		"flagRoot":       "répertoire racine des webfonts",
		"flagOut":        "fichier HTML statique à générer",
		"flagBaseURL":    "URL ou chemin public vers le répertoire webfonts",
		"flagTitle":      "titre de l'interface",
		"parseError":     "arguments invalides",
		"defaultLabel":   "défaut",
		"createQuestion": "Le répertoire %q n'existe pas. Souhaitez-vous créer un nouveau répertoire ? [y/N]",
		"created":        "Répertoire créé : %s",
		"aborted":        "opération annulée",
		"createFailed":   "impossible de créer le répertoire : %w",
		"notDir":         "%q existe mais n'est pas un répertoire",
		"errorPrefix":    "erreur",
		"generated":      "GUI générée : %s",
	},
	"en": {
		"usage":          "Usage: mywebfont [options]",
		"flagLang":       "CLI language (fr, en, es, pt, it, de, pl, ru, uk, he, ar)",
		"flagRoot":       "root directory of the webfonts",
		"flagOut":        "static HTML file to generate",
		"flagBaseURL":    "public URL or path to the webfonts directory",
		"flagTitle":      "interface title",
		"parseError":     "invalid arguments",
		"defaultLabel":   "default",
		"createQuestion": "Directory %q does not exist. Would you like to create a new directory? [y/N]",
		"created":        "Directory created: %s",
		"aborted":        "operation aborted",
		"createFailed":   "could not create directory: %w",
		"notDir":         "%q exists but is not a directory",
		"errorPrefix":    "error",
		"generated":      "GUI generated: %s",
	},
	"es": {
		"usage":          "Uso: mywebfont [opciones]",
		"flagLang":       "idioma del CLI (fr, en, es, pt, it, de, pl, ru, uk, he, ar)",
		"flagRoot":       "directorio raíz de las webfonts",
		"flagOut":        "archivo HTML estático que se generará",
		"flagBaseURL":    "URL o ruta pública al directorio webfonts",
		"flagTitle":      "título de la interfaz",
		"parseError":     "argumentos no válidos",
		"defaultLabel":   "predeterminado",
		"createQuestion": "El directorio %q no existe. ¿Desea crear un nuevo directorio? [y/N]",
		"created":        "Directorio creado: %s",
		"aborted":        "operación cancelada",
		"createFailed":   "no se pudo crear el directorio: %w",
		"notDir":         "%q existe pero no es un directorio",
		"errorPrefix":    "error",
		"generated":      "GUI generada: %s",
	},
	"pt": {
		"usage":          "Uso: mywebfont [opções]",
		"flagLang":       "idioma do CLI (fr, en, es, pt, it, de, pl, ru, uk, he, ar)",
		"flagRoot":       "diretório raiz das webfonts",
		"flagOut":        "ficheiro HTML estático a gerar",
		"flagBaseURL":    "URL ou caminho público para o diretório webfonts",
		"flagTitle":      "título da interface",
		"parseError":     "argumentos inválidos",
		"defaultLabel":   "predefinição",
		"createQuestion": "O diretório %q não existe. Deseja criar um novo diretório? [y/N]",
		"created":        "Diretório criado: %s",
		"aborted":        "operação cancelada",
		"createFailed":   "não foi possível criar o diretório: %w",
		"notDir":         "%q existe mas não é um diretório",
		"errorPrefix":    "erro",
		"generated":      "GUI gerada: %s",
	},
	"it": {
		"usage":          "Uso: mywebfont [opzioni]",
		"flagLang":       "lingua della CLI (fr, en, es, pt, it, de, pl, ru, uk, he, ar)",
		"flagRoot":       "directory radice dei webfont",
		"flagOut":        "file HTML statico da generare",
		"flagBaseURL":    "URL o percorso pubblico della directory webfonts",
		"flagTitle":      "titolo dell'interfaccia",
		"parseError":     "argomenti non validi",
		"defaultLabel":   "predefinito",
		"createQuestion": "La directory %q non esiste. Vuoi creare una nuova directory? [y/N]",
		"created":        "Directory creata: %s",
		"aborted":        "operazione annullata",
		"createFailed":   "impossibile creare la directory: %w",
		"notDir":         "%q esiste ma non è una directory",
		"errorPrefix":    "errore",
		"generated":      "GUI generata: %s",
	},
	"de": {
		"usage":          "Verwendung: mywebfont [Optionen]",
		"flagLang":       "CLI-Sprache (fr, en, es, pt, it, de, pl, ru, uk, he, ar)",
		"flagRoot":       "Wurzelverzeichnis der Webfonts",
		"flagOut":        "zu erzeugende statische HTML-Datei",
		"flagBaseURL":    "öffentliche URL oder Pfad zum Verzeichnis webfonts",
		"flagTitle":      "Titel der Oberfläche",
		"parseError":     "ungültige Argumente",
		"defaultLabel":   "Standard",
		"createQuestion": "Das Verzeichnis %q existiert nicht. Möchten Sie ein neues Verzeichnis erstellen? [y/N]",
		"created":        "Verzeichnis erstellt: %s",
		"aborted":        "Vorgang abgebrochen",
		"createFailed":   "Verzeichnis konnte nicht erstellt werden: %w",
		"notDir":         "%q existiert, ist aber kein Verzeichnis",
		"errorPrefix":    "Fehler",
		"generated":      "GUI erzeugt: %s",
	},
	"pl": {
		"usage":          "Użycie: mywebfont [opcje]",
		"flagLang":       "język CLI (fr, en, es, pt, it, de, pl, ru, uk, he, ar)",
		"flagRoot":       "katalog główny webfontów",
		"flagOut":        "statyczny plik HTML do wygenerowania",
		"flagBaseURL":    "publiczny URL lub ścieżka do katalogu webfonts",
		"flagTitle":      "tytuł interfejsu",
		"parseError":     "nieprawidłowe argumenty",
		"defaultLabel":   "domyślnie",
		"createQuestion": "Katalog %q nie istnieje. Czy chcesz utworzyć nowy katalog? [y/N]",
		"created":        "Utworzono katalog: %s",
		"aborted":        "operacja anulowana",
		"createFailed":   "nie można utworzyć katalogu: %w",
		"notDir":         "%q istnieje, ale nie jest katalogiem",
		"errorPrefix":    "błąd",
		"generated":      "Wygenerowano GUI: %s",
	},
	"ru": {
		"usage":          "Использование: mywebfont [параметры]",
		"flagLang":       "язык CLI (fr, en, es, pt, it, de, pl, ru, uk, he, ar)",
		"flagRoot":       "корневой каталог webfonts",
		"flagOut":        "статический HTML-файл для генерации",
		"flagBaseURL":    "публичный URL или путь к каталогу webfonts",
		"flagTitle":      "заголовок интерфейса",
		"parseError":     "недопустимые аргументы",
		"defaultLabel":   "по умолчанию",
		"createQuestion": "Каталог %q не существует. Хотите создать новый каталог? [y/N]",
		"created":        "Каталог создан: %s",
		"aborted":        "операция отменена",
		"createFailed":   "не удалось создать каталог: %w",
		"notDir":         "%q существует, но не является каталогом",
		"errorPrefix":    "ошибка",
		"generated":      "GUI создан: %s",
	},
	"uk": {
		"usage":          "Використання: mywebfont [параметри]",
		"flagLang":       "мова CLI (fr, en, es, pt, it, de, pl, ru, uk, he, ar)",
		"flagRoot":       "кореневий каталог webfonts",
		"flagOut":        "статичний HTML-файл для генерації",
		"flagBaseURL":    "публічний URL або шлях до каталогу webfonts",
		"flagTitle":      "заголовок інтерфейсу",
		"parseError":     "некоректні аргументи",
		"defaultLabel":   "за замовчуванням",
		"createQuestion": "Каталог %q не існує. Бажаєте створити новий каталог? [y/N]",
		"created":        "Каталог створено: %s",
		"aborted":        "операцію скасовано",
		"createFailed":   "не вдалося створити каталог: %w",
		"notDir":         "%q існує, але не є каталогом",
		"errorPrefix":    "помилка",
		"generated":      "GUI створено: %s",
	},
	"he": {
		"usage":          "שימוש: mywebfont [אפשרויות]",
		"flagLang":       "שפת ה-CLI (fr, en, es, pt, it, de, pl, ru, uk, he, ar)",
		"flagRoot":       "תיקיית השורש של webfonts",
		"flagOut":        "קובץ HTML סטטי ליצירה",
		"flagBaseURL":    "URL ציבורי או נתיב לתיקיית webfonts",
		"flagTitle":      "כותרת הממשק",
		"parseError":     "ארגומנטים לא חוקיים",
		"defaultLabel":   "ברירת מחדל",
		"createQuestion": "התיקייה %q אינה קיימת. האם ליצור תיקייה חדשה? [y/N]",
		"created":        "התיקייה נוצרה: %s",
		"aborted":        "הפעולה בוטלה",
		"createFailed":   "לא ניתן ליצור את התיקייה: %w",
		"notDir":         "%q קיים אך אינו תיקייה",
		"errorPrefix":    "שגיאה",
		"generated":      "ה-GUI נוצר: %s",
	},
	"ar": {
		"usage":          "الاستخدام: mywebfont [خيارات]",
		"flagLang":       "لغة CLI (fr, en, es, pt, it, de, pl, ru, uk, he, ar)",
		"flagRoot":       "المجلد الجذر لخطوط الويب",
		"flagOut":        "ملف HTML ثابت سيتم إنشاؤه",
		"flagBaseURL":    "رابط عام أو مسار إلى مجلد webfonts",
		"flagTitle":      "عنوان الواجهة",
		"parseError":     "وسائط غير صالحة",
		"defaultLabel":   "افتراضي",
		"createQuestion": "المجلد %q غير موجود. هل تريد إنشاء مجلد جديد؟ [y/N]",
		"created":        "تم إنشاء المجلد: %s",
		"aborted":        "تم إلغاء العملية",
		"createFailed":   "تعذر إنشاء المجلد: %w",
		"notDir":         "%q موجود ولكنه ليس مجلداً",
		"errorPrefix":    "خطأ",
		"generated":      "تم إنشاء الواجهة: %s",
	},
}
