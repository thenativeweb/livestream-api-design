# Beispiel für fachlich getriebenes API-Design

Am 23. August 2023 haben wir auf YouTube einen [Livestream zum Thema "fachlich getriebenes API-Design"](https://www.youtube.com/watch?v=SexQcBUp3DM) durchgeführt. In diesem Repository findest Du den Quellcode, der im Rahmen dieses Livestreams entstanden ist.

## Wichtiger Hinweis zum Code

Der Code ist bewusst so geschrieben, dass er möglichst einfach nachvollziehbar ist. Das heißt, dass wir auf viele Dinge verzichtet haben, die in einer produktiven Anwendung unbedingt notwendig sind. Dazu gehören unter anderem:

- Authentifizierung
- Autorisierung
- Security
- Tests
- …

Auch lag der Fokus während des Livestreams auf dem API-Design und nicht auf der Implementierung. Das heißt, dass wir uns nicht immer an bewährte Praktiken gehalten haben. Das betrifft unter anderem:

- Codestruktur
- Namenskonventionen
- Fehlerbehandlung
- …

Die Kernaussage des Livestreams dreht sich um den konzeptionellen Ansatz, nicht um die konkrete Implementierung.

## Überblick über die API

### Commands

- Buch erfassen
- Buch ausleihen

### Queries

- Alle Bücher

## Server starten

Um den Server zu starten, muss Go 1.21 installiert sein. Anschließend kann der Server mit folgendem Befehl gestartet werden:

```shell
$ go run .
```

Als Port wird `3000` verwendet. Dieser kann in der Datei `main.go` angepasst werden.

## Den Server aufrufen

### Buch erfassen

```shell
$ curl \
    -i \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"title": "...", "author": "..."}' \
    http://localhost:3000/bestand/erfasse-buch
```

### Buch ausleihen

```shell
$ curl \
    -i \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"id": "..."}' \
    http://localhost:3000/ausleihe/leihe-buch-aus
```

### Liste aller Bücher abrufen

```shell
$ curl \
    -i \
    -X GET \
    http://localhost:3000/bestand/buecher
```
