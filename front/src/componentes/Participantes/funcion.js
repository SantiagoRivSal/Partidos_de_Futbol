import React, { useState } from "react";
import NewParticipante from './nuevoParticipante'; // Reemplaza con tu componente Form real
import Cookies from "universal-cookie";
import swal from "sweetalert2";

export const InsertParticipante = () => {
    const cookies = new Cookies();
    const idEdicionTorneo = cookies.get("id_edicion_torneo");
    console.log("Valor de la cookie:", idEdicionTorneo);
  const [form, setForm] = useState({
    'id_edicion_torneo': idEdicionTorneo, // Asegúrate de obtener este valor de las cookies
    'id_equipo': "", // Añadido el campo nombre_equipo para representar el nombre del equipo
  });

  const handleChange = (event) => {
    const { name, value } = event.target;
    setForm({
      ...form,
      [name]: name === "id_equipo" || name === "id_edicion_torneo" ? Number(value) : value,
    });
  };

  const handleSubmit = async event => {
    event.preventDefault();
    if (Number(form.id_edicion_torneo)>0 && Number(form.id_equipo)>0) {
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
        const response = await fetch('http://localhost:8090/equipoxedicion', requestOptions);
        if (response.ok) {
          swal.fire({
            icon: 'success',
            text: "Equipo Insertado",
          }).then((result) => {
            if (result.isConfirmed) {
              window.location.href = "http://localhost:3000/participantes"
            }
          });
        } else {
          throw new Error("No se pudo insertar el Equipo");
        }
      } catch (error) {
        console.error("Error:", error);
        swal.fire({
          icon: 'error',
          text: "No se pudo insertar el Equipo",
        })
      }
    } else {
      swal.fire({
        icon: 'error',
        text: "No se pudo Insertar el Equipo. Asegúrate de completar todos los campos.",
      });
    }
  };

  return <NewParticipante
    form={form}
    onChange={handleChange}
    onSubmit={handleSubmit}
  />;
};

