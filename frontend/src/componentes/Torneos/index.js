import React, { useEffect, useState } from 'react';
import './torneo.css';

const Form = ({ onChange, onSubmit, form }) => {
    const [torneoOptions, setTorneoOptions] = useState([]);
  
    useEffect(() => {
      // Aquí realiza una llamada al backend para obtener los valores del torneo
      // Reemplaza la URL con la ruta correcta de tu backend
      fetch('https://backend-4ufveexwpa-uc.a.run.app/torneos') 
        .then(response => response.json())
        .then(data => setTorneoOptions(data))
        .catch(error => console.error('Error fetching torneo options:', error));
    }, []);
  
    return (
      <div className="container" id="container">
        <h1 className='publicar'>
          Creacion de Ediciones de Torneos
        </h1>
        <form
          className='form-publicar'
          onSubmit={onSubmit}
        >
          <div className="form-group">
            <select
              className="form-control inputp"
              name="id_torneo"
              onChange={onChange}
              value={form.id_torneo}
            >
              <option value="" disabled>Selecciona un Torneo</option>
              {torneoOptions.map(torneo => (
                <option key={torneo.id} value={torneo.id}>
                  {torneo.nombre}
                </option>
              ))}
            </select>
          </div>
          <div className="form-group">
            <input
              type="text"
              className="form-control inputp"
              placeholder="Año de la Edicion"
              name="anio"
              onChange={onChange}
              value={form.anio}
            />
          </div>
          <div className="form-group">
            <input
              type="text"
              className="form-control inputp"
              placeholder="Sede Final"
              name="sede_final"
              onChange={onChange}
              value={form.sede_final}
            />
          </div>
          <div>
            <button
              type="submit"
              className="btn-summit"
            >
              Crear Torneo
            </button>
          </div>
        </form>
      </div>
    );
  };
  
  export default Form;
