# Progettazione siti web

## struttura del progetto

 - in src ci sono tutti le pagine, ognuna scritta in markdown
 
   **IMPORTANTE**: non mettere spazi nei nomi dei file;
   
 - in build vengono generati i file;
 
 - build.sh è uno script che compila il sito (si legga più avanti per info);
 
 - style.css è il foglio di stile;
 
 - template.html è il template con il quale vengono renderizzate le pagine.

#### compilare il report

    sh build.sh

#### servire il report

> richiede python

    sh build.sh serve [ port ]

dove port di default è 8000
