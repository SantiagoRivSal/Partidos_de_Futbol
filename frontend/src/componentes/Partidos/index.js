import React, { useEffect, useState } from 'react';
import Cookies from "universal-cookie";
import { EquipoItem } from "../Partidos/EquipoItem";
import { Podio } from './PodioItem';

import swal from "sweetalert2";
import "./partidos.css";

async function GetEquiposByEdicion(id) {
  return fetch('http://https://backend-4ufveexwpa-uc.a.run.app/:4000/equiposxedicion/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(data => data.json());
}

async function GetEquipoById(id) {
  return fetch('http://https://backend-4ufveexwpa-uc.a.run.app/:4000/equipo/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(data => data.json());
}

async function GetResultadoByEdicion(id) {
  return fetch('http://https://backend-4ufveexwpa-uc.a.run.app/:4000/resultadoxedicion/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(data => data.json());
}

export const Match = () => {
  const cookies = new Cookies();
  const id = cookies.get("id_edicion_torneo");
  console.log("Valor de la cookie:", id);

  const [equipos, setEquipos] = useState([]);
  const [podio, setPodio] = useState(null);

  useEffect(() => {
    async function fetchData() {
      const response = await GetEquiposByEdicion(id);
      if (response.status === 400) {
        swal.fire({
          icon: 'error',
          text: "No hay Equipos que participen en esta edicion",
        });
      } else {
        const equiposWithDetails = await Promise.all(
          response.map(async (equipo) => {
            const equipoDetails = await GetEquipoById(equipo.id_equipo);
            return {
              ...equipo,
              nombre: equipoDetails.nombre,
              escudo: equipoDetails.escudo
            };
          })
        );
        setEquipos(equiposWithDetails);
        console.log(equiposWithDetails);
      }

      const resultado = await GetResultadoByEdicion(id);
      if (resultado && resultado.campeon && resultado.subcampeon) {
        const campeonDetails = await GetEquipoById(resultado.campeon);
        const subcampeonDetails = await GetEquipoById(resultado.subcampeon);
        setPodio({ 
          campeon: campeonDetails,
          subcampeon: subcampeonDetails
        });
      }
    }
    fetchData();
  }, [id]);

  return (
    <div className="match">
      <div className="match-header">
        <h1>Datos del Torneo</h1>
      </div>
      <div>
        <div className="Equipos">
          <div className="match-header">
            <h1>Participantes</h1>
          </div>
          {equipos.map(equipo => (
            <EquipoItem
              key={equipo.id_equipo}
              id={equipo.id_equipo}
              nombre={equipo.nombre}
              escudo={equipo.escudo}
            />
          ))}
        </div>
      </div>
      <div>
        <div className="Equipos">
            <div className="match-header">
              <h1>Podio</h1>
            </div>
              {podio?.campeon && podio?.subcampeon && (
                <Podio
                  campeon={podio.campeon}
                  subcampeon={podio.subcampeon}
                />
              )}
        </div>
      </div>
      <div>
        <button type="button" onClick={() => { window.location.href = "/ediciones/opciones"; }}
          className="atras"
        >
          Atrás
        </button>
      </div>
    </div>
  );
};








