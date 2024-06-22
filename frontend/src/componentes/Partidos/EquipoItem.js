import React from "react";
import '../Equipos/Equipo.css';

export const EquipoItem = ({ id, nombre, escudo }) => {
  return (
    <div className="Equipo">
      <a href="#">
        <div className="Equipo_img">
          <img className="image" src={escudo} alt="" />
        </div>
      </a>
      <div className="Equipo_footer">
        <h1>{nombre}</h1>
      </div>
    </div>
  );
};