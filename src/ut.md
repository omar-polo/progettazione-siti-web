title: User Testing

## Introduzione

cose...

## Perimetro dell'indagine

## Sommario dei risultati

## Metodologia

### Briefing

Domande del questionario:

 - Età
 - Professione
 - Competenze digitali
 - Frequenza di utilizzo del web
 - Hai mai usato un sito di informazioni sui viaggi?
 - Hai mai usato il sito "infoviaggiando.it"?

### Profili utente

{{ range .Surveys }}
#### Utente x

TODO: {{ . }}
{{ end }}

### Situazione operativa

Dove e come sono stati effettuati i test.

Per farla breve:

 - giardino dei rizzi
 - all'ombra
 - situazione silenziosa, non c'erano persone nei dintorni
 - pc di daniele (inserire le spec)

I partecipanti sono stati invitati a pensare ad alta voce, per compredere
come ragione e le motivazioni di una certa azione. In alcune occasioni,
sono state poste delle domande, prive di indizi, all'utente per stimolarlo
a parlare. Nei momenti di difficoltà, lo sperimentatore ha calmato
l'utente invitandolo a provare a proseguire.

### Indici di valutazione

#### Indice di successo

 - basso: <!-- TODO definizione -->
 - medio: <!-- TODO definizione -->
 - alto: <!-- TODO definizione -->

#### Indice di gravità del problema

impatto:

 1. il problema riscontrato dall'utente è facilmente superabile
 2. il problema richiede un notevol sforzo per essere aggirato
 3. il problema blocca la possibilità di completare il compito

frequenza:
 1. il problema si verifica poche volte
 2. il problema si verifica ogni volta durante lo svolgimento del compito
 3. il problema si verifica molte volte durante lo svolgimento del compito

In base all'impatto e alla frequenza si definisce la gravità:

<!-- TODO: forse questa tabella la rifacciamo in HTML per poterci mettere
           un po' di stile -->

| gravità    | impatto | frequenza |
|------------|--------:|----------:|
| minore     |       1 |         1 |
| minore     |       1 |         2 |
| importante |       1 |         3 |
| importante |       2 |         1 |
| importante |       2 |         2 |
| critico    |       2 |         3 |
| critico    |       3 |         1 |
| critico    |       3 |         2 |
| critico    |       3 |         3 |

## Risultati


{{ range $i, $v := .Results }}
{{ with $v }}
### Caso d'uso {{ $i | inc }}

{{ .Task.Background }}

Tempo previsto: {{ .Task.Timing }}

Percorso previsto: {{ .Task.Path }}

{{ range $j, $u := .Users }}
{{ with $u }}
#### Utente {{ $j | inc }}

Esito: {{ .Outcome }}

Tempo impiegato: {{ .Timing }}

Livello di successo: {{ .Level }}

Percorso dell'utente: {{ .Path }}

Problemi: {{ .Problems }}

Opinioni del partecipante: {{ .Opinions }}

{{ end }} <!-- with $u -->

{{ end }} <!-- range .Users -->

{{ end }} <!-- with $v -->

{{ end }} <!-- range .Results -->
