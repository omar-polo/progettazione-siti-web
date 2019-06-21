title:WCAG

# Percepibile

## Alternative testuali

### Contenuti non testuali
livello: A
outcome: no

violazioni:
 - manca una descrizione testuale per l'icona (font) cerca nella barra in alto del titolo
 - mancano svariati attributi "alt" alle immagini (in particolare le foto degli articoli)

   al momento, le foto senza alt sono 64 (poco più della metà)

```js
// lista delle immagini alle quali manca l'attributo alt
Array.prototype.filter.call(document.querySelectorAll('img'), i => !i.alt)
```

## Media temporizzati

> Fornire alternative per i media temporizzati

### solo audio e solo titolo
livello: A

 - i video pre-registrati non hanno una trascrizione testuale equivalente
   (l'articolo che discute un evento mostrato nel video)

 - al momento nel sito non sono presenti contenuti di solo audio.

### sottotitoli (preregistrati)
livello: A

I video presenti non hanno sottotitoli

### audiodescrizione o tipo di media alternativo (preregistrato)
livello: A

nada

### sottotitoli (in tempo reale)
livello: AA

al momento sono presenti solo video preregistrati, quindi questo punto
non si applica.

### audiodecrizione (preregistrata)
livello: AA

nada

### lingua dei segni
livello: AAA

nada

### audiodescrizione estesa (preregistrata)
livello: AAA

nada

### tipo di media alternativo (preregistrato)
livello: AAA

nada

### solo audio (in tempo reale)
livello: AAA

nada

## Adattabile

### Informazioni e correlazioni
livello: A
source: info-and-relationships
outcome: no

TODO: prendere la lista degli errori dal validator.w3.org

TL;DR rispettato solo in parte

 - aria: non vengono usati
 - G115 semantic markup: usato solo in parte

  viene usato un markup semantico in alcuni punti (es: aside) ma si
  fallisce miseramente in altri punti della pagina (es. molti titoli
  sono un `p` dentro un `h2`)

 - H49 usare il markup semantico per enfatizzare testo speciale: in
   parte. usano tag come `strong` per il testo in grossetto, ma non usano
   correttamente i `quote` per le citazioni.

 - G117 violato: i contenuti nuovi (aggiornamenti di articoli e/o nuovi
   articoli) non vengono segnalati.

 - G140 violato: si fa largo uso di attributi `style` nell'html per
   cambiare proprietà visuali degli elementi (come colore del testo
   e allineamento)

### Sequenza significativa
livello: A
source: meaningful-sequence

non applicabile: la lista nella quale gli articoli vengono mostrati non
influisce sul significato dei singoli articoli.

### caratteristiche sensoriali
livello: A
source: sensory-characteristics

I video e le foto non hanno descrizioni testuali, quindi questo punto è violato.

### Orientamento
livello: AA
source: orientation

Rispettato (il sito funziona su monitor di diverse forme e orientamenti.)

### Identificare lo scopo degli input
livello: AA
source: identify-input-purpose

in alcuni punti no:
 - il campo di ricerca non ha nessuna descrizione/label/... per permettere all'utente di identificare lo scopo dell'input.
 - alla textarea per i commenti manca un label/aria-label/aria-labelledby. In javascript viene aggiunta una scritta "Partecipa alla discussione" ma non sono presenti indicazioni per collegare tale scritta alla texarea

in alcuni punti male:
 - il campo di ricerca nella pagina meteo mette la scritta "Inserisci qui la tua località" come `value` dei un input di tipo testo. Bisognerebbe usare un label/aria-label/aria-labelledby

in altri casi sì:
 - gli input per il login hanno sia un placeholder che un label associato.

### Identificare lo scopo
livello: AAA
source: identify-purpose
outcome: no

 - utilizzano alcuni tag HTML5 correttamente (footer, header, nav, aside..) ma manca un tag main per indentificare il contenuto primario della pagina. Inoltre, i form non sono strutturati correttamente (vedi punti precedenti) e quindi non tutte le aree della pagina sono comprensibili.

## Distinguibile

### Uso del colore
livello: A
source: use-of-color
outcome: pass

Il colore non è l'unico modo con il quale le differenti aree della pagina sono divise. Vengono usati appropriati tag HTML5 (section e/o aside) per differenziare le aree anche da un punto di vista semantico.

### Controllo del sonoro
livello: A
source: audio-control
outcome: pass

Sono presenti video con riproduzione automatica che forniscono la possibilità di modificare il livello del volume, di mettere in pausa e interrompere la riproduzione audiosonora.

### Contrasto (minimo)
livello: AA
source: contrast-minimum
outcome: no

Nella maggior parte dei testi i livelli di contrasto sono adeguati, ma ci sono alcuni elementi della pagina che non hanno un livello di contrasto adeguato come:
 - tasto condividi e numero di commenti nelle preview degli articoli nella home page e nelle pagine di archivio
 - l'ora e il punteggio dei commenti
 - i tasti per cambiare giornata nel meteo
 - i titoli delle preview degli articoli nella sezione "Top Blog"
 - nella sezione "FQ in edicola" (a scomparsa nell'header) tutte le scritte quando con hover hanno un livello di contrasto insufficiente

### Ridimensionamento del testo
livello: AA
source: resize-text
outcome: pass

Il sito rimane funzionante una volta impostato il livello di zoom a 200%.

Ci preme comunque sottolineare che nonostante sia funzionante sono presenti delle fastidiose barre di scrolling orizzontali: sarebbe opportuno modificare il CSS in modo tale che le pagine rimangano larghe al più quanto il monitor.

### Immagini di testo
livello: AA
source: images-of-text
outcome: no

Vengono usate solo un paio di immagini testuali (non essenziali e/o personalizzabili) quando un testo sarebbe stato più opportuno, come:
 - box "abbonati" nella homepage
 - logo "meteo" dove la scritta non è necessario sia inclusa nell'immagine
 - la scritta 'vignetta di' nel box della vignetta di Natangelo non è necessaria sia inclusa nell'immagine (mentre la scritta 'Natangelo', in quanto firma, rientra nella categoria delle essenziali)

### Contrasto (avanzato)
livello: AAA
source: contrast-enhanced
outcome: no

alcuni esempi:
 - nella pagina di login:
    - la stellina per il richiesto non ha un contrasto sufficiente
    - tutti i pulsanti
    - la scritta 'non sei ancora registrato...'
 - nella pagina meteo (oltre a quelli citati nel punto 'Contrasto'):
    - il pulsante accedi
    - la scritta 'modifica località'
    - tutti i pulsanti per selezionare le varie tipologie (Italia, Europa, Mondo, ...)
    - il selettore dell'ora
 - nella pagina di un articolo
    - il numero dei commenti
    - le categorie nell'aside

### Sottofondo sonoro basso o non presente
livello: AAA
source: low-or-no-background-audio
outcome: pass

non esistono pagine con solo audio preregistrato e i video con autoplay sono mutati in partenza.

### Presentazione visiva
livello: AAA
source: visual-presentation
outcome: no

l'unico punto passato è 'testo non giustificato'

 - [ ] il sito non permette di cambiare i colori in primo piano e dello sfondo
 - [ ] la larghezza del testo dipende dalla dimensione della finestra e non è possibile imporre un limite a 80 caratteri
 - [x] il testo non è giustificato
 - [x] lo spazio tra le righe è almeno di 1.5 e lo spazio tra paragrafi è almeno una volta e mezzo l'interlinea
 - [ ] non è possibile ingrandire il testo e se si zooma con il browser al 200% compaiono delle barre di scorrimento orizzontali

### Immagini di testo (senza eccezioni)
livello: AAA
source: images-of-text-no-exception
outcome: no

No anche per i punti elencati in precedenza.

### Ricalcolo del flusso
livello: AA
source: reflow
outcome: yes

 - [x] il contenuto a scorrimento verticale con una larghezza equivalente a 320 CSS pixel non richiede di scorrere in due dimensioni
 - [x] non c'è contenuto a scorrimento orizzontale

a 400% il sito è usabile e rispetta di sicuro le richieste, ma intorno al 200% bisogna scrollare orizzontalmente e testi importanti come titoli, sottotitoli ecc sono "tagliati fuori"

### Contrasto in conenuti non testuali
livello: AA
source: non-text-contrast
outcome: pass

tutte le immagini testuali hanno un livello di contrasto adeguato, oppure sono loghi.

### Spaziatura del testo
livello: AA
source: text-spacing
outcome: no

 - [x] l'altezza della linea è di 1.5 volte la dimensione del carattere
 - [ ] la spaziatura tra i caratteri è pari alla dimensione del carattere (dovrebbe essere il doppio)
 - [ ] a discrezione dell'UA
 - [ ] a discrezione dell'UA

### Contenuto con Hover o Focus
livello: AA
source: content-on-hover-or-focus
outcome: pass

 - [x] congedabile: tutte le aree a scomparsa oscurano parte dei contenuti
 - [x] passibile: il contenuto aggiuntivo, una volta comparso per via di hover, rimane visibile se il puntatore (o il focus da tastiera) si spostano all'interno dell'area comparsa
 - [x] persistente: il contenuto aggiuntivo rimane visibile fino a quando l'evento hover o focus non viene rimosso

# Utilizzabile

## Accessibile da tastiera

Rendere disponibili tutte le funzionalità tramite tastiera

### Tastiera
livello: A
source: keyboard
outcome: no

Le voci di menu a scomparsa nell'intestazione (ad esempio 'FQ IN EDICOLA') non sono navigabili da tastiera in quanto l'area a scomparsa non si apre.

### Nessun impedimento all'uso della tastiera
livello: A
source: no-keyboard-trap
outcome: yes

Il focus da tastiera può essere spostato da/verso qualunque componente che supporta il focus della pagina.

### Tastiera (nessuna eccezione)
livello: AAA
source: keyboard-no-exception
outcome: no

Non lo è per lo stesso motivo del due sopra. (vedi sopra x2).

### Tasti di scelta rapida
livello: A
source: character-key-shortcuts
outcome: yes

Non hanno implementato nessuna scorciatoia da tastiera quindi questo punto è banalmente rispettato.

## Adeguata disponibilità di tempo

Fornire agli utenti tempo sufficiente per leggere e utilizzare i contenuti.

### Regolazione tempi di esecuzione
livello: A
source: timing-adjustable
outcome: no

La home page si auto-ricarica e:
 - non è possibile disattivarlo
 - non è possibile regolarlo
 - l'utente non viene avvisato prima dello scadere del tempo e non è possibile estendere il limite
 - il limite di tempo non è un evento fondamentale
 - il limite di tempo non è essenziale per l'attività
 - il limite di tempo è inferiore alle 20 ore (sono 10 minuti)

Bastarebbe rimpiazzare il metatag refresh con un controllo in javascript
che permetta almeno uno dei punti sopra citati.

### Pauso, stop, nascondi
livello: A
source: pasue-stop-hide
outcome: no

 - [x] i video che compaiono in sovraimpressione possono essere stoppati e/o nascosti
 - [ ] non è possibile disattivare l'autoaggiornamento

### Nessun tempo di esecuzione
livello: AAA
source: no-timing
outcome: no

Nonostante il refresh non sia una parte esseziale del contenuto, la home
page si auto-ricarica ogni 10 minuti.


### Interruzioni
livello: AAA
source: interruptions
outcome: no

La home page viene ricaricata ogni 10 minuti: tale azione non è
considerabile un'emergenza.

### Riautenticazione
livello: AAA
source: re-authenticating
outcome: yes

TODO: specificare che i commenti sono l'unica tipologia di informazione
inseribile dall'utente nelle pagine che abbiamo testato.

Se uno scrive un commento, la sessione scade, il commento non viene
perso ma viene riproposto tale e quale dopo un login.

#### Termine del tempo
livello: AAA
source: timeouts
outcome: yes

L'inattività non permette di perdere i commenti.

## Convulsioni e reazioni fisiche

Non sviluppare contenuti con tecniche che sia noto causino attacchi
epilettici o reazioni fisiche

### Tre lampeggiamenti o inferiore alla soglia
livello: A
source: three-flashes-or-below-threshold
outcome: yes

Le pagine esaminate non contengono nulla che lampeggi per più di
tre volte in un secondo oppure il lampeggiamento è al di sotto della
soglia. Va sottolineato però che i video che compaiono a lato cambiano
spesso: durante la nostra ispezione i video non contenevano elementi che
potessero invalidare questo punto, ma nulla vieta che in futuro ci siano.

### Tre lampeggiamenti
livello: AAA
source: three-flashes
outcome: yes

La pagina web non contiene nulla che lampeggi per più di tre volte in
un secondo. Si faccia riferimento comunque al punto precedente per il
discorso dei video.


### Animazione da interazioni
livello: AAA
source: animation-from-interaction
outcome: no

Per menu a scomparsa delle sezioni e i video a scomparsa non è possibile
disabilitare l'animazione (transazione) e non è essenziale.

## Navigabile

Fornire delle funizonalità di supporto all'utente per navigar, trovare
contenuti e determinare la propria posizione.

### Salto di blocchi
livello: A
source: bypass-blocks
outcome: no

Non è presente nessun meccanismo per saltare i blocchi di contenuto
che si ripetono su piu pagine.

### Titolazione della pagina
livello: A
source: page-titled
outcome: yes

Le pagine web hanno titoli che ne descrivono l'argomento.

### Ordine del focus
livello: A
source: focus-order
outcome: yes

Rispettato banalmente: la pagine viene navigata in modo sequenziale
usando il tab ma l'ordine di navigazione non influisce sul significato.

### Scopo del collegamento
livello: A
source: link-purpose-in-context
outcome: no

Mancano dei aria-label sui box del meteo e della vignetta in homepage. Le
altre pagine sono a posto.

### Differenti modalità
livello: AA
source: multiple-ways
outcome: yes

Questo punto è rispettato, confrontando con la lista di tecniche
sufficienti per il successo del criterio:

 - [x] forniscono link per navigare a pagine correlati
 - [ ] non forniscono una TOC
 - [ ] non hanno una site map
 - [x] forniscono una funzionalità di ricerca
 - [ ] non forniscono una lista di link a tutte le altre pagine
 - non tutte le pagine sono linkate dalla home page

### Intestazioni ed etichette
livello: AA
source: headings-and-labels
outcome: yes

Le sezioni esaminate hanno intestazioni che ne specificano il contenuto:
 - nella pagina di un articolo l'h1 descrive l'articolo, e gli h2 introducono gli articoli correlati
 - nella pagina del meteo gli h2 introducono i riquadri a lato
 - nella pagina di login l'h3 introduce lo scopo della pagina

TODO: nella pagina di login ci dovrebbe essere un h1 però...

### Focus visibile
livello: AA
source: focus-visible
outcome: yes

Le parti dell'interfaccia vengono evidenziate dallo user agent quando
ricevono focus.

### Posizione
livello: AAA
source: location
outcome: yes

Nelle pagine di archivio e degli articoli sono presenti delle
*breadcrumb*. Nella pagina del meteo la posizione corrente viene mostrata
nella nav.

### Scopo del collegamento (solo collegamento)
livello: AAA
source: link-purpose-link-only
outcome: no

Nella homepage il box del meteo è un'immagine senza testo descrittivo,
così come anche il box della vignetta.

### Intestazioni di sezione
livello: AAA
source: section-headings
outcome: yes

I titoli vengono usati correttamente (ma si veda comunque il punto
[Informazioni e correlazioni](#informazioni-e-correlazioni) per gli
errori semantici)

## Modalità di input

Rendere più facile agli utenti l'utilizzo di funzionalità attraverso
input diversi dalla tastiera.

### Movimenti del puntatore
livello: A
source: pointer-gestures
outcome: yes

Non sono richiesti gisti multi punto o basati su percorsi.

### Cancellazione delle azioni del puntatore
livello: A
source: pointer-cancellation
outcome: yes

non sono presenti down-event nelle pagine esaminate.

### Etichetta nel nome
livello: A
source: label-in-name
outcome: no

 - il controllo di ricerca ha il nome ha nome 'q'
 - nella schermata di login i controlli hanno nome 'username' ma il testo mostrato è 'Nome utente o email'

### Azionamento da movimento
livello: A
source: motion-actuation
outcome: yes

Non sono presenti funzionalità che richiedano di essere attivate dal
movimento del dispositivo o dell'utente.

### Dimensione dell'obbiettivo
livello: AAA
source: target-size
outcome: no

Il link vuoto nella barra laterale non è alto abbastanza (0px)

### Meccanismi di input simultanei
livello: AAA
source: concurrent-input-mechanisms
outcome: yes

L'unica modalità di input richiesta è il click e l'hover.

# Comprensibile

Le informazioni e le operazioni dell'interfaccia utente devono essere
comprensibili.

## Leggibile

Rendere il testo leggibile e comprensibile.

### Lingua della pagina
livello: A
source: language-of-page
outcome: no

l'unica violazione è la pagina del meteo che non ha un attributo lang
sull'elemento html, non ha un meta Content-Language e, quando viene
servita, non è presente  un header Content-Language.

### Parti in lingua
livello: AA
source: language-of-parts
outcome: yes

Negli articoli esaminati, il testo è sempre in italiano fatta eccezione
per i nomi propri, alcune terminologie tecniche od alcune espressioni
diventate parte integrante del parlato.

### Parole inusuali
livello: AAA
source: unusual-words
outcome: no

Non è presente nessun meccanismo per identificare parole o frasi usate
in modo insolito o ristretto.

### Abbreviazioni
livello: AAA
source: abbreviations
outcome: no

Non è presente nessun meccanismo per identificare la forma espansa o
il significato delle abbreviazioni.

### Livello di lettura
livello: AAA
source: reading-level
outcome: yes

Negli articoli esaminati il testo non ci sembra richieda capacità
di lettura più avanzata rispetto al livello di istruzione secondaria
inferiore.

### Pronuncia
livello: AAA
source: pronunciation
outcome: yes

Si perchè il contenunto prevalente è in italiano e non abbiamo i *false friends*.

## Prevedibile

Creare pagine web che abbiano aspetto e funzionalità prevedibili.

### Al Focus
livello: A
source: on-focus
outcome: yes

Nelle pagine esaminate quando un componente dell'interfaccia utente
riceve il focus non si avvia nessun cambiamento del contesto.

### All'input
livello: A
source: on-input
outcome: yes


TODO: tenere a mente quando si progetta il campo di ricerca

### Navigazione Coerente
livello: AA
source: consistent-navigation
outcome: yes

Gli unici meccanismi di navigazioni che sono ripetuti su più pagine
web sono i seguenti e appaiono sempre nello stesso ordine relativo:
 - menu a scomparsa laterale
 - link del footer

### Identificazione coerente
livello: AA
source: consistent-identification
outcome: yes

I componenti che hanno la stessa funzionalità (come i controlli 'Articolo
successivo/precedente', il menu a scomparsa, il footer) sono identificati
consistentemente.

### Cambiamenti su richiesta
livello: AAA
source: change-on-demand
outcome: no

La home page ha una funzione di auto-refresh attiva che periodicamente
aggiorna la pagina. Questo provoca un cambiamento di contesto (la
posizione relativa all'interno della pagina viene persa, così come
anche il focus).

## Assistenza nell'inserimento

Aiutare gli utenti a evitare gli errori e agevolarli nella loro correzione.

### Identificazione di errori
livello: A
source: error-identification
outcome: no

Provare a postare un commento (vuoto) da loggati non risulta in nessuna
forma di feedback che permetta all'utente di capire che l'azione non è
permessa. Non viene identificato nessun elemento e non compare nessun
testo.

Nella pagina di login invece, se vengono fornite credenziali errate,
viene fornito un messaggio di errore.

### Etichette o istruzioni
livello: A
source: labels-or-instructions
outcome: no

Nonostante nella schermata di login ci siano appropriati etichette, la
textarea per i commenti non ha etichette adeguate (è presente una scritta
in sovraimpressione ma viene aggiunta con javascript e posizionata lì
sopra, non ha nessuna vicinanza all'elemento che permetta di comprenderne
l'associazione se non il risultato grafico).

### Seggerimenti per gli errori
livello: AA
source: error-suggestion
outcome: no

Si veda ad esempio la già citata textarea per i commenti.

### Prevenzione degli errori
livello: AA
source: error-prevention-legal-financial-data
outcome: yes

Nelle pagine controllate non sono presenti pagine con vincoli di tipo
giuridico, transazioni finanziarie o la modifica/cancellazione o gestione
di dati controllabili dall'utente.

### Aiuto
livello: AAA
source: help
outcome: yes

 - è presente una pagina (seppur limitata) di aiuto (linkata sempre nel footer)
 - ci sono degli aiuti contestuali

# Robusto

Il contenuto deve essere abbastanza robusto per essere interpretato in
maniera affidabile da una grande varietà di programmi utente, comprese
le tecnologie assitive.

## Compatibile

Garantire la massima compatibilità con i programmi utente attuali e
futuri, comprese le tecnologie assistive.

### Analisi sintattica (parsing)
livello: A
source: parsing
outcome: no

copiare dal validatore:
 - apertura e chiusura corretta dei tag
 - attributi non duplicati
 - id univoci

sappiamo per certo che i tag non sono corretti (chiusi due volti i p
per esempio).

### Nome, ruolo, valore
livello: A
source: name-role-value
outcome: no

Non è possibile calcolare name e ruolo per ogni componente
dell'interfaccia utente (verificato con gli strumenti di accessibilità
di firefox: la barra di ricerca per esempio non viene riconosciuta).

### Messaggi di stato
livello: AA
source: status-messages@
outcome: no

Non sono presenti messaggi di stato che possano essere determinati
programmaticamente.

Infatti, abbiamo visto come per moltissime componenti della pagina non
sia possibile stabilire programmaticamente il ruolo.

# Conformità
