# Teilaufgabe 1

Führen Sie das Folgende zunächst noch im `master`-Branch aus.

## Customer

![image info](./img/TA1.png)

Diesen Schritt muss nur einer aus dem Team durchführen.

1. Loggen Sie sich auf meinem [Drone-Server](https://terraform.cs.hm.edu/) ein.
2. Wählen Sie das Projekt aus und **aktivieren** Sie es.

Um vom Drone-Server die erzeugten Docker-Images auf GitHub pushen zu lassen, müssen Sie
auf GitHub einen _Personal access token_ erzeugen.

3. Gehen Sie dazu auf die Seite [Personal access token](https://github.com/settings/tokens)
   und klicken Sie auf `Generate new token`. Als `Scope` klicken Sie auf `write:packages`.
   Dabei wird `repo` und `read:packages` mit ausgewählt. Unter `Note` tragen Sie am
   Besten sowas wie `GitHub packages` ein, damit Sie später wissen, welchen Token
   Sie wieder löschen müssen, wenn die Veranstaltung zuende ist. Klicken Sie dann ganz unten
   auf `Generate token`.
4. Kopieren Sie den Token und gehen Sie zurück auf meinen Drone-Server in Ihr Projekt.
5. Wählen Sie dort `Settings` aus und scrollen Sie bis zu `Secrets`.
6. Tragen Sie `gh_token` bei `Secret Name` und Ihren Token bei `Secret Value` ein und
   klicken Sie auf `ADD A SECRET`.
7. Tragen Sie `gh_username` bei `Secret Name` und Ihren GitHub Username bei `Secret Value` ein und
   klicken Sie auf `ADD A SECRET`.
8. Für den Download der Docker-Images von GitHub ist noch ein weiteres Secret mit dem
   Namen `dockerconfig` notwendig:

    - Erzeugen Sie zunächst ein Base64-encodedes Authtoken, in dem Sie Ihren GitHub-Username
      und ihr unter 6 verwendetes Token mit einem `:` verbinden und damit folgendes machen:

        ```
        echo -u username:token | base64
        ```

        also z.B.

        ```
        echo -u obcode:123456789012345 | base64
        ```

    - Fügen Sie dieses zwischen die Anführungszeichen in folgende Zeichenkette ein:

        ```
        {
           "auths": {
              "docker.pkg.github.com": {
                 "auth": "hier einfügen"
              }
           }
        }
        ```

        und speichern Sie dieses als `Secret Value`

## Committen und Pushen

Wenn alles passt, dann committen und pushen Sie jetzt. Der Drone CI-Job
wird dennoch fehlschlagen, weil noch kein Go-File vorhanden ist.

## Und jetzt nur noch die Aufgabe lösen und abgeben...

Erzeugen Sie **JETZT** den `develop`-Branch und beginnen Sie mit der eigentlichen
Arbeit. Ab jetzt ist der `master`-Branch für Sie tabu!

Orientieren Sie sich gerne an <https://github.com/vesose/example-micro>...
