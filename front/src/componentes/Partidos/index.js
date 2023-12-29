import React from "react";
import "./partidos.css"

export const Match = () => {
  return (
    <div className="match">
      <div className="match-header">
        <div className="match-status">NOMBRE DE FASE</div>
      </div>
      <div className="match-content">
        {/* Columna para el equipo local */}
        <div className="column">
          <div className="team team--home">
            <div className="team-logo">
              <img src="https://assets.codepen.io/285131/chelsea.svg" alt="Chelsea Logo" />
            </div>
            <h2 className="team-name">Chelsea</h2>
          </div>
        </div>
        {/* Columna para los detalles del partido */}
        <div className="column">
          <div className="match-details">
            <div className="match-score">
              <span className="match-score-number match-score-number--leading">3</span>
              <span className="match-score-divider">-</span>
              <span className="match-score-number">1</span>
            </div>
            <button className="match-bet-place">Insertar Resultado del Partido</button>
          </div>
        </div>
        {/* Columna para el equipo visitante */}
        <div className="column">
          <div className="team team--away">
            <div className="team-logo">
              <img
                src="https://resources.premierleague.com/premierleague/badges/t1.svg"
                alt="Man Utd Logo"
              />
            </div>
            <h2 className="team-name">Man Utd</h2>
          </div>
        </div>
      </div>
    </div>
  );
};

