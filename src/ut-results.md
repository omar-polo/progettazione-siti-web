Questo testo verrà ignorato.

I vari risultati sono separati da '###', mentre gli utenti da
'##'.  Ogni utente deve avere come *metadata*:

 - sesso
 - età (in anni)
 - professione
 - frequenza-web
 - siti-viaggi ("hai mai usato siti di viaggi")
 - infoviaggiando ("hai mai usato infoviaggiando")
 - esperienza (1-5 con 1 bassa 5 alta)

(si veda il file src/ut.md dove sono presenti le domande da porre agli
utenti per comprendere meglio questi punti)

Ogni risultato deve avere i seguenti *metadata*:

 - tempo
 -

La grammatica è la seguente:
```
	U: '##' identifier '\n' metadata+ results
	metadata: key ':' value '\n'
	results: '###' identifier2? '\n' key-value
```

La grammatica dei risultati (`key-value`) è uguale a quella descritta in tasks.md:10.

*NB* i task devono essere inseriti nello stesso ordine del file tasks.md

## utente 1
sesso: maschio
età: 25 anni
professione: studente
frequenza-web: 12 ore al giorno
siti-viaggi: Sì
infoviaggiando: No
esperienza: 4

### Task 1

esito:
non passato

livello:
basso

tempo:
4 minuti 12 secondi

percorso:
 - sicurezza
 - servizi
 - infrastrutture
 - telecamere
 - notizie
 - servizi
 - calcolo pedaggi

problemi:
Molto confusionario e poco chiaro nella divisione delle funzionalità

opinioni:
nessuno

### Task 2

esito:
passato

livello:
basso

tempo:
45 secondi

percorso:
 - servizi
 - previsioni meteo

problemi:
nessuno

opinioni:
Diceva che era meglio fare una pagina apposita per il meteo perchè si aspettava cose diverse dalla pagina servizi

### Task 3

esito:
passato

livello:
basso

tempo:
1 minuto e 42 secondi

percorso:
 - servizi
 - notizie

problemi:
nessuno

opinioni:
se ti serve il traffico dal punto A al punto B quando fai la ricerca ti mostra dov'é il traffico mostrandoti il percorso e i punti precisi in cui prevedono che ci sia il traffico

### Task 4

esito:
passato

livello:
basso

tempo:
10 secondi

percorso:
 - telecamere

problemi:
nessuno

opinioni:
nessuno

### Task 5

esito:
passato

livello:
basso

tempo:
30 secondi

percorso:
 - infrastrutture

problemi:
nessuno

opinioni:
nessuno

### Task 6

esito:
passato

livello:
basso

tempo:
1 minuto e 15 secondi

percorso:
 - infrastrutture
 - sicurezza
 - notizie
 - infrastrutture

problemi:
nessuno

opinioni:
Sarebbe meglio darti l'opzione di selezionare un percorso e poi in automatico ti dice il benzinaio che si trova nbel percorso dove la benzina o il diesel costano meno

## utente 2
sesso: maschio
età: 22 anni
professione: studente
frequenza-web: 15-16 ore al giorno
siti-viaggi: Sì
infoviaggiando: No
esperienza: 4

### Task 1

esito:
passato

livello:
basso

tempo:
30 secondi

percorso:
 - home page pulsante calcolo pedaggi

problemi:
Non è immediato trovare il pulsante per il calcolo del pedaggio.

opinioni:
nessuno

### Task 2

esito:
passato

livello:
basso

tempo:
1 minuto e 45 secondi

percorso:
 - infotraffico
 - notizie
 - infotraffico
 - calcolo pedaggio
 - previsioni del meteo

problemi:
È troppo incasinato è nascosto all'utente

opinioni:
nessuno

### Task 3

esito:
passato

livello:
basso

tempo:
1 minuto e 50 secondi

percorso:
 - notizie
 - pedaggio
 - notizie

problemi:
importante: Difficile capite dove sono si trova un casello perchè il traffico è diviso per tratte complessive autostradali (AUTOVIE VENETE, AUTOSTRADA BS-PD, CAV)
minore: Lo zoom sulla mappa crea problemi con lo scrolling del sito

opinioni:
nessuno

### Task 4

esito:
passato

livello:
basso

tempo:
30 secondi

percorso:
 - telecamere

problemi:
nessuno

opinioni:
nessuno

### Task 5

esito:
non passato

livello:
basso

tempo:
2 minuti e 15 secondi

percorso:
 - infrastrutture

problemi:
critico: Mentre è facile vedere dove sono le stazione di benzina più vicina nella mappa complessiva alcuni filtri non funzionano

opinioni:
I filtri non funzionano e provando ad usare i filtri non ha passato il test

### Task 6

esito:
non passato

livello:
basso

tempo:
3 minuti e 4 secondi

percorso:
 - servizi
 - cerca nel sito
 - telecamere
 - notizie
 - infrastrutture

problemi:
nessuno

opinioni:
Non è immediato il trovare le informazioni

