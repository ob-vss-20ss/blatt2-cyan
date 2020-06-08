# Teilaufgabe 1


### Customer

*Daten:* customerID, name, address

*Funktionalität:*
* neuen Kunden anlegen
* Kunden finden
* Kunden löschen

*Kommunikation mit:* 
* Client (synchron)
* Shipment(synchron)


### Catalog

*Daten:* articleID, name, price 

*Funktionalität:* 
* Artikel hinzufügen
* einzelnen Artikel zurückgeben
* alle Artikel auflisten
* alle Artikel anzeigen, die im Bestand sind
* Artikel löschen

*Kommunikation mit:* 
* Client (synchron)
* Stock (synchron)
* Order (synchron)


### Stock

*Daten:* articleID, amount 

*Funktionalität:* 
* Artikel mit Bestand hinzufügen
* Bestand eines einzelnen Artikels zurückgeben
* Bestand eines Artikels erhöhen/reduzieren
* alle Artikel und Bestände auflisten

*Kommunikation mit:* 
* Client (synchron)
* Catalog (synchron)
* Order (synchron)


### Order 

*Daten:* orderID, customerID, Liste von articleID und dazugehörigen Mengen, isPayed (bool), isSent(bool)

*Funktionalität:* 
* Bestellung aufgeben
* defekte Artikel zurückschicken
* Bestellung stornieren (falls noch nicht versandt)

*Kommunikation mit:* 
* Client (synchron)
* Catalog (synchron)
* Stock (synchron)
* Payment (asynchron, Empfänger)
* Shipment (asynchron, Empänger)


### Payment

*Daten:* -

*Funktionalität:* 
* Bestellung bezahlen

*Kommunikation mit:* 
* Client (synchron)
* Shipment (asynchron, Sender)
* Order (asynchron, Sender)


### Shipment

*Daten:* -

*Funktionalität:* 
* Bestellung versenden

*Kommunikation mit:* 
* Client (synchron)
* Order (synchron)
* Customer (synchron)
* Order (asynchron, Sender)
* Payment (asynchron, Empfänger)


![image info](./img/TA1.png)

![image info](./img/TA2.png)


