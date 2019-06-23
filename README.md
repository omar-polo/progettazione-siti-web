# Progettazione siti web

## struttura del progetto

 - in src ci sono tutti le pagine, ognuna scritta in markdown (con alcune
   estensioni alla sintassi per quanto riguarda la pagina del wcag.)
   
 - in build vengono generati i file;
 
 - le cartelle css e js vengono copiate nella cartella build
 
 - template.html è il template con il quale vengono renderizzate le pagine.
 
#### compilare il programma d'aiuto

> richiede go

    go build

#### compilare il report
    
    ./progettazione-siti-web build

#### servire e (contemporaneamente) compilare quando i file cambiano

    ./progettazione-siti-web dev

#### servire il report

    ./progettazione-siti-web serve

la porta di default è 8000


## TODOs

 - [X] generare i titoli per le pagine
 - [ ] verificare gli aria landmark nelle nostre pagine