import React, { useEffect, useState } from 'react';
import Cookies from "universal-cookie";
import { EquipoItem } from "../Equipos/EquipoItem";
import swal from "sweetalert2";
import "./partidos.css";

async function GetEquiposByEdicion(id) {
  return fetch('http://localhost:8090/equiposxedicion/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(data => data.json());
}

async function GetEquipoById(id) {
  return fetch('http://localhost:8090/equipo/' + id, {
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

  useEffect(() => {
    async function fetchData() {
      const response = await GetEquiposByEdicion(id);
      if (response.status === 400) {
        swal.fire({
          icon: 'error',
          text: "No hay Equipos que participen en esta edicion",
        });
      } else {
        // Obteniendo los detalles de cada equipo
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
    }
    fetchData();
  }, [id]); // El efecto se ejecutará cuando `id` cambie

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
              {
                equipos.map(equipo => (
                  <EquipoItem
                    key={equipo.id_equipo}
                    id={equipo.id_equipo}
                    nombre={equipo.nombre}
                    escudo={equipo.escudo}
                  />
                ))
              }
            </div>
          </div>
          <div>
            <button type="button" onClick={() => { window.location.href = "/ediciones/opciones";
              }}
              className="atras"
            >
              Atrás
            </button>
        </div>
    </div>
  );
};





