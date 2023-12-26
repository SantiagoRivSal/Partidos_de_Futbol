import React from "react";
import './Equipo.css';

export const PaisItem =(
    {id,
    nombre,
    bandera,
})=>{
    return(
        <div className="Pais">
        <div className="Pais_img">
          <img className="image" src={bandera} alt="" />
        </div>
        <div className="Pais_footer">
          <h1>{nombre}</h1>
        </div>
      </div>
    )
}