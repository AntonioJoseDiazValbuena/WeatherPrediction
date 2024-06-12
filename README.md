# Weather prediction

This application is meant to predict the weather between three planets (Vulcano, Ferengi and Betazoide) deppending on the position between them and their sun in a period of 10 years

## Set up

To be able to use this app, first you have to execute the script that is in the scripts folder in a mysql local server.

After executing the script, you have to modify the connection string on the main.go file in the api folder, in line 19, to match the local server and password that you have the database on, and then you will be able to execute the app.

## Execute the app

When you want to execute the app you can do it either with the command "go run main" or debug it with visual studio code itself on the left bar.

If you are executing the app for the first time you have to go first to the configurate section of the app ("http://localhost:8080/configurate").

If you already executed the configuration of the app, then you are already able to use the app changing in the link the "configurate" to "weather?day="
putting the day that you want in front of the equal between 1 and 3650.