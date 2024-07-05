import React, { useState } from "react";
import NewEdition from './nuevoTorneo'; // Replace with your actual NewEdition component
import Cookies from "universal-cookie";
import swal from "sweetalert2";

export const InsertEdition = () => {
    const cookies = new Cookies();
    const [form, setForm] = useState({
        'id_torneo': "",
        'anio': "",
        'sede_final': "",

      });
  
      const handleChange = (event) => {
        const { name, value } = event.target;
        setForm({
          ...form,
          [name]: name === "id_torneo" || name === "anio" ? Number(value) : value,
        });
      };


    const handleSubmit = async event => {
      event.preventDefault();
      if(Number(form.id_torneo)>0 && Number(form.anio)>0){
        const requestOptions = {
        method: "POST",
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json',
          'Access-Control-Allow-Origin': '*'
        },
        body: JSON.stringify({
          ...form,
          idTorneo: Number(form.id_torneo), // Convertir a número
          anio: Number(form.anio), // Convertir a número
        }),
      };
      try {
        const response = await fetch('http://backend-4ufveexwpa-uc.a.run.app:4000/edicion_de_torneo', requestOptions);
        if (response.ok) {
          const data = await response.json();
          swal.fire({
            icon: 'success',
            text: "Torneo creado",
        }).then((result) => { 
          if (result.isConfirmed) {
            window.location.href = "http://backend-4ufveexwpa-uc.a.run.app:3000/torneos"
          }});
          // Aquí podrías hacer algo con la respuesta, si es necesario
        } else {
          throw new Error("No se pudo crear el Torneo");
        }
      } catch (error) {
        console.error("Error:", error);
        swal.fire({
          icon: 'error',
          text: "No se pudo crear el Torneo",
        })
      }
      console.log(form)}
      else{
        swal.fire({
          icon: 'error',
          text: "No se pudo crear el Torneo",
        }) 
      }
    };

    return <NewEdition 
    form={form}
    onChange={handleChange}
    onSubmit={handleSubmit}
/>


};