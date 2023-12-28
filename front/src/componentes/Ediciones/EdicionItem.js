import React from "react";
import Cookies  from "universal-cookie";
import {Link} from "react-router-dom";
import './Edicion.css';

export const EdicionItem = ({ id, anio, sede_final, id_torneo }) => {
  const cookies = new Cookies();
  const handleClick = () => {
    // Guardar el ID de la propiedad en una cookie con nombre "propertyId"
    cookies.set("id_edicion_torneo", id, { path: "/" });
  };
  return (
    <Link to="/ediciones/opciones" className="Edicion" onClick={handleClick}>
      <div className="Edicion_footer">
        <h1>{anio}</h1>
        <h2>{sede_final}</h2>
      </div>
    </Link>
  );
};