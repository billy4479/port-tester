# Port Tester

Un piccolo tool per vedere se le porte sono state aperte correttamente.

## Utilizzo

Scarica da [qui](https://github.com/billy4479/port-tester/releases) l'ultima versione per il tuo sistema operativo.

Su windows apri la cartella in cui hai scaricato `port-tester_windows_x86-64.exe` e apri `cmd` in quella cartella e digita il seguente comando:

```sh
./port-tester.exe -port <porta da testare>
```

Sostituisci `<porta da testare>` con la vera porta (per minecraft è `25565`). Assicurati di averla anche aperta nel modem se no ovviamente non funzionerà.

## Compila

Nel caso tu sia un pazzo furioso e lo voglia compilare dal codice sorgente usa:

```sh
git clone https://github.com/billy4479/port-tester
cd port-tester
make # Compila per il tuo sistema
# make windows # Cross-compila per windows
# Puoi anche usare semplicemente
# go build
```
Dipendenze: 
- go (make)
- GNU Make (make)
