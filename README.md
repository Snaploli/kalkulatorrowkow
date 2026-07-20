# Kalkulator Rowków

[![Wails Version](https://img.shields.io/badge/Wails-v2-blue.svg)](https://wails.io/)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org/)
[![Vite](https://img.shields.io/badge/Vite-3.x-646CFF.svg)](https://vitejs.dev/)

**Kalkulator Rowków** to narzędzie desktopowe ułatwiające inżynierom i mechanikom szybkie wyliczanie wymiarów rowków pod uszczelki, bazując na podanych wymiarach wewnętrznych i zewnętrznych. Aplikacja automatycznie sugeruje odpowiednie uszczelki standardu R oraz BX z wbudowanej bazy danych.

---

## 🚀 Technologie

Projekt został zbudowany przy użyciu nowoczesnego i lekkiego stacku technologicznego, aby zapewnić mały rozmiar pliku wykonywalnego, niesamowitą wydajność oraz świetny wygląd:

- **Backend:** [Go](https://go.dev/) (język kompilowany, obsługuje logikę, obliczenia i wyszukiwanie uszczelek)
- **Framework:** [Wails v2](https://wails.io/) (umożliwia budowanie aplikacji desktopowych z użyciem Go i technologii webowych, korzystając z wbudowanego systemowego silnika przeglądarki)
- **Frontend:** HTML, Vanilla CSS, Vanilla JavaScript
- **Dev Server / Bundler:** [Vite](https://vitejs.dev/)

## 🛠️ Wymagania wstępne

Aby skompilować projekt, potrzebujesz w systemie zainstalowanych poniższych narzędzi:

1. **Go** (wersja 1.18+): [Pobierz stąd](https://go.dev/doc/install)
2. **Node.js** (wersja 16+): [Pobierz stąd](https://nodejs.org/)
3. **Wails CLI**:
   Możesz go zainstalować za pomocą komendy Go:
   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```

## 💻 Tryb deweloperski (Live Development)

Aby uruchomić aplikację w trybie deweloperskim (z włączonym Hot-Reload dla frontendu i backendu):

```bash
# Wewnątrz katalogu głównego projektu wpisz:
wails dev
```
Otworzy to aplikację w oknie, a każda zmiana w pliku `.js`, `.css` lub `.go` natychmiastowo zostanie odzwierciedlona w aplikacji. Dodatkowo uruchomi się serwer na `http://localhost:34115`, dzięki któremu możesz rozwijać i testować frontend z poziomu zwykłej przeglądarki internetowej ze wsparciem DevTools.

## 📦 Kompilacja i budowanie (Build)

Projekt zawiera skrypt pomocniczy do budowania dla systemów **macOS** oraz **Windows**.

Aby zbudować projekt, upewnij się, że masz zainstalowane wymagania wstępne (w tym `wails`) i uruchom:

```bash
./build.sh
```

Co robi ten skrypt?
1. Kopiuje ikonę `app-icon.png` do odpowiednich miejsc we frontendzie oraz plikach budowania.
2. Odpala Wails build dla środowiska macOS (`darwin/universal` - działa natywnie na Intelu oraz Apple Silicon M1/M2/M3).
3. Odpala Wails build dla środowiska Windows (`windows/amd64`).

### Ręczne budowanie

Jeśli chcesz uruchomić czysty proces kompilacji dla swojego systemu, po prostu wpisz:
```bash
wails build -clean
```
Gotowy plik wykonywalny (.exe dla Windows, .app dla macOS, lub binarka pod Linuksem) pojawi się w folderze `build/bin/`.

## 📄 Licencja

Projekt dystrybuowany na licencji MIT.
