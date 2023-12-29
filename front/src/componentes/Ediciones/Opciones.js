import React from "react";
import "./Opciones.css";
import Logo from "../../imagenes/Logo_Torneo.png"
import Cuadro from "../../imagenes/Cuadro.png"


export const MenuOpciones = () => {
  return (
    <div className="opciones-container">
      {/* Botón 1 */}
      <a href="/participantes">
        <button type="button" className="Opciones">
        <div className="logo">
                <img src={Logo} alt="logo" width="200"/>
            </div>
          Insertar Equipos en la Edicion
        </button>
      </a>

      {/* Botón 2 */}
      <a href="/partidos">
        <button type="button" className="Opciones">
        <div className="logo">
                <img src={Cuadro} alt="cuadro" width="200"/>
            </div>
          Crear Partidos
        </button>
      </a>

      {/* Botón de retroceso */}
      <button
        type="button" 
        onClick={() => {
          window.location.href = "/ediciones";
        }}
        className="retroceder"
      >
        Retroceder
      </button>
    </div>
  );
};


