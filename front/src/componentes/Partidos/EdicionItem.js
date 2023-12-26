import React from "react";
import './Edicion.css';

export const EdicionItem = ({ id, anio, sede_final, id_torneo }) => {
  return (
    <div className="Edicion">
      <div className="Edicion_footer">
        <h1>{anio}</h1>
        <h2>{sede_final}</h2>
      </div>
    </div>
  );
};