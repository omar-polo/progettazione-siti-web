title: Linee guida per l'accessibilità dei contenuti Web (WCAG) 2.1

# Linee guida per l'accessibilità dei contenuti Web (WCAG) 2.1

## Introduzione

La valutazione di conformità è stata effettuata nei riguardi delle linee guida [**WCAG 2.1**](https://www.w3.org/Translations/WCAG21-it/). Tali linee guida prevedono diversi livelli di orientamento che comprendono *principi* globali, *linee guida* generali, *criteri di successo* verificabili e una ricca raccolta di tecniche *sufficienti* e *consigliate*.

Principi
: Al livello più alto, sono definiti i quattro principi fondamentali dell'accessibilità sul Web: *percepibile*, *utilizzabile*, *comprensibile* e *robusto*. 

Linee guida
: Dai quattro principi derivano tredici linee guida. Le linee guida forniscono gli obiettivi di base su cui gli autori dovrebbero lavorare per rendere i contenuti accessibili agli utenti. 

Criteri di successo
: Per ogni linea guida vengono forniti criteri di successo verificabili per consentire l'utilizzo delle WCAG 2.1 per test dei requisiti e della conformità. Vengono definiti tre livelli di conformità: A (minimo), AA e AAA (massimo).

Tecniche sufficienti e consigliate
: Le tecniche sono informative e possono essere di due categorie: quelle sufficienti a soddisfare il criterio di successo e quelle consigliate. Le tecniche consigliate vanno oltre ciò che viene richiesto da ciascun singolo criterio di successo e consentono di rispettare meglio le linee guida. 

Tutti questi livelli di orientamento (principi, linee guida, criteri di successo, tecniche sufficienti e consigliate) concorrono a fornire indicazioni su come rendere il contenuto più accessibile. 

Sono stati utilizzati diversi strumenti come supporto alla valutazione di conformità alle linee guida WCAG 2.1.

- I browser Web [Mozilla Firefox Developer Edition](https://www.mozilla.org/it/firefox/channel/desktop/) e [Google Chrome](https://www.google.com/intl/it/chrome/)

- Markup Validation Service [W3C](https://validator.w3.org/)

- CSS Validation Service [W3C](https://jigsaw.w3.org/css-validator/)

- Web Accessibility Evaluation Tool [WAVE](http://wave.webaim.org/)

- Color Contrast Accessibility Validator [a11y](https://color.a11y.com/?wc3)

- [Color Contrast Checker - WebAIM](https://webaim.org/resources/contrastchecker/) che fornisce un rapporto per sapere quanto è accessibile un sito Web e come migliorarlo.

- [Accessibility Inspector](https://developer.mozilla.org/it/docs/Tools/Accessibility_inspector?utm_source=devtools&utm_medium=a11y-panel-description ) che fornisce un mezzo per accedere alle informazioni importanti esposte alle tecnologie assistive nella pagina corrente tramite l'albero di accessibilità.

- [Lighthouse](https://developers.google.com/web/tools/lighthouse/) che è uno strumento automatizzato open-source per migliorare la qualità delle pagine web. Ha Audits per prestazioni, accessibilità, buone pratiche, applicazioni web progressive e altro.

### Perimetro dell'indagine

Per la valutazione dei criteri di conformità alle linee guida WCAG 2.1 sono state selezionate le seguenti pagine:

- Home page del sito [il Fatto Quotidiano](https://www.ilfattoquotidiano.it/)

- Un articolo estratto dalla Home page

- La pagina [Archivi](https://www.ilfattoquotidiano.it/archivi/)

- La pagina di [Login](https://shop.ilfattoquotidiano.it/login/)

- La pagina delle [Previsioni Meteo](http://meteo.ilfattoquotidiano.it/)

### Risultati del report

Il sito analizzato non presenta nessuna alternativa per i contenuti non testuali e ciò rende difficile la fruibilità dei contenuti agli utenti con particolari necessità. 

I video preregistrati non forniscono nessuna trascrizione testuale equivalente, nè presentano sottotitoli.

Le informazioni trasmesse determinate programmaticamente non rispettano la semantica del Web. Il vantaggio della scrittura di HTML semantico deriva da quello che dovrebbe essere l'obiettivo principale di qualsiasi pagina web: il desiderio di comunicare.

Il sito fa largo uso di immagini di testo che non sono essenziali per le informazioni che intendono veicolare, quando sarebbe stato più opportuno utilizzare puro testo.

Dal punto di vista dell'utilizzabilità, nelle voci di menu a scomparsa non è possibile accedere tramite la tastiera.

Il sito, quindi, risulta **NON CONFORME** a nessun livello di conformità previsto dalle linee guida WCAG 2.1.

### Requisiti di conformità
Nella valutazione di conformità dei vari criteri di successo l'esito è stato espresso nei termini di:

**Positivo**
: Tutte le pagine esaminate rispettano completamente le direttive espresse da quel criterio di successo

**Negativo**
: Una o più pagine esaminate non rispettano il criterio di successo

**Non applicabile**
: Per tutte le pagine esaminate il criterio di successo non è applicabile

### Glossario

[↓ Salta glossario](#sommario-dei-risultati)

spec
: Abbreviazione di "specifica"

WCAG
: Web Content Accessibility Guidelines, attualmente alla versione 2.1

CHAPTCHA
: Sigla per "Completely Automated Public Turing test to tell Computers and Humans Apart" (Test di Turing pubblico e completamente automatizzato, allo scopo di distinguere gli esseri umani dai computer)

link
: Collegamento tra due pagine Web

input
: Un elemento HTML che permette all'utente di inserire un contenuto (sia esso testo, file, preferenza o altro)

HTML
: Hyper Text Markup Language, linguaggio di markup con il quale sono realizzate le pagine Web

CSS
: Cascading Style Sheet, linguaggio usato per definire la resa grafica degli elementi di una pagina Web

CSS pixel
: Un CSS pixel è l'unità di misura canonica per tutte le lunghezze e misure in CSS. Questa unità è indipendente dalla densità e distinta dai pixel hardware effettivi presenti in un display. 

Pagina Web
: Risorsa non incorporata ottenuta da un unico URI utilizzando HTTP più qualunque altra risorsa utilizzata per il rendering

Browser Web
: Qualsiasi programma che recuperi e presenti contenuti Web agli utenti

Header HTTP
: Un'intestazione HTTP, come prevista dal protocollo. Sono piccole informazioni che precedono il contenuto della pagina web

ARIA
: *Accessible Rich Internet Applications*. ARIA definisce le modalità con cui sviluppatori di browser, media player, dispositivi mobili e tecnologie assistive, cosí come gli sviluppatori di contenuti, possono migliorare l'accessibilità cross-platform.

Braille
: Un sistema di lettura e scrittura tattile a rilievo per non vedenti e ipovedenti.

Breadcrumb
: letteralmente "Briciole di pane". TODO

### Sommario dei risultati

[↓ Salta sommario risultati](#toc)
