<!DOCTYPE html>
<html>
   <head>
      <title>HTML img Tag</title>
   </head>

   <div class="images">
   <body>
      <h2>Teilaufgabe 1</h2>
      <br>
      <br>
      <b>Customer</b>
      <br>
      <p><i>Daten: </i>customerID, name, address</p>
      <br>
      <i>Funktionalität:</i>
      <ul type="circle">
         <li>neuen Kunden anlegen</li>
         <li>Kunden finden</li>
         <li>Kunden löschen</li>
      </ul>
      <br>
      <i>Kommunikation mit:</i>
      <ul type="circle">
         <li>Client (synchron)</li>
         <li>Shipment(synchron)</li>
      </ul>
      <br>
      <b>Catalog</b>
      <br>
      <p><i>Daten: </i>articleID, name, price</p>
      <br>
      <i>Funktionalität:</i>
      <ul type="circle">
         <li>Artikel hinzufügen</li>
         <li>einzelnen Artikel zurückgeben</li>
         <li>alle Artikel auflisten</li>
         <li>alle Artikel anzeigen, die im Bestand sind</li>
         <li>Artikel löschen</li>
      </ul>
      <br>
      <i>Kommunikation mit:</i>
      <ul type="circle">
         <li>Client (synchron)</li>
         <li>Stock (synchron)</li>
         <li>Order (synchron)</li>
      </ul>
      <br>
      <b>Stock</b>
      <br>
      <p><i>Daten: </i>articleID, amount</p>
      <br>
      <i>Funktionalität:</i>
      <ul type="circle">
         <li>Artikel mit Bestand hinzufügen</li>
         <li>Bestand eines einzelnen Artikels zurückgeben</li>
         <li>Bestand eines Artikels erhöhen/reduzieren</li>
         <li>alle Artikel und Bestände auflisten</li>
      </ul>
      <br>
      <i>Kommunikation mit:</i>
      <ul type="circle">
         <li>Client (synchron)</li>
         <li>Catalog (synchron)</li>
         <li>Order (synchron)</li>
      </ul>
      <br>
      <b>Order</b>
      <br>
      <p><i>Daten: </i>orderID, customerID, Liste von articleID und dazugehörigen Mengen, isPayed (bool), isSent(bool)</p>
      <br>
      <i>Funktionalität:</i>
      <ul type="circle">
         <li>Bestellung aufgeben</li>
         <li>defekte Artikel zurückschicken</li>
         <li>Bestellung stornieren (falls noch nicht versandt)</li>
      </ul>
      <br>
      <i>Kommunikation mit:</i>
      <ul type="circle">
         <li>Client (synchron)</li>
         <li>Catalog (synchron)</li>
         <li>Stock (synchron)</li>
         <li>Payment (asynchron, Empfänger)</li>
         <li>Shipment (asynchron, Empfänger)</li>
      </ul>
      <br>
      <b>Payment</b>
      <br>
      <p><i>Daten: </i>-</p>
      <br>
      <i>Funktionalität:</i>
      <ul type="circle">
         <li>Bestellung bezahlen</li>
      </ul>
      <br>
      <i>Kommunikation mit:</i>
      <ul type="circle">
         <li>Client (synchron)</li>
         <li>Shipment (asynchron, Sender)</li>
         <li>Order (asynchron, Sender)</li>
      </ul>
      <br>
      <b>Shipment</b>
      <br>
      <p><i>Daten: </i>-</p>
      <br>
      <i>Funktionalität:</i>
      <ul type="circle">
         <li>Bestellung versenden</li>
      </ul>
      <br>
      <i>Kommunikation mit:</i>
      <ul type="circle">
         <li>Client (synchron)</li>
         <li>Order (synchron)</li>
         <li>Customer (synchron)</li>
         <li>Order (asynchron, Sender)</li>
         <li>Payment (asynchron, Empfänger)</li>
      </ul>
      <br>
      <b>Ablauf einer Bestellung eines nicht registrierten Kunden:</b>
      <br>
      <b>Der Client fragt zunächst alle Artikel ab, die auf Lager sind. Dazu wird eine Nachricht an den Catalog-Service geschickt. Da der Catalog-Service nicht weiß, welche Artikel auf Lager sind, muss er eine Nachricht an den Stock-Service schicken. Dieser Antwortet mit einer Auflistung aller Artikel, die auf Lager sind. Der Catalog-Service schickt die entsprechenden Artikel dann an den Client.</b>
      <br>
      <b>Der Client entscheidet sich dann einen der angezeigten Artikel zu bestellen und schickt eine Nachricht mit der articleID an den Order-Service. Da der Kunde noch nicht registriert wurde, wird auch keine customerID mitgeschickt. Also Antwortet der Order-Service zunächst mit dem Hinweis, sich beim Customer-Service zu registrieren.</b>
      <br>
      <b>Der Client schickt dann eine Nachricht an den Customer-Service mit seinem Namen und seiner Adresse. Dieser dann einen neuen Kunden mit dem entsprechenden Namen, der Adresse und einer neuen customerID an. Letztere wird in der Antwort an den Kunden mitgegeben.</b>
      <br>
      <b>Nach der Registrierung schickt der Client jetzt erneut seine Bestellung an den Order-Service und gibt dieses mal seine customerID mit. Der Order-Service versichert sich, dass es wirklich einen Kunden mir der übergebenen ID gibt. Danach schickt der Order-Service eine Nachricht an den Stock-Service mit der articleID und der Anzahl, um die der Bestand dieses Artikels reduziert werden soll.</b>
      <br>
      <b>Der Stock-Service reduziert dann den Bestand des entsprechenden Artikels.</b>
      <br>
      <b>Der Order-Service berechnet dann den Gesamtpreis der Bestellung und schickt diesen an den Client zurück mit der Aufforderung ihn beim Payment-Service zu bezahlen.</b>
      <br>
      <b>Der Client schickt dann eine Nachricht mit der orderID an den Payment-Service. Dies entspricht der Bezahlung. Der Payment-Service publisht dann, dass die Bestellung mit der orderID bezahlt ist.</b>
      <br>
      <b>Daraufhin setzt der Order-Service isPayed bei der entsprechenden Bestellung auf true.</b>
      <br>
      <b>Der Shipment-Service reagiert daraufhin ebenfalls. Er fragt die Artikel und die dazugehörigen Mengen vom Order-Service mit der orderID ab. Außerdem holt er sich die customerID der Bestellung und fragt damit den Namen und die Adresse des Kunden beim Customer-Service ab. Anschließend schickt er eine Versandbestätigung an den Kunden und publisht die orderID.</b>
      <br>
      Daraufhin setzt der Order-Service isSent auf true.</b>
      <br>
      <img src="img/TA1.png" alt="first image"/>
      <br>
      <p><small>Bild 1. Vor dem Registrieren</small></p>
      <br>
      <img src="img/TA2.png" alt="second image"/>
      <br>
      <p><small></small>Bild 2. Nach dem Registrieren</small></p>
   </body>
</div>
</html>