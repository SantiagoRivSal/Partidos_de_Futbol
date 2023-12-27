import React from "react";
//import './Equipo.css';

export const TorneoItem =(
    {id,
    nombre,
    logo,
})=>{
    return(
        <div className="Torneo">
        <div className="Torneo_img">
          <img className="image" src={logo} alt="" />
        </div>
        <div className="Torneo_footer">
          <h1>{nombre}</h1>
        </div>
      </div>
    )
}