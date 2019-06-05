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

 - i video pre-registrati non hanno una trascrizione testuale equivalente
   (l'articolo che discute un evento mostrato nel video)

 - al momento nel sito non sono presenti contenuti di solo audio.

# Utilizzabile

# Comprensibile

# Robusto

# Conformità
