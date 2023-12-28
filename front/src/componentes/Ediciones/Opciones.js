import React from "react";
import "./Opciones.css";
import Logo from "../../imagenes/Logo_Torneo.png"

export const MenuOpciones = () => {
  return (
    <div className="opciones-container">
      {/* Botón 1 */}
      <a href="/participantes">
        <button type="button" className="Opciones">
        <div className="logo">
                <img src={Logo} alt="logo" width="150"/>
            </div>
          Insertar Equipos en la Edicion
        </button>
      </a>

      {/* Botón 2 */}
      <a href="/pagina2">
        <button type="button" className="Opciones">
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


