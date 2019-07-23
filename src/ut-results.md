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
2 minuti (interrotto per tempo scaduto)

percorso:
 - sicurezza
 - servizi
 - infrastrutture
 - telecamere
 - notizie

problemi:

opinioni:

Trova il sito "confusionario e poco chiaro". 

### Task 2

esito:
passato

livello:
alto

tempo:
45 secondi

percorso:
 - servizi
 - previsioni meteo

problemi:

opinioni:
Suggerisce la creazione di una pagina apposita per il meteo perchè non si aspettava che si trovasse nella pagina dei servizi.

### Task 3

esito:
passato

livello:
medio

tempo:
1 minuto e 42 secondi

percorso:
 - servizi
 - notizie

problemi:
minore: effettuare lo scrolling nella home page può attivare lo zoom della mappa se il cursore si posiziona sopra per errore

opinioni:
Suggerisce la creazione di una pagina dove venga richiesto all'utente il percorso che desidera effettuare (da un punto A a un punto B) e dove il traffico mostrato riguardi solo quello specifico tratto di autostrada.

### Task 4

esito:
passato

livello:
medio

tempo:
10 secondi

percorso:
 - telecamere

problemi:

opinioni:

### Task 5

esito:
passato

livello:
medio

tempo:
30 secondi

percorso:
 - infrastrutture

problemi:

opinioni:

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

opinioni:
Suggerisce, come ha già fatto, la creazione di una pagina dove l'utente possa inserire gli estremi di un percorso, e dove il sistema fornisca le informazioni pertinenti, come il costo previsto e, appunto, le stazioni di servizio ordinate per costo.

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
minore: Non è immediato trovare il pulsante per il calcolo del pedaggio.

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

opinioni:
Si lamenta (di nuovo) di quanto il sito sia "incasinato".

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

### Task 4

esito:
passato

livello:
medio

tempo:
30 secondi

percorso:
 - telecamere

problemi:

opinioni:

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
importante: alcuni filtri non funzionano

opinioni:
Alcuni filtri non funzionano, ciò gli ha fatto perdere diverso tempo e a superare il limite per questo task.

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
importante: difficoltà nel trovare le informazioni

opinioni:

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

opinioni:

### Task 3

esito:
non passato

livello:
basso

tempo:
3 minuti (fermato per superamento limite di tempo)

percorso:
 - home page
 - click su "Notizie ai viaggiatori e previsioni traffico"
 - click sulla data richiesta
 - non è sicura del nome dell'autostrada/del fatto che riguardi la tratta che le interessi
 - mappa del sito

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
1 minuto (supero limite tempo)

percorso:
 - click su "Vedi tutte le telecamere" (sotto la mappa)
 - mette la mappa fullscreen
 - identifica il punto richiesto
 - torna nella home page
 - click su "Telecamere" nella navbar
 - usa il widget con l'anteprima delle telecamere sotto la mappa cercando quella che le interessa

problemi:
minore: lo scrolling che triggera lo zoom della mappa.

opinioni:
Non ha riconosciuto le icone sulla mappa come cliccabili, e quindi anche
se in relativamente poco tempo (~35 secondi) ha raggiunto la pagina
corretta, è poi tornata nella home per cercare un altro percorso.


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

opinioni:

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
importante: zoom della mappa durante lo scrolling.

opinioni:

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
Trova che la home page contenga "troppe cose e non sono ben mostrate".

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

problemi:
importante: fatica nell'individuare la pagina

opinioni:
Suggerisce l'aggiunta di un link a sè stante nel menù per la pagina meteo. Ha avuto difficoltà a individuare la pagina in cui sono presenti le previsioni meteo.

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
importante: scegliere l'autostrada richiesta

opinioni:
Afferma che sarebbe meglio, invece di divere le autostrade in solo 3 blocchi, dividerle in tratte tra caselli.

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
importante: Lo zoom della mappa

opinioni:

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
importante: Lo zoom della mappa

opinioni:

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

opinioni:
Suggerisce che, per ogni tratta, si evidenzi la stazione di servizio con i costi minori.
