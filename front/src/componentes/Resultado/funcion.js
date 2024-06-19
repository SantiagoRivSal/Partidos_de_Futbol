import React, { useState, useEffect } from "react";
import NewResultado from './nuevoResultado';
import Cookies from "universal-cookie";
import swal from "sweetalert2";

export const InsertResultado = () => {
  const cookies = new Cookies();
  const idEdicionTorneo = cookies.get("id_edicion_torneo");
  console.log("Valor de la cookie:", idEdicionTorneo);

  const [form, setForm] = useState({
    'id_edicion_torneo': idEdicionTorneo,
    'campeon': "",
    'subcampeon': "",
  });

  const [equiposEnEdicion, setEquiposEnEdicion] = useState([]);

  useEffect(() => {
    const obtenerEquiposEnEdicion = async () => {
      try {
        const response = await fetch(`http://localhost:8090/equiposxedicion/` + idEdicionTorneo);
        const data = await response.json();

        if (Array.isArray(data)) {
          const equiposDetalles = await Promise.all(data.map(async (equipo) => {
            const res = await fetch(`http://localhost:8090/equipo/` + equipo.id_equipo);
            const equipoData = await res.json();
            return { id: equipo.id_equipo, nombre: equipoData.nombre };
          }));
          setEquiposEnEdicion(equiposDetalles);
        } else {
          setEquiposEnEdicion([]);
          console.error("La respuesta no es un array:", data);
        }
      } catch (error) {
        console.error("Error al obtener equipos en la edición:", error);
        setEquiposEnEdicion([]);  
      }
    };

    obtenerEquiposEnEdicion();
  }, [idEdicionTorneo]);

  const handleChange = (event) => {
    const { name, value } = event.target;
    setForm({
      ...form,
      [name]: Number(value),
    });
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    // Verificar si ya existe un resultado para la edición del torneo
    try {
      const checkResponse = await fetch(`http://localhost:8090/resultados?edicion=${form.id_edicion_torneo}`);
      const checkData = await checkResponse.json();

      if (checkResponse.ok && checkData.length > 0) {
        swal.fire({
          icon: 'error',
          text: "Este torneo ya tiene un resultado cargado.",
        });
        return;
      }
    } catch (error) {
      console.error("Error al verificar resultado existente:", error);
      swal.fire({
        icon: 'error',
        text: "Error al verificar resultado existente.",
      });
      return;
    }

    if (form.campeon && form.subcampeon && form.id_edicion_torneo) {
      const requestOptions = {
        method: "POST",
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json',
          'Access-Control-Allow-Origin': '*'
        },
        body: JSON.stringify(form),
      };

      try {
        const response = await fetch('http://localhost:8090/resultado', requestOptions);

        if (response.ok) {
          swal.fire({
            icon: 'success',
            text: "Resultado Insertado",
          }).then((result) => {
            if (result.isConfirmed) {
              window.location.href = "http://localhost:3000/resultado"
            }
          });
        } else {
          const errorData = await response.json();
          throw new Error(errorData.message || "No se pudo insertar el Resultado");
        }
      } catch (error) {
        console.error("Error:", error);
        swal.fire({
          icon: 'error',
          text: error.message || "No se pudo insertar el Resultado",
        });
      }
    } else {
      swal.fire({
        icon: 'error',
        text: "No se pudo Insertar el Resultado. Asegúrate de completar todos los campos.",
      });
    }
  };

  return <NewResultado
    form={form}
    onChange={handleChange}
    onSubmit={handleSubmit}
    equipos={equiposEnEdicion}
  />;
};


