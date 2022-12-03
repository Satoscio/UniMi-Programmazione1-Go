# Laboratorio di programmazione - Istruzioni per l'esame

NOTA BENE: le presenti istruzioni e tutti gli script forniti si basano su sistema operativo GNU/Linux (Debian in particolare) e prevedono l'abilità dello studente nell'uso della shell (bash in particolare).
Chi fosse abituato a usare sistemi operativi diversi dovrà adattarsi usando proprie conoscenze e abilità.

## Suggerimenti

Si consiglia una attenta lettura di questo file prima dell'esame.

Si consiglia di leggere attentamente tutti i testi degli esercizi prima di iniziare il lavoro e di usare gli esempi di esecuzione per verificare di aver compreso correttamente le specifiche. 

Si consiglia di fare una valutazione della complessità degli esercizi e di farsi quindi un ordine di lavoro.

Si invita a chiedere chiarimenti in caso di dubbi sulle specifiche.

## Scompattare i file

Lanciando il comando 'tar xvf <nome del file>.tar' viene espanso l'archivio a partire dalla directory in cui ci si trova (e si trova il file stesso). Verrà create una sottodirectory per ogni esercizio con tutto il materiale occorrente per l'esercizio stesso (specifiche, test, ecc.)
Si invitano gli studenti a lavorare nelle sottodirectory così create.



## Contenuto della directory Tema-<data_esame>

Troverete una serie di sottodirectory con i testi degli esercizi, in due forme possibili:

A) file nella forma '<nome esercizio>_test.go' + eventuali file utili per i test.
In testa al file '<nome esercizio>_test.go' si trova, sotto forma di commento, il testo dell'esercizio. Di seguito si trovano i test da eseguire per fare una prima (auto)valutazione del proprio elaborato (vedi più avanti su come usare i test).


B) file nella forma 'testo.txt' + eventuali file di input + script di shell.
In questo caso il testo dell'esercizio si trova in 'testo.txt', mentre gli script sono:
- 'oracolo.sh' è una implementazione delle specifiche usando comandi Unix: è cioè possibile eseguire './oracolo.sh' con i medesimi parametri passati alla propria implementazione e confrontare i due output
- 'test.sh' esegue l'elaborato (che deve già essere compilato!) ed effettua un confronto tra l'output dell'oracolo e l'output della propria implementazione, per eseguirlo basta lanciare da prompt './test.sh'



## Uso dei test Go (**ove presenti**)

Per svolgere il compito bisogna entrare in ogni sottodirectory e creare il relativo file 'go' (se, ad esempio, il test si chiama 'filtro_test.go', va creato il file 'filtro.go'; attenzione alle maiuscole e minuscole, in generale sono caratteri tutti minuscoli) secondo le specifiche.

**NOTABENE** per chi usa la versione di Go dalla 1.16 in poi: per poter usare i test bisogna prima lanciare almeno una volta nella directory di ogni esercizio il comando 'go mod init <nome dell'esercizio senza estensione .go>', ad esempio:
$ go mod init esercizioUno

Per "testare" il codice basta compilare innanzitutto il proprio elaborato con 'go build', poi si può lanciare il testing con 'go test -v'. Ricordarsi di ricompilare dopo ogni modifica prima di lanciare di nuovo i test.

PASS indica che il programma x.go ha superato i test contenuti in x_test.go (ma tra questi potrebbero esserci test che fanno vedere output ottenuto e/o output atteso, senza fare nessun confronto né controllo, per i quali quindi PASS indica solo un superamento formale e non sui risultati)
FAIL indica che il programma x.go NON ha superato almeno uno dei test contenuti in x_test.go

In caso di FAIL si consiglia VIVAMENTE di esaminare l'output dei test per capire cosa va corretto nel proprio programma.

Si consiglia VIVAMENTE inoltre di esaminare anche il codice dei test oltre al testo dell'esercizio al fine di comprendere meglio ciò che viene chiesto per lo svolgimento.

