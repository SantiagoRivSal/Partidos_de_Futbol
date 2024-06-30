Feature('crud');

Scenario('Get Equipos by id', ({ I }) => {
    I.amOnPage("/")
    I.fillField('#idinsert', '1');
    I.click('#getequiposbyid');
    I.see('Equipo: {"id":1,"nombre":Boca Juniors,"escudo":https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjOL0rL1mnaqAIv_9cWnQayBcfquddQxXR82Ho7GG9YdcQjYhK077s3wlCOM1Z1OExJShV-b-f3Yzuoq049fzMYoRvgwR06s99-ezsl4uki7P80_dKIxjtIO4kDkpJNbJTxG-K2jS8xkKzvkTgkVXfy2sosC1ZROpCqNkNCiGJBhKbV9_u86I0Bq4wsNgM/s128/Boca%20Juniors128x.png,"id_pais":1}');
})