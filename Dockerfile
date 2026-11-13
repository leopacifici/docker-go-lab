# PRIMO STAGE: contiene compilatore e strumenti Go.
# AS assegna allo stage il nome "build", usato poi da COPY.
FROM golang:1.25-alpine AS build

# Crea /src se manca e la imposta come directory corrente.
WORKDIR /src

# Copia main.go dal contesto di build in /src/main.go.
COPY main.go .

# Compila il programma durante la build.
# CGO_ENABLED=0 disabilita l'uso di cgo: questa applicazione
# può essere compilata senza dipendere da librerie C.
# Produce il binario /out/server.
RUN CGO_ENABLED=0 go build -o /out/server main.go

# SECONDO STAGE: parte da una nuova base Alpine.
# Non eredita file e configurazione del primo stage.
FROM alpine:3.22
RUN addgroup -S -g 1000 app \
    && adduser -S -D -u 1000 -G app app

# Crea /app e la imposta come directory corrente,
# anche per il comando di avvio del container.
WORKDIR /app

# Copia il binario dal primo stage in /app/server.
# --from=build indica lo stage sorgente.
COPY --from=build /out/server ./server

# Valore predefinito, sostituibile all'avvio con docker run -e.
ENV MESSAGGIO="Laboratorio Docker con Go"

# Documenta la porta del server; non la pubblica sull'host.
EXPOSE 8000
USER app:app

# All'avvio del container esegue /app/server.
CMD ["./server"]
