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

# Utilizzabile

# Comprensibile

# Robusto

# Conformità
