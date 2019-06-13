title:WCAG

# Percepibile

## Alternative testuali

### Contenuti non testuali

livello: A

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
outcome: partial

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
outcome: partial

Nella maggior parte dei testi i livelli di contrasto sono adeguati, ma ci sono alcuni elementi della pagina che non hanno un livello di contrasto adeguato come:
 - tasto condividi e numero di commenti nelle preview degli articoli nella home page e nelle pagine di archivio
 - l'ora e il punteggio dei commenti
 - i tasti per cambiare giornata nel meteo
 - i titoli delle preview degli articoli nella sezione "Top Blog"

### Ridimensionamento del testo
livello: AA
source: resize-text
outcome: pass

Il sito rimane funzionante una volta impostato il livello di zoom a 200%.

Ci preme comunque sottolineare che nonostante sia funzionante sono presenti delle fastidiose barre di scrolling orizzontali: sarebbe opportuno modificare il CSS in modo tale che le pagine rimangano larghe al più quanto il monitor.

### Immagini di testo
livello: AA
source: images-of-text
outcome: partial

Vengono usate solo un paio di immagini testuali (non essenziali e/o personalizzabili) quando un testo sarebbe stato più opportuno, come:
 - box "abbonati" nella homepage
 - logo "meteo" dove la scritta non è necessario sia inclusa nell'immagine

# Utilizzabile

# Comprensibile

# Robusto

# Conformità
