import React from 'react';

const Form = ({ onChange, onSubmit, form, equipos }) => {
  return (
    <div className="container" id="container">
      <h1 className='publicar'>
        Resultado del Torneo
      </h1>
      <form
        className='form-publicar'
        onSubmit={onSubmit}
      >
        <div className="form-group">
          <label>Edicion del Torneo:</label>
          <input
            type="text"
            className="form-control inputp"
            name="id_edicion_torneo"
            onChange={onChange}
            value={form.id_edicion_torneo}
            disabled
          />
        </div>
        
        <div className="form-group">
          <label>Elige al Campeon:</label>
          <select
            className="form-control inputp"
            name="campeon"
            onChange={onChange}
            value={form.campeon}
          >
            <option value="" >Selecciona un Equipo</option>
            {equipos.map(equipo => (
              <option key={equipo.id} value={equipo.id}>
                {equipo.nombre}
              </option>
            ))}
          </select>
        </div>
        
        <div className="form-group">
          <label>Elige al Subcampeon:</label>
          <select
            className="form-control inputp"
            name="subcampeon"
            onChange={onChange}
            value={form.subcampeon}
          >
            <option value="" >Selecciona un Equipo</option>
            {equipos.map(equipo => (
              <option key={equipo.id} value={equipo.id}>
                {equipo.nombre}
              </option>
            ))}
          </select>
        </div>

        <div>
          <button
            type="submit"
            className="btn-summit"
          >
            Insertar Resultado
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





