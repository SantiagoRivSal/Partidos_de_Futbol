import React from "react";
import Oro from "../../imagenes/MedallaOro.png";
import Plata from "../../imagenes/MedallaPlata.png";

import '../Equipos/Equipo.css';

export const Podio = ({ campeon, subcampeon }) => {
  return (
    <div className="Equipos">
      <div className="Equipo">
        <div className="Trofeo_img">
          <img className="medal-image" src={Oro} alt="Medalla de Oro" />
        </div>
        <div className="Equipo_img">
          <img className="image" src={campeon.escudo} alt={`Escudo del Campeón: ${campeon.nombre}`} />
        </div>
        <div className="Equipo_footer">
            <h1>{campeon.nombre}</h1>
        </div>
      </div>
      <div className="Equipo">
        <div className="Trofeo_img">
          <img className="medal-image" src={Plata} alt="Medalla de Plata" />
        </div>
        <div className="Equipo_img">
          <img className="image" src={subcampeon.escudo} alt={`Escudo del Subcampeón: ${subcampeon.nombre}`} />
        </div>
        <div className="Equipo_footer">
            <h1>{subcampeon.nombre}</h1>
        </div>
      </div>
    </div>
  );
};


