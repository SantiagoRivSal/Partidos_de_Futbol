import React, { useState, useEffect } from "react";
import NewResultado from './nuevoResultado';
import Cookies from "universal-cookie";
import swal from "sweetalert2";

export const InsertResultado = () => {
  const cookies = new Cookies();
  const idEdicionTorneo = cookies.get("id_edicion_torneo");
  console.log("Valor de la cookie:", idEdicionTorneo);
 
  const [form, setForm] = useState({
    id_edicion_torneo: idEdicionTorneo,
    campeon: "",
    subcampeon: "",
  });

  const [equiposEnEdicion, setEquiposEnEdicion] = useState([]);
  const [cargando, setCargando] = useState(true);
  const [resultadoExistente, setResultadoExistente] = useState(false);

  useEffect(() => {
    const obtenerEquiposEnEdicion = async () => {
      if (!idEdicionTorneo) {
        console.error("No hay idEdicionTorneo en las cookies");
        setCargando(false);
        return;
      }

      try {
        const response = await fetch(`http://backend-4ufveexwpa-uc.a.run.app:4000/equiposxedicion/${idEdicionTorneo}`);
        const data = await response.json();

        if (Array.isArray(data)) {
          const equiposDetalles = await Promise.all(data.map(async (equipo) => {
            const res = await fetch(`http://backend-4ufveexwpa-uc.a.run.app:4000/equipo/${equipo.id_equipo}`);
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
      } finally {
        setCargando(false);
      }
    };

    const verificarResultadoExistente = async () => {
      if (!idEdicionTorneo) {
        return;
      }

      try {
        const response = await fetch(`http://backend-4ufveexwpa-uc.a.run.app:4000/resultadoxedicion/${idEdicionTorneo}`);
        const data = await response.json();

        if (response.ok && data.length > 0) {
          setResultadoExistente(true);
        } else {
          setResultadoExistente(false);
        }
      } catch (error) {
        console.error("Error al verificar resultado existente:", error);
      }
    };

    obtenerEquiposEnEdicion();
    verificarResultadoExistente();
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

    if (resultadoExistente) {
      swal.fire({
        icon: 'error',
        text: "Este torneo ya tiene un resultado cargado.",
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
        const response = await fetch('http://backend-4ufveexwpa-uc.a.run.app:4000/resultado', requestOptions);

        if (response.ok) {
          swal.fire({
            icon: 'success',
            text: "Resultado Insertado",
          }).then((result) => {
            if (result.isConfirmed) {
              window.location.href = "http://backend-4ufveexwpa-uc.a.run.app:3000/resultado"
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

  if (cargando) {
    return <div>Cargando...</div>;
  }

  return (
    <NewResultado
      form={form}
      onChange={handleChange}
      onSubmit={handleSubmit}
      equipos={equiposEnEdicion}
    />
  );
};
