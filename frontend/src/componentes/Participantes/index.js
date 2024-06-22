import React, { useEffect, useState } from 'react';

const Form = ({ onChange, onSubmit, form }) => {
  const [countryOptions, setCountryOptions] = useState([]);
  const [equipoOptions, setEquipoOptions] = useState([]);

  useEffect(() => {
    // Llamada al backend para obtener los valores de los países
    fetch('http://localhost:8090/paises')
      .then(response => response.json())
      .then(data => setCountryOptions(data))
      .catch(error => console.error('Error fetching country options:', error));
  }, []);

  useEffect(() => {
    // Llamada al backend para obtener los valores de los equipos filtrados por país
    if (form.id_pais) {
      fetch(`http://localhost:8090/equiposxpais/` + form.id_pais)
        .then(response => response.json())
        .then(data => {
          // Ordenar los equipos alfabéticamente por nombre
          const sortedEquipos = data.sort((a, b) => a.nombre.localeCompare(b.nombre));
          setEquipoOptions(sortedEquipos);
        })
        .catch(error => console.error('Error fetching equipo options:', error));
    }
  }, [form.id_pais]);

  return (
    <div className="container" id="container">
      <h1 className='publicar'>
        Equipos Participantes
      </h1>
      <form
        className='form-publicar'
        onSubmit={onSubmit}
      >
        {/* Input para el id_torneo */}
        <div className="form-group">
          <label>Edicion del Torneo:</label>
            <input
            type="text"
            className="form-control inputp"
            name="id_edicion_torneo"  // Cambia aquí
            onChange={onChange}
            value={form.id_edicion_torneo}
            disabled
            />
        </div>
        {/* Selector para el país */}
        
        <div className="form-group">
          <label>Selecciona un País:</label>
          <select className="form-control inputp" name="id_pais"
            onChange={onChange}
            value={form.id_pais}
          >
            <option value="" >Selecciona un País</option>
            {countryOptions.map(pais => (
              <option key={pais.id} value={pais.id}>
                {pais.nombre}
              </option>
            ))}
          </select>
        </div>
        {/* Selector para el equipo */}
        <div className="form-group">
          <label>Selecciona un Equipo:</label>
          <select
            className="form-control inputp"
            name="id_equipo"
            onChange={onChange}
            value={form.id_equipo}
          >
            <option value="" disabled>Selecciona un Equipo</option>
            {equipoOptions.map(equipo => (
              <option key={equipo.id} value={equipo.id}>
                {equipo.nombre}
              </option>
            ))}
          </select>
        </div>
        {/* Botón de submit */}
        <div>
          <button
            type="submit"
            className="btn-summit"
          >
            Insertar Equipo
          </button>
        </div>
      </form>
      <button
        type="button" 
        onClick={() => {
          window.location.href = "/ediciones/opciones";
        }}
        className="atras"
      >
        Atrás
      </button>
    </div>
  );
};

export default Form;


