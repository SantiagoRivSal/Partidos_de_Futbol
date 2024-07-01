Feature('fifa');

Scenario('Creacion de Torneos', ({ I }) => {
    I.amOnPage('http://localhost:3000/torneos');
    I.see('Creacion de Ediciones de Torneos');

    // Abre el desplegable y selecciona la opción deseada
    I.click('select[name="id_torneo"]');
    I.selectOption('select[name="id_torneo"]', 'Copa Libertadores'); // Ajusta el texto según el valor correcto

    // Rellena los otros campos
    I.fillField('input[name="anio"]', '2023');
    I.fillField('input[name="sede_final"]', 'Rio de Janeiro');
    
    // Clic en Crear Torneo
    I.click('button[type="submit"]');
    
    // Verifica que el torneo ha sido creado
    I.see('Torneo Creado');
});



