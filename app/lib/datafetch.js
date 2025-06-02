// Beispiel: GET-Anfrage an den API-Endpunkt
fetch('http://localhost:3006/api/dein-endpunkt') // Ersetze 'dein-endpunkt' durch den tatsächlichen Pfad
  .then(response => {
    if (!response.ok) {
      throw new Error('Netzwerkantwort war nicht ok');
    }
    return response.json(); // Falls der Server JSON zurückgibt
  })
  .then(data => {
    console.log('Antwort vom Server:', data);
    // Hier kannst du die Daten z. B. im DOM anzeigen
  })
  .catch(error => {
    console.error('Fehler beim Abrufen der Daten:', error);
  });
