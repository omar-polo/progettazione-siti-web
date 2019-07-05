title:Linee guida per l'accessibilità dei contenuti Web (WCAG) 2.1

# Linee guida per l'accessibilità dei contenuti Web (WCAG) 2.1

La valutazione di conformità si è basata sulle linee guida **WCAG 2.1** che prevedono diversi livelli di orientamento che comprendono *principi* globali, *linee* guida generali, *criteri di successo* verificabili e una ricca raccolta di tecniche *sufficienti* e *consigliate*.

- **Principi** - Al livello più alto, sono definiti i quattro principi che fanno da pilastri all'accessibilità del Web: *percepibile*, *utilizzabile*, *comprensibile* e *robusto*. 

- **Linee guida** - Dai quattro principi derivano tredici linee guida. Le linee guida forniscono gli obiettivi di base su cui gli autori dovrebbero lavorare per rendere il contenuto più accessibile agli utenti con disabilità diverse. 

- **Criteri di successo** - Per ogni linea guida, sono forniti criteri di successo verificabili per consentire l'utilizzo delle WCAG 2.1 dove sono necessari test dei requisiti e della conformità. 
  - Vengono definiti tre livelli di conformità: A (minimo), AA e AAA (massimo).

- **Tecniche sufficienti e consigliate** -  Le tecniche sono informative e possono essere di due categorie: quelle sufficienti a soddisfare il criterio di successo e quelle consigliate. Le tecniche consigliate vanno oltre ciò che viene richiesto da ciascun singolo criterio di successo e consentono di rispettare meglio le linee guida. 

Tutti questi livelli di orientamento (principi, linee guida, criteri di successo, tecniche sufficienti e consigliate) concorrono a fornire indicazioni su come rendere il contenuto più accessibile. 

Sono stati utilizzati diversi strumenti come supporto alla valutazione di conformità alle linee guida WCAG 2.1.

