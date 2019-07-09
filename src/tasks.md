Questo testo verrà ignorato.

I vari casi d'uso sono separati da '###'. Ogni caso d'uso deve avere le
seguenti informazioni:

 - storia : la descrizione del caso d'uso
 - tempo: il tempo massimo previsto
 - percorso: il percorso previsto

La grammatica è la seguente:
```
	UC: head '\n'* data+
	head: '###'
	data: key ':' '\n' definition '\n' '\n'
```
dove `key` è una keyword (descritte sopra) e definition una sequenza
di linee di testo.

###

storia:
Stai progettando un viaggio (e.g. Udine - Verona) e vuoi sapere quanto
ti costerà il pedaggio.

tempo:
2 minuti

percorso:
 - home page
 - click sul box "calcolo pedaggi"
 - inserire le informazioni e premere il tasto "calcola"

###

storia:
Vuoi verificare il meteo per una certa città (e.g. Venezia) per domenica
prossima.

tempo:
3 minuti

percorso:
 - home page
 - servizi
 - previsioni meteo

###

storia:
Vuoi verificare il traffico su una autostrada per una specifica data.

tempo:
3 minuti

percorso:
 - home page
 - uso del widget per il traffico

###

storia:
Ricerca di una specifica telecamera

tempo: 
1 minuto

percorso:
 - home page
 - telecamere
 - uso della mappa

###

storia:
Viene richiesto all'utente di trovare la stazione di benzina più vicina
ad una città (e.g. Udine venendo da Venezia.)

tempo:
2 minuti

percorso:
 - home page
 - infrastrutture
 - uso della mappa

###

storia:
Viene richiesto all'utente di trovare la stazione di servizio con benzina
con costo minore su una tratta di un'autostrada.

tempo:
2 minuti

percorso:
 - home page
 - infrastrutture
 - ricerca manuale nella lista