Nota Bene: i test effettivi eseguiti dai docenti in fase di correzione potrebbero essere in numero e tipo diversi da quelli presentati nel tema d'esame. I.e., un PASS non significa automaticamente che l'esercizio sia privo di errori. Inoltre a determinare il voto concorreranno anche altri aspetti, quali la semplicità del codice, la leggibilità, l'uso della memoria, ecc.



## Consegna

Va effettuata caricando i SINGOLI file '<nome esercizio>.go' (quindi NON vanno consegnati gli eseguibili, i test e gli altri file) su https://upload.di.unimi.it utilizzando la propria login e scegliendo la SESSIONE CORRETTA (nel dubbio chiedere ai docenti).
Consegnare in una sessione sbagliata comporta la perdita della consegna e il conseguente annullamento dell'esame.

Si può caricare più volte lo stesso file, verrà valutata soltanto l'ultima versione depositata. Si consiglia di caricare il file anche 'in itinere' in modo da avere un backup in caso di problemi alla postazione.

Nota bene: la sessione (login) di upload scade dopo circa 15 minuti di inattività, quindi nel caso occorre semplicemente fornire di nuovo le proprie credenziali (evitare trovarsi in questa situazione a un minuto dall'ora di consegna).

***ATTENZIONE***: la sessione d'esame si CHIUDE AUTOMATICAMENTE all'orario "di consegna" che viene comunicato in aula dai docenti. Fino a quell'ora si possono consegnare i propri elaborati, oltre no e NON SI FARANNO ECCEZIONI. Al termine della consegna il sistema di upload NON ACCETTA più il caricamento dei file. Si consiglia di non ridursi all'ultimo momento.


## Per ritirarsi

Occorre caricare un file vuoto dal nome 'ritirato.txt' (entro l'orario di consegna).



## Chi ha terminato (sia per consegna che per ritiro)

E` pregato semplicemente di scollegarsi dalla sessione di upload, senza spegnere il PC.



## Valutazione

***ATTENZIONE***: condizione necessaria affinché tutto l'elaborato venga valutato è che sia presente ALMENO UN esercizio che compila e gira perfettamente (cioè produca l'output atteso).

In ogni caso verranno valutati (cioè ne verrà esaminato e giudicato il sorgente) solo gli esercizi che compilano senza errori, gli altri non concorreranno alla formazione del voto.
Gli esercizi i cui test non dovessero compilare verranno penalizzati d'ufficio.

Alla formazione del voto concorreranno anche i seguenti aspetti:
- che il programma mostri un funzionamento corretto su casi di test significativi
- la semplicità della soluzione e della sua implementazione
- l'uso di nomi significativi per variabili, tipi e funzioni (leggibilità)
- che nel programma sia evitata la duplicazione di codice
- la struttura del programma e l'uso di funzioni (modularità)
- l'uso della memoria, che deve essere limitato alla sola memoria necessaria.

Gli esercizi hanno pesi diversi, per cui non c'è un "numero minimo di esercizi giusti per passare l'esame". I pesi vengono assegnati a valle di una valutazione sul campo della difficoltà di svolgimento, quindi chiedere "quali sono i pesi?" durante l'esame non può ricevere risposta.



## Materiale utilizzabile

* documentazione online di Go (https://pkg.go.dev/std)
* libro di testo o manuale di Go (eventualmente pdf) (non eserciziari)
* penna
* carta (fornita dai docenti)
* finestra per l'upload
* editor
* terminale



## Materiale NON utilizzabile

Ogni altro oggetto non elencato nella sezione precedente. Quindi, ad esempio, NON si possono tenere in vista telefoni, appunti, altri libri, quaderni, etc.



## Comportamenti sanzionabili

Comunicare con chiunque (docenti esclusi) a esame iniziato comporta espulsione e annullamento dell'esame, per TUTTI i soggetti coinvolti nella comunicazione.
Idem in caso di utilizzo di materiale non consentito.
Se si ha bisogno di delucidazioni sul tema d'esame o per altre domande rivolgersi ai docenti (è sufficiente alzare la mano).