## utente 3
sesso: femmina
età: 18 anni
professione: studente
frequenza-web: 5-6 ore al giorno
siti-viaggi: Sì
infoviaggiando: No
esperienza: 3

### Task 1

esito:
passato

livello:
alto

tempo:
42 secondi

percorso:
 - home page
 - servizi (nella nav)
 - calcolo pedaggi

problemi:
nessuno

opinioni:
Trova che la home page sia troppo confusionaria: "ho avuto paura di perdermi" afferma. In ogni caso, risolve velocemente il task.

### Task 2

esito:
passato

livello:
alto

tempo:
34 secondi

percorso:
 - mappa del sito
 - previsioni meteo
 - venezia

problemi:
nessuno

opinioni:
nessuno

### Task 3

esito:
non passato

livello:
medio-basso

tempo:
4 minuti e 4 secondi

percorso:
 - home page
 - click su "Notizie ai viaggiatori e previsioni traffico"
 - click sulla data richiesta
 - non è sicura del nome dell'autostrada/del fatto che riguardi la tratta che le interessi
 - click su mappa del sito
 - click su Previsioni traffico
 - si convice che "autovie venete" è quello che le interessa

problemi:
importante: Lo zoom della mappa durante lo scrolling se il cursore ci si posiziona sopra.

opinioni:
Dopo due minuti ha avuto bisogno di un piccolo incoraggiamento per
continuare. Pensa che una soluzione, simile a quanto fatto su altri siti,
di chiedere all'utente gli estremi di un percorso e quindi fornire le
informazioni solo riguardante quella tratta.

### Task 4

esito:
non passato

livello:
basso

tempo:
2 minuti e 12

percorso:
 - click su "Vedi tutte le telecamere" (sotto la mappa)
 - mette la mappa fullscreen
 - identifica il punto richiesto
 - torna nella home page
 - click su "Telecamere" nella navbar
 - usa il widget con l'anteprima delle telecamere sotto la mappa cercando quella che le interessa
 - click su "come usare le telecamere"
 - legge il testo
 - clicca sulla telecamera corretta nella mappa

problemi:
minore: lo scrolling che triggera lo zoom della mappa.

opinioni:
Non ha riconosciuto le icone sulla mappa come cliccabili, e quindi anche
se in relativamente poco tempo (~35 secondi) ha raggiunto la pagina
corretta, è poi tornata nella home per cercare ancora.


### Task 5

esito:
passato

livello:
medio

tempo:
40 secondi

percorso:
 - infrastrutture
 - consulta velocemente la lista ma poi preferisce usare la mappa
 - trova la stazione richiesta sulla mappa

problemi:
nessuno

opinioni:
nessuno

### Task 6

esito:
non passato

livello:
basso

tempo:
2 minuti e 20 secondi.

percorso:
 - home
 - infrastrutture
 - uso della mappa
 - click ripetuti sulla scritta "A4" sperando che sia interagibile
 - cerca A4 nella barra in alto
 - torna nella home
 - nota il widget "Dove il pieno costa meno" ma nota che non è interagibile
 - infrastrutture
 - nota la lista delle stazione di servizio sotto la mappa e trova quello richiesto.

problemi:
Zoom della mappa durante lo scrolling.

opinioni:
nessuno

## utente 4
sesso: femmina
età: 51 anni
professione: impiegata
frequenza-web: 0-1 ore al giorno
siti-viaggi: No
infoviaggiando: No
esperienza: 1

### Task 1

esito:
passato

livello:
basso

tempo:
1 minuto 3 secondi

percorso:
 - calcolo pedaggio

problemi:
Home page troppo confusionaria

opinioni:
La home page ha troppe cose e non sono ben mostrate

### Task 2

esito:
passato

livello:
basso

tempo:
2 minuto e 18 secondi

percorso:
 - notizie
 - servizi
 - previsioni meteo

problemi:
Individuare la pagina in cui sono presenti le previsioni meteo

opinioni:
La pagina meteo meriterebbe un link a se stante nel menù

### Task 3

esito:
passato

livello:
basso

tempo:
1 minuto e 26 secondi

percorso:
 - servizi
 - notizie

problemi:
Non si capisce quale delle 3 autostrade devi scegliere

opinioni:
Sarebbe meglio iinvece di dividere le autostrate in solo 3 blocchi divederle ogni due caselli (caselli adiacenti)

### Task 4

esito:
passato

livello:
basso

tempo:
47 secondi

percorso:
 - telecamere

problemi:
Lo zoom della mappa

opinioni:
nessuno

### Task 5

esito:
non passato

livello:
basso

tempo:
2 minuti e 4 secondi

percorso:
 - servizi
 - infrastrutture

problemi:
Lo zoom della mappa

opinioni:
nessuno

### Task 6

esito:
passato

livello:
basso

tempo:
1 minuti e 51 secondi

percorso:
 - infrastrutture

problemi:
nessuno

opinioni:
Sarebbe stato meglio che per ogni tratta il benzinaio con i costi minori fosse evidenziato