- I browser Web [Mozilla Firefox Developer Edition](https://www.mozilla.org/it/firefox/channel/desktop/) e [Google Chrome](https://www.google.com/intl/it/chrome/)

- Markup Validation Service [W3C](https://validator.w3.org/)

- CSS Validation Service [W3C](https://jigsaw.w3.org/css-validator/)

- Web Accessibility Evaluation Tool [WAVE](http://wave.webaim.org/)

- Color Contrast Accessibility Validator [a11y](https://color.a11y.com/?wc3)

- [Accessibility Inspector](https://developer.mozilla.org/it/docs/Tools/Accessibility_inspector?utm_source=devtools&utm_medium=a11y-panel-description ) che fornisce un mezzo per accedere alle informazioni importanti esposte alle tecnologie assistive nella pagina corrente tramite l'albero di accessibilità. 

- [Lighthouse](https://developers.google.com/web/tools/lighthouse/) che è uno strumento automatizzato open-source per migliorare la qualità delle pagine web. Ha Audits per prestazioni, accessibilità, buone pratiche, applicazioni web progressive e altro.

## Percepibile

> Le informazioni e i componenti dell'interfaccia utente devono essere presentati agli utenti in modi in cui essi possano percepirli.

### Alternative testuali

> Fornire alternative testuali per qualsiasi contenuto non di testo in modo che questo possa essere trasformato in altre forme fruibili secondo le necessità degli utenti come stampa a caratteri ingranditi, Braille, sintesi vocale, simboli o un linguaggio più semplice.

#### Contenuti non testuali
livello: A
outcome: no
source: non-text-content

> Tutti i contenuti non testuali presentati all'utente hanno un'alternativa testuale equivalente che serve allo stesso scopo, ad eccezione dei seguenti casi
>
> **Controlli**, **input**: Se il contenuto non testuale è un controllo o accetta l'input degli utenti, allora ha un nome che ne descrive la finalità.
>
> **Media temporizzati**: Se il contenuto non testuale è un media temporizzato, allora le alternative testuali forniscono almeno una identificazione descrittiva per il contenuto non testuale.
> **Test**: Se il contenuto non testuale è un test o un esercizio che potrebbe essere non valido se presentato come testo, allora le alternative testuali forniscono almeno una descrizione identificativa del contenuto non testuale.
>
> **Esperienze sensoriali**: Se il contenuto non testuale ha lo scopo primario di creare una specifica esperienza sensoriale, allora le alternative testuali forniscono almeno una descrizione identificativa del contenuto non testuale.
>
> **CAPTCHA**: Se la finalità del contenuto non testuale è confermare che il contenuto sia utilizzato da una persona e non da un computer, allora sono fornite alternative testuali che identifichino e descrivano lo scopo del contenuto non testuale, e forme alternative di CAPTCHA che usino diverse modalità di output per differenti tipologie di percezioni sensoriali al fine di soddisfare differenti disabilità.
>
> **Decorazioni**, **formattazioni**, **contenuti invisibili**: Se il contenuto non testuale è puramente decorativo, è utilizzato solamente per formattazione visuale oppure non è presentato agli utenti, allora è implementato in modo da poter essere ignorato dalla tecnologia assistiva.

##### Violazioni riscontrate:

Manca una descrizione testuale per l'icona "search" nella barra di navigazione.

Non sono presenti gli attributi "alt" che forniscono una descrizione testuale dell'immagine, la quale ne descrive i contenuti per chi non può vederla (in particolare nelle foto degli articoli).

Al momento, le immagini senza l'attributo "alt" sono 64.

Abbiamo utilizzato questo script JavaScript per ricavare la lista delle immagini senza l'attributo "alt"

```javascript
  // Lista delle immagini alle quali manca l'attributo alt
  Array.prototype.filter.call(document.querySelectorAll('img'), i => !i.alt)
```

**Screen d'esempio**

![Esempio di icone senza descrizione testuale](img/contenuti-non-testuali.png)

![Esempio di icone senza descrizione testuale](img/../../img/social-non-testuali.png)

**Code snippets**

```html
  <!-- Logo Home page -->
  <img src="https://st.ilfattoquotidiano.it/wp-content/themes/ifq/assets/img/logo-header-navbar.png" alt="">

  <!-- Articoli -->
  <img class="attachment-primopiano img-landscape" src="https://st.ilfattoquotid…/2015/04/finanza-990.jpg" alt="" width="990" height="192">
  <img class="lazyload" data-src="https://st.ilfattoquotid…adalajara675-320x132.jpg" alt="" width="320" height="132">​
  <img class="lazyload" data-src="https://st.ilfattoquotid…2/bersani675-320x132.jpg" alt="" width="320" height="132">​
  <img class="lazyload" data-src="https://st.ilfattoquotid…2/viagola675-320x132.jpg" alt="" width="320" height="132">​
  <img class="lazyload" data-src="https://st.ilfattoquotid…vecchia-1300-320x132.jpg" alt="" width="320" height="132">
  <img class="lazyload" data-src="https://st.ilfattoquotid…vecchia-1300-320x132.jpg" alt="" width="320" height="132">

  <!-- Social Icons -->
  <a href="http://www.facebook.com/ilFattoQuotidiano" target="_blank" itemprop="sameAs">
  <i class="icon-fb"></i>Facebook</a>
  <a href="http://www.twitter.com/FattoQuotidiano" target="_blank" itemprop="sameAs">
  <i class="icon-tw"></i>Twitter</a>
  <a href="http://www.youtube.com/antefattoblog" target="_blank" itemprop="sameAs">
  <i class="icon-yt"></i>YouTube</a>
  <a href="https://www.instagram.com/ilfattoquotidianoit/" target="_blank" itemprop="sameAs">
  <i class="icon-instagram"></i>Instagram</a>
```

### Media temporizzati

> Fornire alternative per i media temporizzati

#### Solo audio e solo video (preregistrati)
livello: A
outcome: no
source: audio-only-and-video-only-prerecorded

> Per i tipi di media preregistrati di solo audio e di solo video, a meno che questi non costituiscano un tipo di media alternativo ad un contenuto testuale chiaramente etichettato come tale, sono soddisfatti i seguenti punti:

  > **Solo audio preregistrato**: È fornita un'alternativa per il tipo di media temporizzato che presenti informazioni equivalenti al contenuto di solo audio preregistrato.

  > **Solo video preregistrato**: È fornita un'alternativa per il tipo di media temporizzato oppure una traccia audio che presenti informazioni equivalenti al contenuto di solo video preregistrato.

##### Violazioni riscontrate:

I video pre-registrati non hanno una trascrizione testuale equivalente.

Al momento della nostra analisi, nel sito non sono presenti contenuti di solo audio.

#### sottotitoli (preregistrati)
livello: A
outcome: no
source: captions-prerecorded

> Per tutti i contenuti audio preregistrati presenti in tipi di media sincronizzati sono forniti sottotitoli, eccetto quando tali contenuti sono alternativi ad un contenuto testuale e sono chiaramente etichettati come tali.

##### Violazioni riscontrate:

I video presenti nelle pagine del sito non hanno sottotitoli.

#### Audiodescrizione o tipo di media alternativo (preregistrato)
livello: A
outcome: no
source: audio-description-or-media-alternative-prerecorded

> Per i media sincronizzati è fornita un'alternativa ai media temporizzati, oppure una audiodescrizione dei contenuti video preregistrati, eccetto quando il contenuto audio o video è alternativo ad un contenuto testuale ed è chiaramente etichettato come tale.

##### Violazioni riscontrate:

Non sono state fornite alternative ai media temporizzati e non vi sono audiodescrizioni dei contenuti video preregistrati.

#### Sottotitoli (in tempo reale)
livello: AA
outcome: na
source: captions-live

> Per tutti i contenuti audio in tempo reale sotto forma di media sincronizzati sono forniti sottotitoli.

##### Violazioni riscontrate:

Al momento sono presenti solo video preregistrati, quindi questo punto
non si applica.

#### Audiodescrizione (preregistrata)
livello: AA
outcome: no
source: audio-description-prerecorded

> Per tutti i contenuti video preregistrati sotto forma di media sincronizzati è fornita una audiodescrizione.

##### Violazioni riscontrate:

Non sono presenti audiodescrizioni in nessun video preregistrato.

#### Lingua dei segni (preregistrato)
livello: AAA
outcome: na
source: sign-language-prerecorded

> Per tutti i contenuti audio preregistrati sotto forma di media sincronizzati è fornita l'interpretazione tramite lingua dei segni.

##### Violazioni riscontrate:

Non sono presenti contenuti audio preregistrati, quindi questo criterio di successo non si applica.

#### Audiodescrizione estesa (preregistrata)
livello: AAA
outcome: no
source: extended-audio-description-prerecorded

> Per tutti i contenuti video preregistrati in media sincronizzati, se le pause nell'audio principale sono troppo brevi per consentire alle audiodescrizioni di comunicare il senso del video, sono fornite delle audiodescrizioni estese.

##### Violazioni riscontrate:

Non sono presenti audiodescrizioni, quindi nemmeno audiodescrizioni estese.

#### Tipo di media alternativo (preregistrato)
livello: AAA
outcome: no
source: media-alternative-prerecorded

> Per tutti i contenuti preregistrati di media sincronizzati e per tutti i tipi di media preregistrati di solo video è fornito un tipo di media alternativo.

##### Violazioni riscontrate:

Non sono forniti media alternativi per nessuno dei media preregistrati di solo video.

#### Solo audio (in tempo reale)
livello: AAA
outcome: no
source: audio-only-live

> Per i media temporizzati che presentano informazioni equivalenti a contenuti solo audio in tempo reale è fornita un'alternativa.

##### Violazioni riscontrate:

Non è presente nessuna alternativa ai contenuti di solo audio in tempo reale.

### Adattabile

> Creare contenuti che possano essere rappresentati in modalità differenti (ad esempio, con layout più semplici), senza perdere informazioni o struttura.

#### Informazioni e correlazioni
livello: A
outcome: no
source: info-and-relationships

> Le informazioni, la struttura e le correlazioni trasmesse dalla presentazione possono essere determinate programmaticamente oppure sono disponibili tramite testo.

##### Violazioni riscontrate:

L'elemento *p* non può essere figlio di un *h2*

**Code snippets**

```html
<h2 class="title" data-vr-excerpttitle id="desktop-1-1-strilli-row1-1-secondopiano2col-title">
  <p>
    <a href="https://www.ilfattoquotidiano.it/2019/07/04/ndrangheta-le-mani-su-tre-parcheggi-di-malpensa-34-arresti-in-8-province-anche-un-consigliere-comunale-di-ferno/5300253/">Le mani della &#8216;ndrangheta sui parcheggi<br/>
    dell&#8217;aeroporto di Malpensa: 34 arrestati<br/>
    <span style="color: #dd0015;">Ombre su elezione ex sindaco di Lonate</span>
    </a>
  </p>
</h2>
```

L'elemento *p* non può essere figlio di un elemento di tipo *span*

**Code snippets**

```html
<span class="no-mobile">
  <p>Si è spesso discusso del rapporto dei trattati europei con la Costituzione italiana. Di solito,&hellip;</p>
</span>
```

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

#### Sequenza significativa
livello: A
outcome: na
source: meaningful-sequence

> Quando la sequenza in cui il contenuto è presentato influisce sul suo significato, la corretta sequenza di lettura può essere determinata programmaticamente.

##### Violazioni riscontrate:

La lista nella quale gli articoli vengono mostrati non influisce sul significato dei singoli articoli, quindi questo criterio di successo non si applica.

#### Caratteristiche sensoriali
livello: A
outcome: no
source: sensory-characteristics

> Le istruzioni fornite per comprendere ed operare sui contenuti non si basano unicamente su caratteristiche sensoriali dei componenti quali forma, colore, dimensione, ubicazione visiva, orientamento o suono.

##### Violazioni riscontrate:

I video e le foto non hanno descrizioni testuali e si basano unicamente su caratteristiche sensoriali, quindi questo punto non è rispettato.

#### Orientamento
livello: AA
outcome: yes
source: orientation

> La visualizzazione e il funzionamento di un contenuto non dipendono dall'orientamento dello schermo, ad esempio verticale o orizzontale, a meno che questo non sia essenziale.

##### Osservazioni

Il sito funziona su monitor di diverse forme e orientamenti, quindi questo punto è rispettato.

#### Identificare lo scopo degli input
livello: AA
outcome: no
source: identify-input-purpose

> Lo scopo di ciascun campo di input per le informazioni sull'utente può essere determinato programmaticamente quando:

> Il campo di input ha uno scopo noto, identificato nella sezione scopo dell'input per i componenti dell'interfaccia utente;
  
> Il contenuto è implementato utilizzando tecnologie che supportino l'identificazione del significato atteso dei dati inseriti del modulo.

##### Violazioni riscontrate:

Il campo di ricerca non ha nessuna descrizione/label per permettere all'utente di identificare lo scopo dell'input.

Alla textarea per i commenti manca una label/aria-label/aria-labelledby. In JavaScript viene inserita una scritta "Partecipa alla discussione", ma non sono presenti indicazioni per collegare tale scritta alla texarea.

Il campo di ricerca nella pagina meteo presenta la dicitura "Inserisci qui la tua località" come `value` di un input di tipo testo. Invece, bisognerebbe usare un label/aria-label/aria-labelledby.

#### Identificare lo scopo
livello: AAA
outcome: no
source: identify-purpose

> Nei contenuti implementati utilizzando i linguaggi di markup, è possibile determinare programmaticamente lo scopo dei componenti dell'interfaccia utente, delle icone e delle aree.

##### Violazioni riscontrate:

Alcuni tag HTML5 sono utilizzati correttamente (footer, header, nav, aside..), ma manga un tag 'main' per indentificare il contenuto primario della pagina. Inoltre, i form non sono strutturati correttamente, e quindi non tutte le aree della pagina sono comprensibili.

### Distinguibile

> Rendere più semplice agli utenti la visione e l'ascolto dei contenuti, separando i contenuti in primo piano dallo sfondo.

#### Uso del colore
livello: A
outcome: yes
source: use-of-color

> Il colore non deve essere utilizzato come unica modalità visiva per rappresentare informazioni, indicare azioni, richiedere risposte o come elemento di distinzione visiva.

##### Osservazioni

Il colore non è l'unico modo con il quale le differenti aree della pagina sono divise. Vengono usati appropriati tag HTML5 (section e/o aside) per differenziare le diverse aree anche da un punto di vista semantico.

#### Controllo del sonoro
livello: A
outcome: yes
source: audio-control

> Se un contenuto audio all'interno di una pagina Web è eseguito automaticamente per più di tre secondi o si fornisce una funzionalità per metterlo in pausa o interromperlo, oppure si fornisce una modalità per il controllo dell'audio che sia indipendente dal controllo predefinito del sistema.

##### Osservazioni

Sono presenti video con riproduzione automatica che forniscono la possibilità di modificare il livello del volume, di mettere in pausa e interrompere la riproduzione audiosonora.

#### Contrasto (minimo)
livello: AA
outcome: no
source: contrast-minimum

> La rappresentazione visiva del testo e di immagini contenenti testo ha un rapporto di contrasto di almeno 4.5:1, fatta eccezione per i seguenti casi:

  > **Testo grande**: Testo grande e immagini contenenti testo grande devono avere un rapporto di contrasto di almeno 3:1;

  > **Testo non essenziale**: Testo o immagini contenenti testo che siano parti inattive di componenti dell'interfaccia utente, che siano di pura decorazione, non visibili a nessuno, oppure che facciano parte di immagini contenenti contenuti visuali maggiormente significativi, non hanno alcun requisito di contrasto.

  > **Logotipi**: Un testo che sia parte di un logo o marchio non ha alcun requisito minimo di contrasto.

##### Violazioni riscontrate:

Nella maggior parte dei testi i livelli di contrasto sono adeguati, ma ci sono alcuni elementi della pagina che non hanno un livello di contrasto ottimale come:

 - Il tasto condividi e il numero di commenti nelle preview degli articoli della Home page e nelle pagine di archivio
  
 - L'ora e il punteggio dei commenti
  
 - I tasti per cambiare giornata nella pagina del meteo
  
 - I titoli degli articoli nella sezione "Top Blog"

#### Ridimensionamento del testo
livello: AA
outcome: yes
source: resize-text

> Il testo, ad eccezione dei sottotitoli e delle immagini contenenti testo, può essere ridimensionato fino al 200 percento senza l'ausilio di tecnologie assistive e senza perdita di contenuto e funzionalità.

##### Osservazioni

Il sito rimane funzionante una volta impostato il livello di Zoom a 200%.
Ci preme comunque sottolineare che nonostante sia funzionante sono presenti delle fastidiose barre di scrolling orizzontali: sarebbe opportuno modificare il CSS in modo tale che le pagine rimangano larghe al più quanto il monitor.

#### Immagini di testo
livello: AA
outcome: no
source: images-of-text

> Se le tecnologie utilizzate consentono la rappresentazione visiva dei contenuti, per veicolare informazioni è usato il testo, e non le immagini di testo, ad eccezione dei seguenti casi:

  > **Personalizzabile**: L'immagine di testo può essere personalizzata visivamente per le esigenze dell'utente;

  > **Essenziale**: Una particolare rappresentazione del testo è essenziale per il tipo di informazioni veicolate.

##### Violazioni riscontrate:

Vengono usate solo un paio di immagini testuali (non essenziali e/o personalizzabili), quando un testo sarebbe stato più opportuno, come:

 - Il box "Abbonati" nella Homepage
  
 - Il logo "Meteo" dove la scritta non è necessaria che sia inclusa nell'immagine
  
 - La scritta 'la vignetta di' nel box della vignetta di Natangelo non è necessaria che sia inclusa nell'immagine (mentre la scritta 'Natangelo', in quanto firma, rientra nella categoria delle essenziali).

#### Contrasto (avanzato)
livello: AAA
outcome: no
source: contrast-enhanced

> La rappresentazione visiva del testo e immagini contenenti testo ha un rapporto di contrasto di almeno 7:1, fatta eccezione per i seguenti casi:

  > **Testo grande**: Testo grande e immagini contenenti testo grande devono avere un rapporto di contrasto di almeno 4.5:1;

  > **Testo non essenziale**: Testo o immagini contenenti testo che siano parti inattive di componenti dell'interfaccia utente, che siano di pura decorazione, non visibili a nessuno, oppure che facciano parte di immagini contenenti contenuti visuali maggiormente significativi, non hanno alcun requisito di contrasto.

  > **Logotipi**: Un testo che sia parte di un logo o marchio non ha alcun requisito minimo di contrasto.

##### Violazioni riscontrate:

Alcuni esempi:

 - Nella pagina di login:
  
    - la stellina accanto alle voci 'Nome utente o email' e 'Password' non ha un contrasto sufficiente
  
    - tutti i pulsanti
  
    - la scritta 'Non sei ancora registrato?' non ha un contrasto adeguato
  
 - Nella pagina del meteo
  
    - il pulsante 'Accedi'
  
    - la scritta 'modifica località' è poco visibile
  
    - tutti i pulsanti per selezionare le varie zone del mondo (Italia, Europa, Mondo, ...)
  
    - il selettore dell'ora
  
 - Nella pagina di un articolo
  
    - il numero dei commenti
  
    - le categorie nell'aside

#### Sottofondo sonoro basso o non presente
livello: AAA
outcome: yes
source: low-or-no-background-audio

> Per i contenuti di solo audio preregistrato che (1) contengano principalmente parlato in primo piano (2) non siano CAPTCHA audio o loghi audio e (3) non siano una vocalizzazione intesa per essere principalmente espressione musicale come canto o rap, si applica almeno uno dei seguenti casi:

  > **Nessun sottofondo**: Il sonoro non contiene suoni di sottofondo.

  > **Spegnimento**: Il sottofondo sonoro può essere disattivato.

  > **20 dB**: Il sottofondo sonoro deve essere inferiore di almeno 20 decibel rispetto al parlato in primo piano, con l'eccezione di suoni occasionali di durata massima di uno o due secondi.

##### Osservazioni

Non esistono pagine con solo audio preregistrato e i video con autoplay sono muti in partenza, quindi questo punto è rispettato.

#### Presentazione visiva
livello: AAA
outcome: no
source: visual-presentation

> Per la presentazione visiva di blocchi di testo, è disponibile una modalità per conseguire i seguenti obiettivi:

  > I colori del testo in primo piano e dello sfondo possono essere scelti dall'utente.
  > La larghezza non supera gli 80 caratteri o glifi (40 se CJK).
  > Il testo non è giustificato (allineato sia al margine destro che al sinistro).
  > Lo spazio tra righe (interlinea) è almeno di uno spazio e mezzo all'interno del paragrafo e lo spazio tra paragrafi, è almeno una volta e mezzo l'interlinea.
  > Il testo può essere ridimensionato fino al 200 percento senza il supporto di tecnologie assistive in modo da non richiedere all'utente di dover scorrere orizzontalmente per leggere una riga di testo in una finestra a tutto schermo.

##### Violazioni riscontrate:

L'unico punto passato è 'il testo non è giustificato'

 - [ ] Il sito non permette di cambiare i colori in primo piano e dello sfondo
  
 - [ ] La larghezza del testo dipende dalla dimensione della finestra e non è possibile imporre un limite a 80 caratteri
  
 - [x] Il testo non è giustificato
  
 - [x] Lo spazio tra le righe è almeno di 1.5 e lo spazio tra paragrafi è almeno una volta e mezzo l'interlinea
  
 - [ ] Non è possibile ingrandire il testo e se si ingrandisce con lo Zoom con il browser al 200% compaiono delle barre di scorrimento orizzontali

#### Immagini di testo (senza eccezioni)
livello: AAA
outcome: no
source: images-of-text-no-exception

> Le immagini contenenti testo sono utilizzate soltanto per pura decorazione o quando una particolare presentazione del testo sia essenziale per il tipo di informazione veicolata.

##### Violazioni riscontrate:

Le immagini contenenti testo non sono utilizzate per pura decorazione e non sono essenziali per il tipo di informazione veicolata.

#### Ricalcolo del flusso
livello: AA
outcome: yes
source: reflow

> Il contenuto può essere ripresentato senza perdita di informazioni o funzionalità e senza richiedere lo scorrimento in due dimensioni per:

  > Contenuto a scorrimento verticale con una larghezza equivalente a 320 CSS pixel;

  > Contenuto a scorrimento orizzontale ad un'altezza equivalente a 256 CSS pixel.

> Tranne per le parti del contenuto che richiedono layout bidimensionale per l'utilizzo o per comprenderne il senso.

##### Osservazioni

 - [x] Il contenuto a scorrimento verticale con una larghezza equivalente a 320 CSS pixel non richiede di scorrere in due dimensioni
  
 - [x] Non c'è contenuto a scorrimento orizzontale

Ingrandendo a 400% il sito è usabile e rispetta le richieste, ma intorno al 200% bisogna scrollare orizzontalmente ed i testi importanti come titoli, sottotitoli sono "tagliati fuori".

#### Contrasto in contenuti non testuali
livello: AA
outcome: yes
source: non-text-contrast

> Nella presentazione visiva il rapporto di contrasto è di almeno 3:1 rispetto al colore o ai colori adiacenti per:

  > **Componenti dell'interfaccia utente**: Le informazioni visive richieste per identificare i componenti dell'interfaccia utente e gli stati (ad eccezione dei componenti inattivi o dove l'aspetto del componente è determinato dal programma utente e non modificato dall'autore);

  > **Oggetti grafici**: Parti di grafica necessarie per comprendere il contenuto, tranne quando una particolare presentazione di grafica sia essenziale per le informazioni trasmesse. 

##### Osservazioni

Tutte le immagini testuali hanno un livello di contrasto adeguato, oppure sono loghi.

#### Spaziatura del testo
livello: AA
outcome: no
source: text-spacing

> Nei contenuti implementati utilizzando linguaggi di markup che supportano le seguenti proprietà di stile per il testo, non si verifica alcuna perdita di contenuto o funzionalità impostando quanto segue senza modificare altre proprietà di stile:

  > Altezza della linea (interlinea) di almeno 1,5 volte la dimensione del carattere;
  > Spaziatura dopo i paragrafi di almeno 2 volte la dimensione del carattere;
  > Spaziatura tra le lettere di almeno 0,12 volte la dimensione del carattere;
  > Spaziatura tra le parole di almeno 0,16 volte la dimensione del carattere.

> Eccezione: le lingue e le scritture che non utilizzano una o più di queste proprietà nel testo scritto sono conformi quando si può applicare il criterio alle sole proprietà esistenti per quella combinazione di lingua e scrittura.

##### Violazioni riscontrate:

 - [x] L'altezza della linea (interlinea) è di 1,5 volte la dimensione del carattere
  
 - [x] La spaziatura dopo i paragrafi è almeno 2 volte la dimensione del carattere
  
 - [ ] Spaziatura tra le lettere di almeno 0,12 volte la dimensione del carattere
  
 - [ ] Spaziatura tra le parole di almeno 0,16 volte la dimensione del carattere
   
#### Contenuto con Hover o Focus
livello: AA
outcome: yes
source: content-on-hover-or-focus

> Nel caso in cui il passaggio del puntatore del mouse (hover) o il focus della tastiera rendono visibili e nascosti dei contenuti, sono soddisfatte le seguenti condizioni:

  > **Congedabile**: È disponibile un meccanismo per eliminare il contenuto aggiuntivo senza spostare il puntatore del mouse o il focus della tastiera, a meno che il contenuto aggiuntivo non comunichi un errore di immissione dei dati o non oscuri o sostituisca altri contenuti;

  > **Passabile**: Se il passaggio del puntatore del mouse sul contenuto può attivare il contenuto aggiuntivo, il puntatore può essere spostato sul contenuto aggiuntivo senza che questo scompaia;

  > **Persistente**: Il contenuto aggiuntivo rimane visibile fino a quando l'evento hover o focus non viene rimosso, l'utente lo elimina o le sue informazioni non sono più valide.

> Eccezione: la presentazione visiva del contenuto aggiuntivo è controllata dal programma utente e non viene modificata dall'autore.

##### Osservazioni

 - [x] Congedabile: tutte le aree a scomparsa oscurano parte dei contenuti
  
 - [x] Passibile: il contenuto aggiuntivo, una volta comparso per via di hover, rimane visibile se il puntatore (o il focus da tastiera) si spostano all'interno dell'area comparsa
  
 - [x] Persistente: il contenuto aggiuntivo rimane visibile fino a quando l'evento Hover o Focus non viene rimosso

## Utilizzabile

### Accessibile da tastiera

Rendere disponibili tutte le funzionalità tramite tastiera

#### Tastiera
livello: A
source: keyboard
outcome: no

Le voci di menu a scomparsa nell'intestazione (ad esempio 'FQ IN EDICOLA') non sono navigabili da tastiera in quanto l'area a scomparsa non si apre.

#### Nessun impedimento all'uso della tastiera
livello: A
source: no-keyboard-trap
outcome: yes

Il focus da tastiera può essere spostato da/verso qualunque componente che supporta il focus della pagina.

#### Tastiera (nessuna eccezione)
livello: AAA
source: keyboard-no-exception
outcome: no

Non lo è per lo stesso motivo del due sopra. (vedi sopra x2).

#### Tasti di scelta rapida
livello: A
source: character-key-shortcuts
outcome: yes

Non hanno implementato nessuna scorciatoia da tastiera quindi questo punto è banalmente rispettato.

### Adeguata disponibilità di tempo

Fornire agli utenti tempo sufficiente per leggere e utilizzare i contenuti.

#### Regolazione tempi di esecuzione
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

#### Pauso, stop, nascondi
livello: A
source: pasue-stop-hide
outcome: no

 - [x] i video che compaiono in sovraimpressione possono essere stoppati e/o nascosti
 - [ ] non è possibile disattivare l'autoaggiornamento

#### Nessun tempo di esecuzione
livello: AAA
source: no-timing
outcome: no

Nonostante il refresh non sia una parte esseziale del contenuto, la home
page si auto-ricarica ogni 10 minuti.


#### Interruzioni
livello: AAA
source: interruptions
outcome: no

La home page viene ricaricata ogni 10 minuti: tale azione non è
considerabile un'emergenza.

#### Riautenticazione
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

### Convulsioni e reazioni fisiche

Non sviluppare contenuti con tecniche che sia noto causino attacchi
epilettici o reazioni fisiche

#### Tre lampeggiamenti o inferiore alla soglia
livello: A
source: three-flashes-or-below-threshold
outcome: yes

Le pagine esaminate non contengono nulla che lampeggi per più di
tre volte in un secondo oppure il lampeggiamento è al di sotto della
soglia. Va sottolineato però che i video che compaiono a lato cambiano
spesso: durante la nostra ispezione i video non contenevano elementi che
potessero invalidare questo punto, ma nulla vieta che in futuro ci siano.

#### Tre lampeggiamenti
livello: AAA
source: three-flashes
outcome: yes

La pagina web non contiene nulla che lampeggi per più di tre volte in
un secondo. Si faccia riferimento comunque al punto precedente per il
discorso dei video.


#### Animazione da interazioni
livello: AAA
source: animation-from-interaction
outcome: no

Per menu a scomparsa delle sezioni e i video a scomparsa non è possibile
disabilitare l'animazione (transazione) e non è essenziale.

### Navigabile

Fornire delle funizonalità di supporto all'utente per navigar, trovare
contenuti e determinare la propria posizione.

#### Salto di blocchi
livello: A
source: bypass-blocks
outcome: no

Non è presente nessun meccanismo per saltare i blocchi di contenuto
che si ripetono su piu pagine.

#### Titolazione della pagina
livello: A
source: page-titled
outcome: yes

Le pagine web hanno titoli che ne descrivono l'argomento.

#### Ordine del focus
livello: A
source: focus-order
outcome: yes

Rispettato banalmente: la pagine viene navigata in modo sequenziale
usando il tab ma l'ordine di navigazione non influisce sul significato.

#### Scopo del collegamento
livello: A
source: link-purpose-in-context
outcome: no

Mancano dei aria-label sui box del meteo e della vignetta in homepage. Le
altre pagine sono a posto.

#### Differenti modalità
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

#### Intestazioni ed etichette
livello: AA
source: headings-and-labels
outcome: yes

Le sezioni esaminate hanno intestazioni che ne specificano il contenuto:
 - nella pagina di un articolo l'h1 descrive l'articolo, e gli h2 introducono gli articoli correlati
 - nella pagina del meteo gli h2 introducono i riquadri a lato
 - nella pagina di login l'h3 introduce lo scopo della pagina

TODO: nella pagina di login ci dovrebbe essere un h1 però...

#### Focus visibile
livello: AA
source: focus-visible
outcome: yes

Le parti dell'interfaccia vengono evidenziate dallo user agent quando
ricevono focus.

#### Posizione
livello: AAA
source: location
outcome: yes

Nelle pagine di archivio e degli articoli sono presenti delle
*breadcrumb*. Nella pagina del meteo la posizione corrente viene mostrata
nella nav.

#### Scopo del collegamento (solo collegamento)
livello: AAA
source: link-purpose-link-only
outcome: no

Nella homepage il box del meteo è un'immagine senza testo descrittivo,
così come anche il box della vignetta.

#### Intestazioni di sezione
livello: AAA
source: section-headings
outcome: yes

I titoli vengono usati correttamente (ma si veda comunque il punto
[Informazioni e correlazioni](#informazioni-e-correlazioni) per gli
errori semantici)

### Modalità di input

Rendere più facile agli utenti l'utilizzo di funzionalità attraverso
input diversi dalla tastiera.

#### Movimenti del puntatore
livello: A
source: pointer-gestures
outcome: yes

Non sono richiesti gisti multi punto o basati su percorsi.

#### Cancellazione delle azioni del puntatore
livello: A
source: pointer-cancellation
outcome: yes

non sono presenti down-event nelle pagine esaminate.

#### Etichetta nel nome
livello: A
source: label-in-name
outcome: no

 - il controllo di ricerca ha il nome ha nome 'q'
 - nella schermata di login i controlli hanno nome 'username' ma il testo mostrato è 'Nome utente o email'

#### Azionamento da movimento
livello: A
source: motion-actuation
outcome: yes

Non sono presenti funzionalità che richiedano di essere attivate dal
movimento del dispositivo o dell'utente.

#### Dimensione dell'obbiettivo
livello: AAA
source: target-size
outcome: no

Il link vuoto nella barra laterale non è alto abbastanza (0px)

#### Meccanismi di input simultanei
livello: AAA
source: concurrent-input-mechanisms
outcome: yes

L'unica modalità di input richiesta è il click e l'hover.

## Comprensibile

> Le informazioni e le operazioni dell'interfaccia utente devono essere
> comprensibili.

### Leggibile

> Rendere il testo leggibile e comprensibile.

#### Lingua della pagina
livello: A
source: language-of-page
outcome: no

l'unica violazione è la pagina del meteo che non ha un attributo lang
sull'elemento html, non ha un meta Content-Language e, quando viene
servita, non è presente  un header Content-Language.

#### Parti in lingua
livello: AA
source: language-of-parts
outcome: yes

Negli articoli esaminati, il testo è sempre in italiano fatta eccezione
per i nomi propri, alcune terminologie tecniche od alcune espressioni
diventate parte integrante del parlato.

#### Parole inusuali
livello: AAA
source: unusual-words
outcome: no

Non è presente nessun meccanismo per identificare parole o frasi usate
in modo insolito o ristretto.

#### Abbreviazioni
livello: AAA
source: abbreviations
outcome: no

Non è presente nessun meccanismo per identificare la forma espansa o
il significato delle abbreviazioni.

#### Livello di lettura
livello: AAA
source: reading-level
outcome: yes

Negli articoli esaminati il testo non ci sembra richieda capacità
di lettura più avanzata rispetto al livello di istruzione secondaria
inferiore.

#### Pronuncia
livello: AAA
source: pronunciation
outcome: yes

Si perchè il contenunto prevalente è in italiano e non abbiamo i *false friends*.

### Prevedibile

Creare pagine web che abbiano aspetto e funzionalità prevedibili.

#### Al Focus
livello: A
source: on-focus
outcome: yes

Nelle pagine esaminate quando un componente dell'interfaccia utente
riceve il focus non si avvia nessun cambiamento del contesto.

#### All'input
livello: A
source: on-input
outcome: yes


TODO: tenere a mente quando si progetta il campo di ricerca

#### Navigazione Coerente
livello: AA
source: consistent-navigation
outcome: yes

Gli unici meccanismi di navigazioni che sono ripetuti su più pagine
web sono i seguenti e appaiono sempre nello stesso ordine relativo:
 - menu a scomparsa laterale
 - link del footer

#### Identificazione coerente
livello: AA
source: consistent-identification
outcome: yes

I componenti che hanno la stessa funzionalità (come i controlli 'Articolo
successivo/precedente', il menu a scomparsa, il footer) sono identificati
consistentemente.

#### Cambiamenti su richiesta
livello: AAA
source: change-on-demand
outcome: no

La home page ha una funzione di auto-refresh attiva che periodicamente
aggiorna la pagina. Questo provoca un cambiamento di contesto (la
posizione relativa all'interno della pagina viene persa, così come
anche il focus).

### Assistenza nell'inserimento

Aiutare gli utenti a evitare gli errori e agevolarli nella loro correzione.

#### Identificazione di errori
livello: A
source: error-identification
outcome: no

Provare a postare un commento (vuoto) da loggati non risulta in nessuna
forma di feedback che permetta all'utente di capire che l'azione non è
permessa. Non viene identificato nessun elemento e non compare nessun
testo.

Nella pagina di login invece, se vengono fornite credenziali errate,
viene fornito un messaggio di errore.

#### Etichette o istruzioni
livello: A
source: labels-or-instructions
outcome: no

Nonostante nella schermata di login ci siano appropriati etichette, la
textarea per i commenti non ha etichette adeguate (è presente una scritta
in sovraimpressione ma viene aggiunta con javascript e posizionata lì
sopra, non ha nessuna vicinanza all'elemento che permetta di comprenderne
l'associazione se non il risultato grafico).

#### Seggerimenti per gli errori
livello: AA
source: error-suggestion
outcome: no

Si veda ad esempio la già citata textarea per i commenti.

#### Prevenzione degli errori
livello: AA
source: error-prevention-legal-financial-data
outcome: yes

Nelle pagine controllate non sono presenti pagine con vincoli di tipo
giuridico, transazioni finanziarie o la modifica/cancellazione o gestione
di dati controllabili dall'utente.

#### Aiuto
livello: AAA
source: help
outcome: yes

 - è presente una pagina (seppur limitata) di aiuto (linkata sempre nel footer)
 - ci sono degli aiuti contestuali

## Robusto

Il contenuto deve essere abbastanza robusto per essere interpretato in
maniera affidabile da una grande varietà di programmi utente, comprese
le tecnologie assitive.

### Compatibile

Garantire la massima compatibilità con i programmi utente attuali e
futuri, comprese le tecnologie assistive.

#### Analisi sintattica (parsing)
livello: A
source: parsing
outcome: no

copiare dal validatore:
 - apertura e chiusura corretta dei tag
 - attributi non duplicati
 - id univoci

sappiamo per certo che i tag non sono corretti (chiusi due volti i p
per esempio).

#### Nome, ruolo, valore
livello: A
source: name-role-value
outcome: no

Non è possibile calcolare name e ruolo per ogni componente
dell'interfaccia utente (verificato con gli strumenti di accessibilità
di firefox: la barra di ricerca per esempio non viene riconosciuta).

#### Messaggi di stato
livello: AA
source: status-messages@
outcome: no

Non sono presenti messaggi di stato che possano essere determinati
programmaticamente.

Infatti, abbiamo visto come per moltissime componenti della pagina non
sia possibile stabilire programmaticamente il ruolo.

## Conformità
