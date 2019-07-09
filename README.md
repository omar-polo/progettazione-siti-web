# Progettazione siti web

## struttura del progetto

 - in `src` ci sono tutti i dati. La maggior parte sono scritte in markdown, altre con formati specifici.
   
 - in build vengono generati i file;
 
 - le cartelle `css`, `js` e `img` vengono copiate nella cartella build
 
 - template.gohtml è il template con il quale vengono renderizzate le pagine.
 
#### compilare il programma d'aiuto

> richiede go

    go build

dovrebbe produrre un eseguibile dal nome `progettazione-siti-web`
(prende il nome dalla cartella che lo contiene). Su windows tale file
avrà estensione `.exe`.

#### compilare il report
    
    ./progettazione-siti-web build

#### servire e (contemporaneamente) compilare quando i file cambiano

    ./progettazione-siti-web dev

#### servire il report

    ./progettazione-siti-web serve

la porta di default è 8000
