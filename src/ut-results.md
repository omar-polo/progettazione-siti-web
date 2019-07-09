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
età: 25
professione: studente
frequenza-web: 12 ore al giorno
siti-viaggi: yes
infoviaggiando: no
esperienza: 4

### Task 1

esito:
non passato

livello:
basso

tempo:
4 minuti 12 secondi

percorso:
TODO

problemi:
tanti

opinioni:
poche

### Task 2

esito:
non passato

livello:
basso

tempo:
45 secondi

percorso:
TODO

problemi:
tanti

opinioni:
poche

### Task 3

esito:
non passato

livello:
basso

tempo:
1 minuto e 42 secondi

percorso:
TODO

problemi:
tanti

opinioni:
poche

### Task 4

esito:
non passato

livello:
basso

tempo:
10 secondi

percorso:
TODO

problemi:
tanti

opinioni:
poche

### Task 5

esito:
non passato

livello:
basso

tempo:
30 secondi

percorso:
TODO

problemi:
tanti

opinioni:
poche

### Task 6

esito:
non passato

livello:
basso

tempo:
1 minuto e 15 secondi

percorso:
TODO

problemi:
tanti

opinioni:
poche

## utente 2
sesso: maschio
età: 22
professione: studente
frequenza-web: 15-16 ore al giorno
siti-viaggi: yes
infoviaggiando: no
esperienza: 4

### Task 1

esito:
non passato

livello:
basso

tempo:
30 secondi

percorso:
TODO

problemi:
tanti

opinioni:
poche

### Task 2

esito:
non passato

livello:
basso

tempo:
1 minuto e 45 secondi

percorso:
TODO

problemi:
tanti

opinioni:
poche

### Task 3

esito:
non passato

livello:
basso

tempo:
1 minuto e 50 secondi

percorso:
TODO

problemi:
tanti

opinioni:
poche

### Task 4

esito:
non passato

livello:
basso

tempo:
30 secondi

percorso:
TODO

problemi:
tanti

opinioni:
poche

### Task 5

esito:
non passato

livello:
basso

tempo:
2 minuti e 15 secondi

percorso:
TODO

problemi:
tanti

opinioni:
poche

### Task 6

esito:
non passato

livello:
basso

tempo:
3 minuti e 4 secondi

percorso:
TODO

problemi:
tanti

opinioni:
poche
