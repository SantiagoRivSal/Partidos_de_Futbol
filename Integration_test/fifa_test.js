Feature('fifa');

Scenario('Creacion de Torneos', ({ I }) => {
    I.amOnPage('http://localhost:3000/torneos');
    I.see('Creacion de Ediciones de Torneos');

    // Abre el desplegable y selecciona la opción deseada
    I.click('select[name="id_torneo"]');
    I.selectOption('select[name="id_torneo"]', 'Copa Sudamericana'); // Ajusta el texto según el valor correcto

    // Rellena los otros campos
    I.fillField('input[name="anio"]', '2021');
    I.fillField('input[name="sede_final"]', 'Cordoba');
    
    // Clic en Crear Torneo
    I.click('button[type="submit"]');
    I.wait(10);
});

Scenario('Carga de países y selección de equipos por país', async ({ I }) => {
    I.amOnPage('/');
    I.seeElement('.paises');
  
    // Hacer clic en cada botón de país y verificar que los equipos se cargan
    const paises = [
      'Argentina',
      'Brasil',
      'Uruguay',
      'Colombia',
      'Chile',
      'Peru',
      'Ecuador',
      'Paraguay',
      'Venezuela'
    ];
  
    for (const pais of paises) {
      I.click(`button:has-text("${pais}")`);
      I.waitForElement('.Equipos', 10);
      I.seeElement('.Equipos div');
      I.wait(5); // Espera 5 segundos para ver los equipos antes de pasar al siguiente país
    }
  });





