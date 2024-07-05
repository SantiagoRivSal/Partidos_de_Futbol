
import React, { useState, useEffect } from "react";
import './Edicion.css';
import { EdicionItem } from "./EdicionItem";
//import { TorneoItem } from "./TorneoItem";
import "bootstrap/dist/css/bootstrap.min.css";
import swal from "sweetalert2";

async function GetEdicionesByTorneo(id) {
  return fetch(`https://backend-4ufveexwpa-uc.a.run.app/ediciones_de_torneo/`+ id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  }).then((data) => data.json());
}

export const TorneosLista = () => {
  const [torneos, setTorneos] = useState([]);
  const fetchTorneos = async () => {
    const response = await fetch("https://backend-4ufveexwpa-uc.a.run.app/torneos").then((response) =>
      response.json() 
    );
    setTorneos(response);
  };
  useEffect(() => {
    fetchTorneos();
  }, []);

  const [ediciones, setEdiciones] = useState([]);

  async function Handle(id) {
    const response = await GetEdicionesByTorneo(id);
    if (response.status === 400) {
      swal.fire({
        icon: 'error',
        text: "No hay Ediciones de ese Torneo",
      });
    } else {
      // Ordenar las ediciones de menor a mayor según el año
      const edicionesOrdenadas = response.sort((a, b) => a.anio - b.anio);
      setEdiciones(edicionesOrdenadas);
      console.log(edicionesOrdenadas);
    }
  };

    const Render =(
      <div className="Ediciones">
          {
              ediciones.map(edicion =>(
                <EdicionItem key={edicion.id}
                id={edicion.id}
                anio={edicion.anio}
                sede_final={edicion.sede_final}
                id_torneo={edicion.id_torneo}
                /> 
              ))
          }
          
      </div>
)

  return (
    <>
        <div className="Torneo">
            {
                torneos.map(torneo =>(
                    <button className="torneo" onClick={() => Handle(torneo.id)}>
                    <div className="logo">
                        <img src={torneo.logo} alt="" />
                    </div>
                    </button>
                ))
            }
        </div>
      <div>
          {Render}
      </div>

    </>
  );
};
