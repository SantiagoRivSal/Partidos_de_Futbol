import React, { useContext, useEffect, useState} from "react";
import { EquipoItem } from "./EquipoItem";
import "bootstrap/dist/css/bootstrap.min.css";
import './Equipo.css';
import swal from "sweetalert2";

async function GetEquiposByIdPais(id) {
    return fetch('http://https://backend-4ufveexwpa-uc.a.run.app/:4000/equiposxpais/' +id, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
      .then(data => data.json())
   }

export const PaisesLista =()=>{   
    const [paises,setPaises] = useState([]);
    const fetchApi = async()=>{
    const response = await fetch('http://https://backend-4ufveexwpa-uc.a.run.app/:4000/paises')
    .then((response) => response.json());
    setPaises(response);
    };
    useEffect(()=>{
    fetchApi();
    },[])

    const [equipos,setEquipos] = useState([]);
    async function Handle (id) {
    const response = await GetEquiposByIdPais(id)
    if (response.status == 400) {
      swal.fire({
        icon: 'error',
        text: "No hay Equipos que pertenezcan a dicho pais",
      })
   }else{
    setEquipos(response)
    console.log(response);
   }
    };
const Render =(
        <div className="Equipos">
            {
                equipos.map(equipo =>(
                  <EquipoItem key={equipo.id}
                  id={equipo.id}
                  nombre={equipo.nombre}
                  escudo={equipo.escudo}
                  id_pais={equipo.id_pais}
                  /> 
                ))
            }
            
        </div>
)
    return(
        <>
        <div className="paises"> 
            {
                paises.map(pais =>(
                    <button className="paises" onClick={() => Handle(pais.id)}>
                    <div className="bandera">
                        <img src={pais.bandera} alt="" />
                    </div>
                    <span className="nombre">{pais.nombre}</span>
                    </button>
                ))
            }
        </div>
          <div>
              {Render}
          </div>
        </>
    );
}